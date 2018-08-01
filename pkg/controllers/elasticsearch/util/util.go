package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"

	"github.com/jetstack/navigator/pkg/apis/navigator"
	v1alpha1 "github.com/jetstack/navigator/pkg/apis/navigator/v1alpha1"
)

const (
	typeName = "es"
	kindName = "ElasticsearchCluster"
)

func NewControllerRef(c *v1alpha1.ElasticsearchCluster) metav1.OwnerReference {
	return *metav1.NewControllerRef(c, schema.GroupVersionKind{
		Group:   navigator.GroupName,
		Version: "v1alpha1",
		Kind:    kindName,
	})
}

func ResourceBaseName(c *v1alpha1.ElasticsearchCluster) string {
	return typeName + "-" + c.Name
}

func SelectorForCluster(clusterName string) (labels.Selector, error) {
	clusterTypeReq, err := labels.NewRequirement(v1alpha1.ClusterTypeLabel, selection.Equals, []string{kindName})
	if err != nil {
		return nil, err
	}
	clusterNameReq, err := labels.NewRequirement(v1alpha1.ClusterNameLabel, selection.Equals, []string{clusterName})
	if err != nil {
		return nil, err
	}
	return labels.NewSelector().Add(*clusterTypeReq, *clusterNameReq), nil
}
