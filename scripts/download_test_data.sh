#!/bin/bash

set -euo pipefail

mkdir -p test/data

base_url="https://raw.githubusercontent.com/maxmind/MaxMind-DB/refs/heads/main/test/data"
files=(
	"GeoIP2-City-Test.mmdb"
	"GeoIP2-Country-Test.mmdb"
	"GeoIP2-Enterprise-Test.mmdb"
	"GeoIP2-Anonymous-IP-Test.mmdb"
	"GeoLite2-ASN-Test.mmdb"
	"GeoIP2-Connection-Type-Test.mmdb"
	"GeoIP2-Domain-Test.mmdb"
	"GeoIP2-ISP-Test.mmdb"
)

for file in "${files[@]}"; do
	target_path="test/data/${file}"
	if [[ -f "${target_path}" ]]; then
		echo "${target_path} already exists, skipping"
		continue
	fi

	wget -O "${target_path}" "${base_url}/${file}"
done
