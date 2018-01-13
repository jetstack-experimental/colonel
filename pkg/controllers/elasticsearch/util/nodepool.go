package util

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"

	"github.com/coreos/go-semver/semver"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"

	v1alpha1 "github.com/jetstack/navigator/pkg/apis/navigator/v1alpha1"
	hashutil "github.com/jetstack/navigator/pkg/util/hash"
)

const (
	NodePoolNameLabelKey      = "navigator.jetstack.io/elasticsearch-node-pool-name"
	NodePoolHashAnnotationKey = "navigator.jetstack.io/elasticsearch-node-pool-hash"
)

// ComputeHash returns a hash value calculated from pod template and a collisionCount to avoid hash collision
func ComputeNodePoolHash(c *v1alpha1.ElasticsearchCluster, np *v1alpha1.ElasticsearchClusterNodePool, collisionCount *int32) string {
	hashVar := struct {
		Plugins    []string
		ESImage    v1alpha1.ElasticsearchImage
		PilotImage v1alpha1.ElasticsearchPilotImage
		Sysctl     []string
		NodePool   *v1alpha1.ElasticsearchClusterNodePool
	}{
		Plugins:    c.Spec.Plugins,
		ESImage:    c.Spec.Image,
		PilotImage: c.Spec.Pilot,
		Sysctl:     c.Spec.Sysctl,
		NodePool:   np,
	}

	hasher := fnv.New32a()
	hashutil.DeepHashObject(hasher, hashVar)

	// Add collisionCount in the hash if it exists.
	if collisionCount != nil {
		collisionCountBytes := make([]byte, 8)
		binary.LittleEndian.PutUint32(collisionCountBytes, uint32(*collisionCount))
		hasher.Write(collisionCountBytes)
	}

	return fmt.Sprintf("%d", hasher.Sum32())
}

func ClusterLabels(c *v1alpha1.ElasticsearchCluster) map[string]string {
	return map[string]string{
		"app":               "elasticsearch",
		ClusterNameLabelKey: c.Name,
	}
}

func NodePoolLabels(c *v1alpha1.ElasticsearchCluster, poolName string, roles ...v1alpha1.ElasticsearchClusterRole) map[string]string {
	labels := ClusterLabels(c)
	if poolName != "" {
		labels[NodePoolNameLabelKey] = poolName
	}
	for _, role := range roles {
		labels[string(role)] = "true"
	}
	return labels
}

func NodePoolResourceName(c *v1alpha1.ElasticsearchCluster, np *v1alpha1.ElasticsearchClusterNodePool) string {
	return fmt.Sprintf("%s-%s", ResourceBaseName(c), np.Name)
}

func SelectorForNodePool(c *v1alpha1.ElasticsearchCluster, np *v1alpha1.ElasticsearchClusterNodePool) (labels.Selector, error) {
	nodePoolNameReq, err := labels.NewRequirement(NodePoolNameLabelKey, selection.Equals, []string{np.Name})
	if err != nil {
		return nil, err
	}
	clusterSelector, err := SelectorForCluster(c)
	if err != nil {
		return nil, err
	}
	return clusterSelector.Add(*nodePoolNameReq), nil
}

const defaultElasticsearchImageRepository = "docker.elastic.co/elasticsearch/elasticsearch"
const defaultElasticsearchRunAsUser = 1000

func DefaultElasticsearchImageForVersion(s string) (v1alpha1.ElasticsearchImage, error) {
	// ensure the version follows semver
	_, err := semver.NewVersion(s)
	if err != nil {
		return v1alpha1.ElasticsearchImage{}, err
	}

	return v1alpha1.ElasticsearchImage{
		FsGroup: defaultElasticsearchRunAsUser,
		ImageSpec: v1alpha1.ImageSpec{
			Repository: defaultElasticsearchImageRepository,
			Tag:        s,
			PullPolicy: string(corev1.PullIfNotPresent),
		},
	}, nil
}
