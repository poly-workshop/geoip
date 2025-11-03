#!/bin/bash

set -euo pipefail

mkdir -p data

base_url="https://github.com/LOVECHEN/GeoLite.mmdb/raw/download"
files=(
	"GeoLite2-ASN.mmdb"
	"GeoLite2-City.mmdb"
    "GeoLite2-Country.mmdb"
)

for file in "${files[@]}"; do
	target_path="data/${file}"
	if [[ -f "${target_path}" ]]; then
		echo "${target_path} already exists, skipping"
		continue
	fi

	wget -O "${target_path}" "${base_url}/${file}"
done