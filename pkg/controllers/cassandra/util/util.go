package util

import (
	"fmt"
	"strings"

	"k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	appslisters "k8s.io/client-go/listers/apps/v1beta1"

	"github.com/jetstack/navigator/pkg/apis/navigator"
	v1alpha1 "github.com/jetstack/navigator/pkg/apis/navigator/v1alpha1"
)

const (
	typeName             = "cass"
	kindName             = "CassandraCluster"
	ClusterNameLabelKey  = "navigator.jetstack.io/cassandra-cluster-name"
	NodePoolNameLabelKey = "navigator.jetstack.io/cassandra-node-pool-name"
)

func NewControllerRef(c metav1.Object) metav1.OwnerReference {
	return *metav1.NewControllerRef(c, schema.GroupVersionKind{
		Group:   navigator.GroupName,
		Version: "v1alpha1",
		Kind:    kindName,
	})
}

func ResourceBaseName(c *v1alpha1.CassandraCluster) string {
	return typeName + "-" + c.Name
}

func NodePoolResourceName(c *v1alpha1.CassandraCluster, np *v1alpha1.CassandraClusterNodePool) string {
	return fmt.Sprintf("%s-%s", ResourceBaseName(c), np.Name)
}

func ParseNodePoolResourceName(name string) (string, string, error) {
	err := fmt.Errorf(
		"not a NodePool statefulset name. "+
			"Expected 'cass-<clustername>-<nodepoolname>'. "+
			"Got %q.",
		name,
	)
	parts := strings.Split(name, "-")
	partsCount := len(parts)
	if partsCount < 3 {
		return "", "", err
	}
	if parts[0] != typeName {
		return "", "", err
	}
	clusterName := strings.Join(parts[1:len(parts)-1], "-")
	nodePoolName := parts[len(parts)-1]

	return clusterName, nodePoolName, nil
}

func SeedsServiceName(c *v1alpha1.CassandraCluster) string {
	return fmt.Sprintf("%s-seeds", ResourceBaseName(c))
}

func ServiceAccountName(c *v1alpha1.CassandraCluster) string {
	return ResourceBaseName(c)
}

func PilotRBACRoleName(c *v1alpha1.CassandraCluster) string {
	return fmt.Sprintf("%s-pilot", ResourceBaseName(c))
}

func ClusterLabels(c metav1.Object) map[string]string {
	return map[string]string{
		"app":               "cassandracluster",
		ClusterNameLabelKey: c.GetName(),
	}
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

func NodePoolLabels(
	c *v1alpha1.CassandraCluster,
	poolName string,
) map[string]string {
	labels := ClusterLabels(c)
	labels[NodePoolNameLabelKey] = poolName
	return labels
}

func OwnerCheck(
	obj metav1.Object,
	owner metav1.Object,
) error {
	if !metav1.IsControlledBy(obj, owner) {
		ownerRef := metav1.GetControllerOf(obj)
		return fmt.Errorf(
			"'%s/%s' is foreign owned: "+
				"it is owned by '%v', not '%s/%s'.",
			obj.GetNamespace(), obj.GetName(),
			ownerRef,
			owner.GetNamespace(), owner.GetName(),
		)
	}
	return nil
}

func StatefulSetsForCluster(
	cluster *v1alpha1.CassandraCluster,
	statefulSetLister appslisters.StatefulSetLister,
) (results map[string]*v1beta1.StatefulSet, err error) {
	results = map[string]*v1beta1.StatefulSet{}
	lister := statefulSetLister.StatefulSets(cluster.Namespace)
	selector, err := SelectorForCluster(cluster)
	if err != nil {
		return nil, err
	}
	existingSets, err := lister.List(selector)
	if err != nil {
		return nil, err
	}
	for _, set := range existingSets {
		err := OwnerCheck(set, cluster)
		if err != nil {
			continue
		}
		results[set.Name] = set
	}
	return results, nil
}
