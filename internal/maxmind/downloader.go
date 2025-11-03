package maxmind

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/poly-workshop/geoip/internal/configs"
)

const (
	maxMindDownloadBaseURL = "https://download.maxmind.com/geoip/databases"
	userAgent              = "poly-workshop-geoip/1.0"
)

type datasetDefinition struct {
	editionID string
	path      string
}

// EnsureDatabases downloads missing MaxMind database files using the provided license key.
// If MAXMIND_LICENSE_KEY is not set, the function logs the omission and returns without downloading.
func EnsureDatabases(ctx context.Context, cfg configs.GRPCServerConfig) error {
	licenseKey := strings.TrimSpace(cfg.MaxMindLicenseKey)
	accountID := strings.TrimSpace(cfg.MaxMindAccountID)
	if licenseKey == "" || accountID == "" {
		slog.Info("MaxMind credentials not provided, skipping automatic database download")
		return nil
	}

	datasets := []datasetDefinition{
		{editionID: "GeoIP2-City", path: cfg.CityDBPath},
		{editionID: "GeoIP2-Country", path: cfg.CountryDBPath},
		{editionID: "GeoIP2-Enterprise", path: cfg.EnterpriseDBPath},
		{editionID: "GeoIP2-Anonymous-IP", path: cfg.AnonymousIPDBPath},
		{editionID: "GeoLite2-ASN", path: cfg.ASNDBPath},
		{editionID: "GeoIP2-Connection-Type", path: cfg.ConnectionTypeDBPath},
		{editionID: "GeoIP2-Domain", path: cfg.DomainDBPath},
		{editionID: "GeoIP2-ISP", path: cfg.ISPDBPath},
	}

	client := &http.Client{Timeout: 5 * time.Minute}

	for _, dataset := range datasets {
		if dataset.path == "" {
			continue
		}

		if _, err := os.Stat(dataset.path); err == nil {
			slog.Debug("MaxMind dataset already present", "path", dataset.path)
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("cannot stat dataset %q: %w", dataset.path, err)
		}

		slog.Info("Downloading MaxMind dataset", "edition", dataset.editionID, "path", dataset.path)
		if err := downloadDataset(ctx, client, accountID, licenseKey, dataset.editionID, dataset.path); err != nil {
			return fmt.Errorf("download %s: %w", dataset.editionID, err)
		}
	}

	return nil
}

func downloadDataset(ctx context.Context, client *http.Client, accountID, licenseKey, editionID, destPath string) error {
	u := fmt.Sprintf("%s/%s/download", maxMindDownloadBaseURL, url.PathEscape(editionID))
	params := url.Values{}
	params.Set("suffix", "tar.gz")
	u = fmt.Sprintf("%s?%s", u, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)
	req.SetBasicAuth(accountID, licenseKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 8*1024))
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("creating target directory: %w", err)
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(destPath), "maxmind-*.mmdb")
	if err != nil {
		return fmt.Errorf("creating temp file: %w", err)
	}
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if err := extractMMDB(resp.Body, tmpFile); err != nil {
		return err
	}

	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("closing temp file: %w", err)
	}

	if err := os.Rename(tmpFile.Name(), destPath); err != nil {
		return fmt.Errorf("replacing target file: %w", err)
	}

	return nil
}

func extractMMDB(reader io.Reader, dst io.Writer) error {
	gz, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("creating gzip reader: %w", err)
	}
	defer gz.Close()

	tarReader := tar.NewReader(gz)

	for {
		header, err := tarReader.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("reading tar: %w", err)
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		if filepath.Ext(header.Name) != ".mmdb" {
			continue
		}

		if _, err := io.Copy(dst, tarReader); err != nil {
			return fmt.Errorf("copying mmdb contents: %w", err)
		}

		return nil
	}

	return fmt.Errorf("mmdb file not found in archive")
}
