#!/usr/bin/env bash

set -euo pipefail
${TRACE:+set -x}

err_exit() {
  echo "Error: $1"
  exit 1
}

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." >/dev/null 2>&1 && pwd)"
CHARTS_DIR="${PROJECT_DIR}/charts"
VERSION="${VERSION:-latest}"

TMP_DIR="${PROJECT_DIR}/charts-compiled"
rm -rf $TMP_DIR
mkdir $TMP_DIR

cd "${PROJECT_DIR}"
CHARTS=$(find charts -mindepth 1 -maxdepth 1 -type d | cut -d '/' -f 2)

cd "${TMP_DIR}"

cp "${CHARTS_DIR}/Dockerfile" "${TMP_DIR}"
cp "${CHARTS_DIR}/set-charts-url.sh" "${TMP_DIR}"

helm repo index .

for chart_dir in "${CHARTS_DIR}"/*; do
  [ -d "${chart_dir}" ] || continue

  chart=$(basename "${chart_dir}")
  if [ "${chart}" == "wayfinder" ]; then
    continue
  fi

  cp -a "${chart_dir}" "${TMP_DIR}/${chart}"
  for ver_chart in "${TMP_DIR}/${chart}"/*; do
    [ -f "${ver_chart}/Chart.yaml" ] || continue
    helm package "${ver_chart}"
  done

done

helm repo index . --url "https://charts.wayfinder.run" --merge index.yaml
