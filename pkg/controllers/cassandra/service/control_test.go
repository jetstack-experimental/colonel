package service_test

import (
	"testing"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/jetstack-experimental/navigator/pkg/apis/navigator/v1alpha1"
	"github.com/jetstack-experimental/navigator/pkg/controllers/cassandra/service"
)

type fixture struct {
	t       *testing.T
	control service.Interface
	kclient *fake.Clientset
}

func newFixture(t *testing.T) *fixture {
	kclient := fake.NewSimpleClientset()
	factory := informers.NewSharedInformerFactory(kclient, 0)
	serviceLister := factory.Core().V1().Services().Lister()
	control := service.NewControl(kclient, serviceLister)

	return &fixture{
		t:       t,
		control: control,
		kclient: kclient,
	}
}

func (f *fixture) run(cluster *v1alpha1.CassandraCluster) {
	err := f.control.Sync(cluster)
	if err != nil {
		f.t.Error(err)
	}
}

func (f *fixture) expectService(service *apiv1.Service) {
	_, err := f.kclient.CoreV1().Services(service.Namespace).Get(
		service.Name,
		metav1.GetOptions{},
	)
	if err != nil {
		f.t.Log("Actions:")
		f.t.Log(f.kclient.Actions())
		f.t.Error(err)
	}
}

func newCassandraCluster() *v1alpha1.CassandraCluster {
	c := &v1alpha1.CassandraCluster{}
	c.SetNamespace("foo")
	c.SetName("bar")
	return c
}

func TestServiceControl(t *testing.T) {
	t.Run(
		"service created",
		func(t *testing.T) {
			cluster := newCassandraCluster()
			expectedService := service.ServiceForCluster(cluster)
			f := newFixture(t)
			f.run(cluster)
			f.expectService(expectedService)
		},
	)
	t.Run(
		"resync",
		func(t *testing.T) {
			cluster := newCassandraCluster()
			expectedService := service.ServiceForCluster(cluster)
			f := newFixture(t)
			f.run(cluster)
			f.expectService(expectedService)
			f.run(cluster)
			f.expectService(expectedService)
		},
	)
}
