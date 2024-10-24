#!/bin/sh

CHARTS_URL="${CHARTS_PROTOCOL:-http}://${CHARTS_HOSTNAME}"

sed s~https://charts.wayfinder.run~$CHARTS_URL~g /usr/share/nginx/html/index.yaml.tmpl > /usr/share/nginx/html/index.yaml