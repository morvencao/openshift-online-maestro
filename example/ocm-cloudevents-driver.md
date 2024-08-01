## Enable CloudEvents Driver form ClusterManager

```bash
cat << EOF | oc apply -f -
apiVersion: operator.open-cluster-management.io/v1
kind: ClusterManager
metadata:
  name: cluster-manager
spec:
  addOnManagerImagePullSpec: quay.io/open-cluster-management/addon-manager:latest
  deployOption:
    mode: Default
  placementImagePullSpec: quay.io/open-cluster-management/placement:latest
  registrationConfiguration:
    featureGates:
    - feature: DefaultClusterSet
      mode: Enable
  registrationImagePullSpec: quay.io/open-cluster-management/registration:latest
  workImagePullSpec: quay.io/morvencao/work:dev
  workConfiguration:
    workDriver: mqtt
    featureGates:
    - feature: ManifestWorkReplicaSet
      mode: Enable
    - feature: CloudEventsDrivers
      mode: Enable
EOF
```

## Create a secret for the CloudEvents Driver

```bash
cat << EOF | oc apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: work-driver-config
  namespace: open-cluster-management
stringData:
  config.yaml: |
    brokerHost: mosquitto.mqtt:1883
    topics:
      sourceEvents: sources/mwrsctrl/clusters/+/sourceevents
      agentEvents: sources/mwrsctrl/clusters/+/agentevents
      sourceBroadcast: sources/mwrsctrl/sourcebroadcast
EOF
```

### ManifestworkReplicaset Controller

#### Create ManagedClusterSetBinding

```bash
clusteradm clusterset bind global --namespace default
```
or

```bash
cat << EOF | oc apply -n default -f -
apiVersion: cluster.open-cluster-management.io/v1beta2
kind: ManagedClusterSetBinding
metadata:
  name: global-binding
  namespace: default
spec:
  clusterSet: global
```

#### Create Placement

```bash
cat << EOF | oc apply -n default -f -
apiVersion: cluster.open-cluster-management.io/v1beta1
kind: Placement
metadata:
  name: placement1
  namespace: default
spec:
  numberOfClusters: 3
  clusterSets:
  - global
EOF
```

#### Create ManifestWorkReplicaSet

```bash
cat << EOF | oc apply -n default -f -
apiVersion: work.open-cluster-management.io/v1alpha1
kind: ManifestWorkReplicaSet
metadata:
  name: mwrset-nginx
spec:
  placementRefs:
    - name: placement1
  manifestWorkTemplate:
    workload:
      manifests:
      - apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: nginx
          namespace: default
        spec:
          replicas: 1
          selector:
            matchLabels:
              app: nginx
          template:
            metadata:
              labels:
                app: nginx
            spec:
              containers:
              - image: nginxinc/nginx-unprivileged
                name: nginx
EOF
```

#### Clean up

```bash
oc delete ManifestWorkReplicaSet mwrset-nginx -n default
```

#### Debug

```bash
oc get ManifestWorkReplicaSet -A
oc get appliedmanifestwork -A
oc -n default get pod
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources | jq
curl -k -X GET -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources | jq -r '.items[0].id'
curl -k -X DELETE -H "Content-Type: application/json" https://maestro.${external_apps_domain}/api/maestro/v1/resources/c4df9ff0-bfeb-5bc6-a0ab-4c9128d698b4
oc exec -it maestro-db-748dc568f4-p9pzd -- psql -d maestro -U maestro
delete from resources;
delete from events;
```
