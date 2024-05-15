#!/bin/sh

CHARTS_URL="${CHARTS_PROTOCOL:-http}:\/\/${CHARTS_HOSTNAME}"

sed s/https:\/\/appvia.github.io\/wayfinder-charts/$CHARTS_URL/g /usr/share/nginx/html/index.yaml.tmpl > /usr/share/nginx/html/index.yaml