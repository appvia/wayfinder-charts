SHELL = /bin/bash -e
AUTHOR_EMAIL=wayfinder@appvia.io
ROOT_DIR=${PWD}
UNAME := $(shell uname)

.PHONY: charts
charts:
	@echo "--> Building charts"
	@./scripts/build_charts.sh

.PHONY: update-charts
update-charts:
	@echo "--> Updating charts"
	@go run ./cmd/updatecharts/ update-charts
