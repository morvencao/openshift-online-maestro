package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/dao"
	"github.com/openshift-online/maestro/test"
	"k8s.io/apimachinery/pkg/util/rand"
)

func TestHealthCheckServer(t *testing.T) {
	h, _ := test.RegisterIntegration(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	instanceDao := dao.NewInstanceDao(&h.Env().Database.SessionFactory)
	// insert one existing instances
	_, err := instanceDao.Create(ctx, &api.ServerInstance{
		Meta: api.Meta{
			ID: "instance1",
		},
		LastHeartbeat: time.Now(),
		Ready:         true,
	})
	Expect(err).NotTo(HaveOccurred())

	// create a consumer
	clusterName := "cluster-" + rand.String(5)
	_ = h.CreateConsumer(clusterName)

	instanceID := &h.Env().Config.MessageBroker.ClientID
	Eventually(func() error {
		instances, err := instanceDao.All(ctx)
		if err != nil {
			return err
		}

		if len(instances) != 2 {
			return fmt.Errorf("expected 1 instance, got %d", len(instances))
		}

		var instance *api.ServerInstance
		for _, i := range instances {
			if i.ID == *instanceID {
				instance = i
			}
		}

		if instance.LastHeartbeat.IsZero() {
			return fmt.Errorf("expected instance.LastHeartbeat to be non-zero")
		}

		if !instance.Ready {
			return fmt.Errorf("expected instance.Ready to be true")
		}

		if instance.ID != *instanceID {
			return fmt.Errorf("expected instance.ID to be %s, got %s", *instanceID, instance.ID)
		}

		return nil
	}, 10*time.Second, 1*time.Second).Should(Succeed())
}
