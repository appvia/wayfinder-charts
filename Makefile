SHELL = /bin/bash -e
AUTHOR_EMAIL=wayfinder@appvia.io
ROOT_DIR=${PWD}
UNAME := $(shell uname)
REGISTRY ?= quay.io
REGISTRY_ORG ?= appvia-wf-dev
VERSION ?= latest

.PHONY: charts
charts:
	@echo "--> Building charts"
	@./scripts/build_charts.sh

.PHONY: update-charts
update-charts:
	@echo "--> Updating charts"
	@go run ./cmd/updatecharts/ update-charts

charts-image: charts
	@echo "--> Building charts docker image ${REGISTRY}/${REGISTRY_ORG}/charts:${VERSION}"
	docker build --platform amd64 -t ${REGISTRY}/${REGISTRY_ORG}/charts:${VERSION} -f charts-compiled/Dockerfile charts-compiled
	docker push ${REGISTRY}/${REGISTRY_ORG}/charts:${VERSION}