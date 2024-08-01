# Maestro

## Deploy Maestro

### Run Maestro in KinD

```bash
make e2e-test/setup
```

### Run Maestro in OpenShift

```bash
export container_tool=docker
export GOPATH=$HOME/go
export USER=demo
export CLIENT_ID=demo
export CLIENT_SECRET=demo
export ENABLE_JWT=false
export ENABLE_AUTHZ=false
export external_apps_domain=$(oc get ingresses.config.openshift.io cluster -o jsonpath='{.spec.domain}')
oc create ns maestro-demo
make deploy
```

```bash
oc -n maestro-demo get svc
oc -n maestro-demo get pod
oc -n maestro-demo logs -f deploy/maestro -f
```

```bash
export consumer_id=$(curl -k -X POST -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/consumers -d '{"name": "cluster1"}' | jq -r .id)
export agent_namespace=maestro-agent-demo
make deploy-agent
```

```bash
oc -n maestro-agent-demo get pod
oc -n maestro-agent-demo get pod -o yaml | grep consumer-name
```

### Run Maestro in ARO

Follow the document at: https://github.com/Azure/ARO-HCP/blob/main/dev-infrastructure/docs/development-setup.md

## How to Use Maestro

### RESTFul API

```bash
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/consumers | jq
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources | jq
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resource-bundles | jq
curl -k -X POST -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources --data-binary @example/resource.json | jq
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources/be3cf895-be90-43b8-862b-0a401fa1f858 | jq
curl -k -X PATCH -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources/be3cf895-be90-43b8-862b-0a401fa1f858 --data-binary @example/resource.json | jq
curl -k -X DELETE -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources/5c61adfb-14bb-433a-929b-7b98a69cbdde
```

```bash
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/consumers | jq
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources | jq
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resource-bundles | jq
curl -k -X POST -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources --data-binary @example/resource.json | jq
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources/d338d754-c8f8-4b1e-8e4c-cc6a574a1af7 | jq
curl -k -X PATCH -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources/d338d754-c8f8-4b1e-8e4c-cc6a574a1af7 --data-binary @example/resource.json | jq
curl -k -X DELETE -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources/c31ea123-bf13-4dba-a70f-0a747cb23336 | jq
```

### GRPC API

1. Enable GRPC Server:

```bash
oc -n maestro-demo edit deploy/maestro
oc -n maestro-demo patch deploy/maestro --type=json -p='[{"op": "add", "path": "/spec/template/spec/containers/0/command/-", "value": "--enable-grpc-server=true"}]'
```

2. Port Forward for GRPC Server:

```bash
oc -n maestro-demo port-forward svc/maestro-grpc 8090:8090
```

## CURD Manifest:

```bash
# get consumer
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/consumers | jq

# create
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent.json

# check resource
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/resources | jq

# update
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent.update.json

# delete
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent.delete.json
```

## CURD ManifestBundle:

```bash
# get consumer
curl -k -X GET -H "Content-Type: application/json" https://localhost:30080/api/maestro/v1/consumers | jq

# create
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent-bundle.json

# check resource bundle
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resource-bundles | jq

# update
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent-bundle.update.json

# delete
go run ./example/grpcclient.go -grpc_server localhost:30090 -cloudevents_json_file ./example/cloudevent-bundle.delete.json
```
