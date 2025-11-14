SHELL:=/bin/bash

e2e_dir=$(shell cd ${PWD}/../.. && pwd -P)

gcp/setup-pubsub:
	./setup/pubsub/setup_for_server.sh
	./setup/pubsub/setup_for_agent.sh
.PHONY: gcp/setup-pubsub

gcp/setup-maestro:
	./setup/maestro.sh
.PHONY: gcp/setup-maestro

gcp/setup-agent:
	./setup/agent.sh
.PHONY: gcp/setup-agent

gcp/setup-e2e:
	./setup/e2e.sh
.PHONY: gcp/setup-e2e

gcp/e2e-test: gcp/setup-e2e
	ginkgo -v --fail-fast --label-filter="!(e2e-tests-spec-resync-reconnect||e2e-tests-status-resync-reconnect)" \
	--output-dir="$(e2e_dir)/report" --json-report=report.json --junit-report=report.xml \
	${e2e_dir}/pkg -- \
	-api-server="http://127.0.0.1:8000" \
	-grpc-server="127.0.0.1:8090" \
	-server-kubeconfig=$(KUBECONFIG) \
	-agent-kubeconfig=$(KUBECONFIG) \
	-consumer-name=${PWD}/_output/consumer_id
.PHONY: gcp/e2e-test

gcp/teardown:
	./setup/teardown.sh
.PHONY: gcp/teardown
