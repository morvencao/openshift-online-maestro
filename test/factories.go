package test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/api/openapi"
	"github.com/openshift-online/maestro/pkg/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gorm.io/datatypes"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	workv1 "open-cluster-management.io/api/work/v1"
	cetypes "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/types"
	workpayload "open-cluster-management.io/sdk-go/pkg/cloudevents/work/payload"
)

var testManifestJSON = `
{
	"apiVersion": "apps/v1",
	"kind": "Deployment",
	"metadata": {
	  "name": "nginx",
	  "namespace": "default"
	},
	"spec": {
	  "replicas": %d,
	  "selector": {
		"matchLabels": {
		  "app": "nginx"
		}
	  },
	  "template": {
		"metadata": {
		  "labels": {
			"app": "nginx"
		  }
		},
		"spec": {
		  "containers": [
			{
			  "image": "nginxinc/nginx-unprivileged",
			  "name": "nginx"
			}
		  ]
		}
	  }
	}
}
`

var testManifestJSONWithSA = `
{
	"apiVersion": "apps/v1",
	"kind": "Deployment",
	"metadata": {
	  "name": "nginx",
	  "namespace": "default"
	},
	"spec": {
	  "replicas": %d,
	  "selector": {
		"matchLabels": {
		  "app": "nginx"
		}
	  },
	  "template": {
		"metadata": {
		  "labels": {
			"app": "nginx"
		  }
		},
		"spec": {
		  "serviceAccount": "%s",
		  "containers": [
			{
			  "image": "nginxinc/nginx-unprivileged",
			  "name": "nginx"
			}
		  ]
		}
	  }
	}
}
`

var testManifestIndexJSON = `
{
	"apiVersion": "apps/v1",
	"kind": "Deployment",
	"metadata": {
	  "name": "nginx-%d",
	  "namespace": "default"
	},
	"spec": {
	  "replicas": %d,
	  "selector": {
		"matchLabels": {
		  "app": "nginx"
		}
	  },
	  "template": {
		"metadata": {
		  "labels": {
			"app": "nginx"
		  }
		},
		"spec": {
		  "containers": [
			{
			  "image": "nginxinc/nginx-unprivileged",
			  "name": "nginx"
			}
		  ]
		}
	  }
	}
}
`

var testReadOnlyManifestJSON = `
{
	"apiVersion": "apps/v1",
	"kind": "Deployment",
	"metadata": {
	  "name": "nginx",
	  "namespace": "default"
	},
	"update_strategy": {
	  "type": "ReadOnly"
	}
}
`

func (helper *Helper) NewAPIResource(consumerName string, replicas int) openapi.Resource {
	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(testManifestJSON, replicas)), &testManifest); err != nil {
		helper.T.Errorf("error unmarshalling test manifest: %q", err)
	}

	return openapi.Resource{
		Manifest:     testManifest,
		ConsumerName: &consumerName,
	}
}

func (helper *Helper) NewAPIResourceWithSA(consumerName string, replicas int, sa string) openapi.Resource {
	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(testManifestJSONWithSA, replicas, sa)), &testManifest); err != nil {
		helper.T.Errorf("error unmarshalling test manifest: %q", err)
	}

	return openapi.Resource{
		Manifest:     testManifest,
		ConsumerName: &consumerName,
	}
}

func (helper *Helper) NewAPIResourceWithIndex(consumerName string, replicas, index int) openapi.Resource {
	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(testManifestIndexJSON, index, replicas)), &testManifest); err != nil {
		helper.T.Errorf("error unmarshalling test manifest: %q", err)
	}

	return openapi.Resource{
		Manifest:     testManifest,
		ConsumerName: &consumerName,
	}
}

func (helper *Helper) GetTestNginxJSON(replicas int) []byte {
	return []byte(fmt.Sprintf(testManifestJSON, replicas))
}

func (helper *Helper) NewReadOnlyAPIResource(consumerName string) openapi.Resource {
	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprint(testReadOnlyManifestJSON)), &testManifest); err != nil {
		helper.T.Errorf("error unmarshalling test manifest: %q", err)
	}

	return openapi.Resource{
		Manifest:     testManifest,
		ConsumerName: &consumerName,
	}
}

func (helper *Helper) NewResource(consumerName string, replicas int) *api.Resource {
	testResource := helper.NewAPIResource(consumerName, replicas)
	testPayload, err := api.EncodeManifest(testResource.Manifest, testResource.DeleteOption, testResource.UpdateStrategy)
	if err != nil {
		helper.T.Errorf("error encoding manifest: %q", err)
	}

	resource := &api.Resource{
		ConsumerName: consumerName,
		Type:         api.ResourceTypeSingle,
		Payload:      testPayload,
		Version:      1,
	}

	return resource
}

func (helper *Helper) CreateResource(consumerName string, replicas int) *api.Resource {
	resourceService := helper.Env().Services.Resources()
	resource := helper.NewResource(consumerName, replicas)

	res, err := resourceService.Create(context.Background(), resource)
	if err != nil {
		helper.T.Errorf("error creating resource: %q", err)
	}

	return res
}

func (helper *Helper) CreateResourceList(consumerName string, count int) (resources []*api.Resource) {
	for i := 1; i <= count; i++ {
		resources = append(resources, helper.CreateResource(consumerName, 1))
	}
	return resources
}

// EncodeManifestBundle converts resource manifests into a CloudEvent JSONMap representation.
func (helper *Helper) EncodeManifestBundle(manifest map[string]interface{}) (datatypes.JSONMap, error) {
	if len(manifest) == 0 {
		return nil, nil
	}

	delOption := &workv1.DeleteOption{
		PropagationPolicy: workv1.DeletePropagationPolicyTypeForeground,
	}

	upStrategy := &workv1.UpdateStrategy{
		Type: workv1.UpdateStrategyTypeServerSideApply,
	}

	// create a cloud event with the manifest as the data
	evt := cetypes.NewEventBuilder("maestro", cetypes.CloudEventsType{}).NewEvent()
	eventPayload := &workpayload.ManifestBundle{
		Manifests: []workv1.Manifest{
			{
				RawExtension: runtime.RawExtension{
					Object: &unstructured.Unstructured{Object: manifest},
				},
			},
		},
		DeleteOption: delOption,
		ManifestConfigs: []workv1.ManifestConfigOption{
			{
				FeedbackRules: []workv1.FeedbackRule{
					{
						Type: workv1.JSONPathsType,
						JsonPaths: []workv1.JsonPath{
							{
								Name: "status",
								Path: ".status",
							},
						},
					},
				},
				UpdateStrategy: upStrategy,
				ResourceIdentifier: workv1.ResourceIdentifier{
					Group:     "apps",
					Resource:  "deployments",
					Name:      "nginx",
					Namespace: "default",
				},
			},
		},
	}

	if err := evt.SetData(cloudevents.ApplicationJSON, eventPayload); err != nil {
		return nil, fmt.Errorf("failed to set cloud event data: %v", err)
	}

	// convert cloudevent to JSONMap
	manifest, err := api.CloudEventToJSONMap(&evt)
	if err != nil {
		return nil, fmt.Errorf("failed to convert cloudevent to resource manifest: %v", err)
	}

	return manifest, nil
}

func (helper *Helper) NewResourceBundle(name, consumerName string, replicas int) *api.Resource {
	testResource := helper.NewAPIResource(consumerName, replicas)
	testPayload, err := helper.EncodeManifestBundle(testResource.Manifest)
	if err != nil {
		helper.T.Errorf("error encoding manifest bundle: %q", err)
	}

	resource := &api.Resource{
		Name:         name,
		ConsumerName: consumerName,
		Type:         api.ResourceTypeBundle,
		Payload:      testPayload,
		Version:      1,
	}

	return resource
}

func (helper *Helper) CreateResourceBundle(name, consumerName string, replicas int) *api.Resource {
	resourceService := helper.Env().Services.Resources()
	resourceBundle := helper.NewResourceBundle(name, consumerName, replicas)

	res, err := resourceService.Create(context.Background(), resourceBundle)
	if err != nil {
		helper.T.Errorf("error creating resource bundle: %q", err)
	}

	return res
}

func (helper *Helper) CreateResourceBundleList(consumerName string, count int) (resources []*api.Resource) {
	for i := 1; i <= count; i++ {
		resources = append(resources, helper.CreateResourceBundle(fmt.Sprintf("resource%d", i), consumerName, 1))
	}
	return resources
}

func (helper *Helper) CreateConsumer(name string) *api.Consumer {
	return helper.CreateConsumerWithLabels(name, nil)
}

func (helper *Helper) CreateConsumerWithLabels(name string, labels map[string]string) *api.Consumer {
	consumerService := helper.Env().Services.Consumers()

	consumer, err := consumerService.Create(context.Background(), &api.Consumer{Name: name, Labels: db.EmptyMapToNilStringMap(&labels)})
	if err != nil {
		helper.T.Errorf("error creating resource: %q", err)
	}
	return consumer
}

func (helper *Helper) CreateConsumerList(count int) (consumers []*api.Consumer) {
	for i := 1; i <= count; i++ {
		consumers = append(consumers, helper.CreateConsumer(fmt.Sprintf("consumer-%d", i)))
	}
	return consumers
}

// ManifestToEvent converts a manifest into a CloudEvent representation with manifest data.
func (helper *Helper) ManifestToEvent(replicas int, source, action, consumerName, resourceID string,
	resourceVersion int64, deleting bool) (*cloudevents.Event, error) {

	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(testManifestJSON, replicas)), &testManifest); err != nil {
		return nil, fmt.Errorf("error unmarshalling test manifest: %v", err)
	}

	eventType := cetypes.CloudEventsType{
		CloudEventsDataType: workpayload.ManifestEventDataType,
		SubResource:         cetypes.SubResourceSpec,
		Action:              cetypes.EventAction(action),
	}
	evtBuilder := cetypes.NewEventBuilder(source, eventType).
		WithClusterName(consumerName).
		WithResourceID(resourceID).
		WithResourceVersion(resourceVersion)
	if deleting {
		evtBuilder.WithDeletionTimestamp(time.Now())
	}
	evt := evtBuilder.NewEvent()

	eventPayload := &workpayload.Manifest{
		Manifest: unstructured.Unstructured{Object: testManifest},
		DeleteOption: &workv1.DeleteOption{
			PropagationPolicy: workv1.DeletePropagationPolicyTypeForeground,
		},
		ConfigOption: &workpayload.ManifestConfigOption{
			FeedbackRules: []workv1.FeedbackRule{
				{
					Type: workv1.JSONPathsType,
					JsonPaths: []workv1.JsonPath{
						{
							Name: "status",
							Path: ".status",
						},
					},
				},
			},
			UpdateStrategy: &workv1.UpdateStrategy{
				Type: workv1.UpdateStrategyTypeServerSideApply,
			},
		},
	}

	if err := evt.SetData(cloudevents.ApplicationJSON, eventPayload); err != nil {
		return nil, fmt.Errorf("failed to set cloud event data: %v", err)
	}

	return &evt, nil
}

// ManifestsToBundleEvent converts a list of manifests into a CloudEvent representation with manifest bundle data.
func (helper *Helper) ManifestsToBundleEvent(replicas int, source, action, consumerName, resourceID string,
	resourceVersion int64, deleting bool) (*cloudevents.Event, error) {

	testManifest := map[string]interface{}{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(testManifestJSON, replicas)), &testManifest); err != nil {
		return nil, fmt.Errorf("error unmarshalling test manifest: %v", err)
	}

	eventType := cetypes.CloudEventsType{
		CloudEventsDataType: workpayload.ManifestBundleEventDataType,
		SubResource:         cetypes.SubResourceSpec,
		Action:              cetypes.EventAction(action),
	}

	// create a cloud event with the manifest as the data
	evtBuilder := cetypes.NewEventBuilder(source, eventType).
		WithClusterName(consumerName).
		WithResourceID(resourceID).
		WithResourceVersion(resourceVersion)
	if deleting {
		evtBuilder.WithDeletionTimestamp(time.Now())
	}
	evt := evtBuilder.NewEvent()

	eventPayload := &workpayload.ManifestBundle{
		Manifests: []workv1.Manifest{
			{
				RawExtension: runtime.RawExtension{
					Object: &unstructured.Unstructured{Object: testManifest},
				},
			},
		},
		DeleteOption: &workv1.DeleteOption{
			PropagationPolicy: workv1.DeletePropagationPolicyTypeForeground,
		},
		ManifestConfigs: []workv1.ManifestConfigOption{
			{
				FeedbackRules: []workv1.FeedbackRule{
					{
						Type: workv1.JSONPathsType,
						JsonPaths: []workv1.JsonPath{
							{
								Name: "status",
								Path: ".status",
							},
						},
					},
				},
				UpdateStrategy: &workv1.UpdateStrategy{
					Type: workv1.UpdateStrategyTypeServerSideApply,
				},
				ResourceIdentifier: workv1.ResourceIdentifier{
					Group:     "apps",
					Resource:  "deployments",
					Name:      "nginx",
					Namespace: "default",
				},
			},
		},
	}

	if err := evt.SetData(cloudevents.ApplicationJSON, eventPayload); err != nil {
		return nil, fmt.Errorf("failed to set cloud event data: %v", err)
	}

	return &evt, nil
}

func (helper *Helper) CreateGRPCAuthRule(ctx context.Context, kubeClient kubernetes.Interface, ruleName, resourceType, resourceID string, actions []string) error {
	// create the cluster rolefor grpc authz
	nonResourceUrl := ""
	switch resourceType {
	case "source":
		nonResourceUrl = fmt.Sprintf("/sources/%s", resourceID)
	case "cluster":
		nonResourceUrl = fmt.Sprintf("/clusters/%s", resourceID)
	default:
		return fmt.Errorf("unsupported resource type: %s", resourceType)
	}

	_, err := kubeClient.RbacV1().ClusterRoles().Create(ctx, &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: ruleName,
		},
		Rules: []rbacv1.PolicyRule{
			{
				NonResourceURLs: []string{nonResourceUrl},
				Verbs:           actions,
			},
		},
	}, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		// update the cluster role
		_, err = kubeClient.RbacV1().ClusterRoles().Update(ctx, &rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name: ruleName,
			},
			Rules: []rbacv1.PolicyRule{
				{
					NonResourceURLs: []string{nonResourceUrl},
					Verbs:           actions,
				},
			},
		}, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	return err
}

func (helper *Helper) CreateGrpcConn(grpcServerAddress, serverCAFile, clientCertFile, clientKeyFile string) (*grpc.ClientConn, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	caPEM, err := os.ReadFile(serverCAFile)
	if err != nil {
		return nil, err
	}

	ok := certPool.AppendCertsFromPEM(caPEM)
	if !ok {
		return nil, fmt.Errorf("failed to append server CA certificate")
	}

	clientCerts, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCerts},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	return grpc.Dial(grpcServerAddress, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
}
