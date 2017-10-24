package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"

	"github.com/jetstack-experimental/navigator/pkg/apis/navigator"
	v1alpha1 "github.com/jetstack-experimental/navigator/pkg/apis/navigator/v1alpha1"
)

const (
	typeName = "cass"
	kindName = "CassandraCluster"
)

const (
	ClusterNameLabelKey = "navigator.jetstack.io/cassandra-cluster-name"
)

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func NewControllerRef(c *v1alpha1.CassandraCluster) metav1.OwnerReference {
	return *metav1.NewControllerRef(c, schema.GroupVersionKind{
		Group:   navigator.GroupName,
		Version: "v1alpha1",
		Kind:    kindName,
	})
}

func ResourceBaseName(c *v1alpha1.CassandraCluster) string {
	return typeName + "-" + c.Name
}

func SelectorForCluster(c *v1alpha1.CassandraCluster) (labels.Selector, error) {
	clusterNameReq, err := labels.NewRequirement(
		ClusterNameLabelKey,
		selection.Equals,
		[]string{c.Name},
	)
	if err != nil {
		return nil, err
	}
	return labels.NewSelector().Add(*clusterNameReq), nil
}

func ClusterLabels(c *v1alpha1.CassandraCluster) map[string]string {
	return map[string]string{
		"app":               "cassandracluster",
		ClusterNameLabelKey: c.Name,
	}
}
