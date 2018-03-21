package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	cassDefaultDatacenter = "navigator-default-datacenter"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_CassandraClusterNodePool(np *CassandraClusterNodePool) {
	if np.Datacenter == "" {
		np.Datacenter = cassDefaultDatacenter
	}

	if np.Rack == "" {
		np.Rack = np.Name
	}
}
