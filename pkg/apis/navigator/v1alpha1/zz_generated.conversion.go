// +build !ignore_autogenerated

/*
Copyright 2017 Jetstack Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	navigator "github.com/jetstack/navigator/pkg/apis/navigator"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_CassandraCluster_To_navigator_CassandraCluster,
		Convert_navigator_CassandraCluster_To_v1alpha1_CassandraCluster,
		Convert_v1alpha1_CassandraClusterList_To_navigator_CassandraClusterList,
		Convert_navigator_CassandraClusterList_To_v1alpha1_CassandraClusterList,
		Convert_v1alpha1_CassandraClusterNodePool_To_navigator_CassandraClusterNodePool,
		Convert_navigator_CassandraClusterNodePool_To_v1alpha1_CassandraClusterNodePool,
		Convert_v1alpha1_CassandraClusterNodePoolStatus_To_navigator_CassandraClusterNodePoolStatus,
		Convert_navigator_CassandraClusterNodePoolStatus_To_v1alpha1_CassandraClusterNodePoolStatus,
		Convert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec,
		Convert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec,
		Convert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus,
		Convert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus,
		Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster,
		Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster,
		Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList,
		Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList,
		Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool,
		Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool,
		Convert_v1alpha1_ElasticsearchClusterNodePoolStatus_To_navigator_ElasticsearchClusterNodePoolStatus,
		Convert_navigator_ElasticsearchClusterNodePoolStatus_To_v1alpha1_ElasticsearchClusterNodePoolStatus,
		Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig,
		Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig,
		Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec,
		Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec,
		Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus,
		Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus,
		Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage,
		Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage,
		Convert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage,
		Convert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage,
		Convert_v1alpha1_ElasticsearchPilotStatus_To_navigator_ElasticsearchPilotStatus,
		Convert_navigator_ElasticsearchPilotStatus_To_v1alpha1_ElasticsearchPilotStatus,
		Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec,
		Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec,
		Convert_v1alpha1_Pilot_To_navigator_Pilot,
		Convert_navigator_Pilot_To_v1alpha1_Pilot,
		Convert_v1alpha1_PilotCondition_To_navigator_PilotCondition,
		Convert_navigator_PilotCondition_To_v1alpha1_PilotCondition,
		Convert_v1alpha1_PilotElasticsearchSpec_To_navigator_PilotElasticsearchSpec,
		Convert_navigator_PilotElasticsearchSpec_To_v1alpha1_PilotElasticsearchSpec,
		Convert_v1alpha1_PilotList_To_navigator_PilotList,
		Convert_navigator_PilotList_To_v1alpha1_PilotList,
		Convert_v1alpha1_PilotSpec_To_navigator_PilotSpec,
		Convert_navigator_PilotSpec_To_v1alpha1_PilotSpec,
		Convert_v1alpha1_PilotStatus_To_navigator_PilotStatus,
		Convert_navigator_PilotStatus_To_v1alpha1_PilotStatus,
	)
}

func autoConvert_v1alpha1_CassandraCluster_To_navigator_CassandraCluster(in *CassandraCluster, out *navigator.CassandraCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CassandraCluster_To_navigator_CassandraCluster is an autogenerated conversion function.
func Convert_v1alpha1_CassandraCluster_To_navigator_CassandraCluster(in *CassandraCluster, out *navigator.CassandraCluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraCluster_To_navigator_CassandraCluster(in, out, s)
}

func autoConvert_navigator_CassandraCluster_To_v1alpha1_CassandraCluster(in *navigator.CassandraCluster, out *CassandraCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_CassandraCluster_To_v1alpha1_CassandraCluster is an autogenerated conversion function.
func Convert_navigator_CassandraCluster_To_v1alpha1_CassandraCluster(in *navigator.CassandraCluster, out *CassandraCluster, s conversion.Scope) error {
	return autoConvert_navigator_CassandraCluster_To_v1alpha1_CassandraCluster(in, out, s)
}

func autoConvert_v1alpha1_CassandraClusterList_To_navigator_CassandraClusterList(in *CassandraClusterList, out *navigator.CassandraClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]navigator.CassandraCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CassandraClusterList_To_navigator_CassandraClusterList is an autogenerated conversion function.
func Convert_v1alpha1_CassandraClusterList_To_navigator_CassandraClusterList(in *CassandraClusterList, out *navigator.CassandraClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraClusterList_To_navigator_CassandraClusterList(in, out, s)
}

func autoConvert_navigator_CassandraClusterList_To_v1alpha1_CassandraClusterList(in *navigator.CassandraClusterList, out *CassandraClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]CassandraCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_navigator_CassandraClusterList_To_v1alpha1_CassandraClusterList is an autogenerated conversion function.
func Convert_navigator_CassandraClusterList_To_v1alpha1_CassandraClusterList(in *navigator.CassandraClusterList, out *CassandraClusterList, s conversion.Scope) error {
	return autoConvert_navigator_CassandraClusterList_To_v1alpha1_CassandraClusterList(in, out, s)
}

func autoConvert_v1alpha1_CassandraClusterNodePool_To_navigator_CassandraClusterNodePool(in *CassandraClusterNodePool, out *navigator.CassandraClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1alpha1_CassandraClusterNodePool_To_navigator_CassandraClusterNodePool is an autogenerated conversion function.
func Convert_v1alpha1_CassandraClusterNodePool_To_navigator_CassandraClusterNodePool(in *CassandraClusterNodePool, out *navigator.CassandraClusterNodePool, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraClusterNodePool_To_navigator_CassandraClusterNodePool(in, out, s)
}

func autoConvert_navigator_CassandraClusterNodePool_To_v1alpha1_CassandraClusterNodePool(in *navigator.CassandraClusterNodePool, out *CassandraClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	return nil
}

// Convert_navigator_CassandraClusterNodePool_To_v1alpha1_CassandraClusterNodePool is an autogenerated conversion function.
func Convert_navigator_CassandraClusterNodePool_To_v1alpha1_CassandraClusterNodePool(in *navigator.CassandraClusterNodePool, out *CassandraClusterNodePool, s conversion.Scope) error {
	return autoConvert_navigator_CassandraClusterNodePool_To_v1alpha1_CassandraClusterNodePool(in, out, s)
}

func autoConvert_v1alpha1_CassandraClusterNodePoolStatus_To_navigator_CassandraClusterNodePoolStatus(in *CassandraClusterNodePoolStatus, out *navigator.CassandraClusterNodePoolStatus, s conversion.Scope) error {
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}

// Convert_v1alpha1_CassandraClusterNodePoolStatus_To_navigator_CassandraClusterNodePoolStatus is an autogenerated conversion function.
func Convert_v1alpha1_CassandraClusterNodePoolStatus_To_navigator_CassandraClusterNodePoolStatus(in *CassandraClusterNodePoolStatus, out *navigator.CassandraClusterNodePoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraClusterNodePoolStatus_To_navigator_CassandraClusterNodePoolStatus(in, out, s)
}

func autoConvert_navigator_CassandraClusterNodePoolStatus_To_v1alpha1_CassandraClusterNodePoolStatus(in *navigator.CassandraClusterNodePoolStatus, out *CassandraClusterNodePoolStatus, s conversion.Scope) error {
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}

// Convert_navigator_CassandraClusterNodePoolStatus_To_v1alpha1_CassandraClusterNodePoolStatus is an autogenerated conversion function.
func Convert_navigator_CassandraClusterNodePoolStatus_To_v1alpha1_CassandraClusterNodePoolStatus(in *navigator.CassandraClusterNodePoolStatus, out *CassandraClusterNodePoolStatus, s conversion.Scope) error {
	return autoConvert_navigator_CassandraClusterNodePoolStatus_To_v1alpha1_CassandraClusterNodePoolStatus(in, out, s)
}

func autoConvert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec(in *CassandraClusterSpec, out *navigator.CassandraClusterSpec, s conversion.Scope) error {
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	out.NodePools = *(*[]navigator.CassandraClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec(&in.PilotImage, &out.PilotImage, s); err != nil {
		return err
	}
	out.CqlPort = in.CqlPort
	return nil
}

// Convert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec(in *CassandraClusterSpec, out *navigator.CassandraClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraClusterSpec_To_navigator_CassandraClusterSpec(in, out, s)
}

func autoConvert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec(in *navigator.CassandraClusterSpec, out *CassandraClusterSpec, s conversion.Scope) error {
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	out.NodePools = *(*[]CassandraClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec(&in.Image, &out.Image, s); err != nil {
		return err
	}
	if err := Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec(&in.PilotImage, &out.PilotImage, s); err != nil {
		return err
	}
	out.CqlPort = in.CqlPort
	return nil
}

// Convert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec is an autogenerated conversion function.
func Convert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec(in *navigator.CassandraClusterSpec, out *CassandraClusterSpec, s conversion.Scope) error {
	return autoConvert_navigator_CassandraClusterSpec_To_v1alpha1_CassandraClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus(in *CassandraClusterStatus, out *navigator.CassandraClusterStatus, s conversion.Scope) error {
	out.NodePools = *(*map[string]navigator.CassandraClusterNodePoolStatus)(unsafe.Pointer(&in.NodePools))
	return nil
}

// Convert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus(in *CassandraClusterStatus, out *navigator.CassandraClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CassandraClusterStatus_To_navigator_CassandraClusterStatus(in, out, s)
}

func autoConvert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus(in *navigator.CassandraClusterStatus, out *CassandraClusterStatus, s conversion.Scope) error {
	out.NodePools = *(*map[string]CassandraClusterNodePoolStatus)(unsafe.Pointer(&in.NodePools))
	return nil
}

// Convert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus is an autogenerated conversion function.
func Convert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus(in *navigator.CassandraClusterStatus, out *CassandraClusterStatus, s conversion.Scope) error {
	return autoConvert_navigator_CassandraClusterStatus_To_v1alpha1_CassandraClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in *ElasticsearchCluster, out *navigator.ElasticsearchCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in *ElasticsearchCluster, out *navigator.ElasticsearchCluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchCluster_To_navigator_ElasticsearchCluster(in, out, s)
}

func autoConvert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in *navigator.ElasticsearchCluster, out *ElasticsearchCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster is an autogenerated conversion function.
func Convert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in *navigator.ElasticsearchCluster, out *ElasticsearchCluster, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchCluster_To_v1alpha1_ElasticsearchCluster(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in *ElasticsearchClusterList, out *navigator.ElasticsearchClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]navigator.ElasticsearchCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in *ElasticsearchClusterList, out *navigator.ElasticsearchClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterList_To_navigator_ElasticsearchClusterList(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in *navigator.ElasticsearchClusterList, out *ElasticsearchClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ElasticsearchCluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in *navigator.ElasticsearchClusterList, out *ElasticsearchClusterList, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterList_To_v1alpha1_ElasticsearchClusterList(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in *ElasticsearchClusterNodePool, out *navigator.ElasticsearchClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]navigator.ElasticsearchClusterRole)(unsafe.Pointer(&in.Roles))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	if err := Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(&in.Persistence, &out.Persistence, s); err != nil {
		return err
	}
	out.Config = *(*map[string]string)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in *ElasticsearchClusterNodePool, out *navigator.ElasticsearchClusterNodePool, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterNodePool_To_navigator_ElasticsearchClusterNodePool(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in *navigator.ElasticsearchClusterNodePool, out *ElasticsearchClusterNodePool, s conversion.Scope) error {
	out.Name = in.Name
	out.Replicas = in.Replicas
	out.Roles = *(*[]ElasticsearchClusterRole)(unsafe.Pointer(&in.Roles))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	if err := Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(&in.Persistence, &out.Persistence, s); err != nil {
		return err
	}
	out.Config = *(*map[string]string)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in *navigator.ElasticsearchClusterNodePool, out *ElasticsearchClusterNodePool, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterNodePool_To_v1alpha1_ElasticsearchClusterNodePool(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterNodePoolStatus_To_navigator_ElasticsearchClusterNodePoolStatus(in *ElasticsearchClusterNodePoolStatus, out *navigator.ElasticsearchClusterNodePoolStatus, s conversion.Scope) error {
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterNodePoolStatus_To_navigator_ElasticsearchClusterNodePoolStatus is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterNodePoolStatus_To_navigator_ElasticsearchClusterNodePoolStatus(in *ElasticsearchClusterNodePoolStatus, out *navigator.ElasticsearchClusterNodePoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterNodePoolStatus_To_navigator_ElasticsearchClusterNodePoolStatus(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterNodePoolStatus_To_v1alpha1_ElasticsearchClusterNodePoolStatus(in *navigator.ElasticsearchClusterNodePoolStatus, out *ElasticsearchClusterNodePoolStatus, s conversion.Scope) error {
	out.ReadyReplicas = in.ReadyReplicas
	return nil
}

// Convert_navigator_ElasticsearchClusterNodePoolStatus_To_v1alpha1_ElasticsearchClusterNodePoolStatus is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterNodePoolStatus_To_v1alpha1_ElasticsearchClusterNodePoolStatus(in *navigator.ElasticsearchClusterNodePoolStatus, out *ElasticsearchClusterNodePoolStatus, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterNodePoolStatus_To_v1alpha1_ElasticsearchClusterNodePoolStatus(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in *ElasticsearchClusterPersistenceConfig, out *navigator.ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in *ElasticsearchClusterPersistenceConfig, out *navigator.ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterPersistenceConfig_To_navigator_ElasticsearchClusterPersistenceConfig(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in *navigator.ElasticsearchClusterPersistenceConfig, out *ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Size = in.Size
	out.StorageClass = in.StorageClass
	return nil
}

// Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in *navigator.ElasticsearchClusterPersistenceConfig, out *ElasticsearchClusterPersistenceConfig, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterPersistenceConfig_To_v1alpha1_ElasticsearchClusterPersistenceConfig(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in *ElasticsearchClusterSpec, out *navigator.ElasticsearchClusterSpec, s conversion.Scope) error {
	out.Plugins = *(*[]string)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]navigator.ElasticsearchClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage(&in.Pilot, &out.Pilot, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in *ElasticsearchClusterSpec, out *navigator.ElasticsearchClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterSpec_To_navigator_ElasticsearchClusterSpec(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in *navigator.ElasticsearchClusterSpec, out *ElasticsearchClusterSpec, s conversion.Scope) error {
	out.Plugins = *(*[]string)(unsafe.Pointer(&in.Plugins))
	out.NodePools = *(*[]ElasticsearchClusterNodePool)(unsafe.Pointer(&in.NodePools))
	if err := Convert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage(&in.Pilot, &out.Pilot, s); err != nil {
		return err
	}
	if err := Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(&in.Image, &out.Image, s); err != nil {
		return err
	}
	out.Sysctl = *(*[]string)(unsafe.Pointer(&in.Sysctl))
	return nil
}

// Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in *navigator.ElasticsearchClusterSpec, out *ElasticsearchClusterSpec, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterSpec_To_v1alpha1_ElasticsearchClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in *ElasticsearchClusterStatus, out *navigator.ElasticsearchClusterStatus, s conversion.Scope) error {
	out.NodePools = *(*map[string]navigator.ElasticsearchClusterNodePoolStatus)(unsafe.Pointer(&in.NodePools))
	return nil
}

// Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in *ElasticsearchClusterStatus, out *navigator.ElasticsearchClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchClusterStatus_To_navigator_ElasticsearchClusterStatus(in, out, s)
}

func autoConvert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in *navigator.ElasticsearchClusterStatus, out *ElasticsearchClusterStatus, s conversion.Scope) error {
	out.NodePools = *(*map[string]ElasticsearchClusterNodePoolStatus)(unsafe.Pointer(&in.NodePools))
	return nil
}

// Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus is an autogenerated conversion function.
func Convert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in *navigator.ElasticsearchClusterStatus, out *ElasticsearchClusterStatus, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchClusterStatus_To_v1alpha1_ElasticsearchClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in *ElasticsearchImage, out *navigator.ElasticsearchImage, s conversion.Scope) error {
	if err := Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec(&in.ImageSpec, &out.ImageSpec, s); err != nil {
		return err
	}
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in *ElasticsearchImage, out *navigator.ElasticsearchImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchImage_To_navigator_ElasticsearchImage(in, out, s)
}

func autoConvert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in *navigator.ElasticsearchImage, out *ElasticsearchImage, s conversion.Scope) error {
	if err := Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec(&in.ImageSpec, &out.ImageSpec, s); err != nil {
		return err
	}
	out.FsGroup = in.FsGroup
	return nil
}

// Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage is an autogenerated conversion function.
func Convert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in *navigator.ElasticsearchImage, out *ElasticsearchImage, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchImage_To_v1alpha1_ElasticsearchImage(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage(in *ElasticsearchPilotImage, out *navigator.ElasticsearchPilotImage, s conversion.Scope) error {
	if err := Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec(&in.ImageSpec, &out.ImageSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage(in *ElasticsearchPilotImage, out *navigator.ElasticsearchPilotImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchPilotImage_To_navigator_ElasticsearchPilotImage(in, out, s)
}

func autoConvert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage(in *navigator.ElasticsearchPilotImage, out *ElasticsearchPilotImage, s conversion.Scope) error {
	if err := Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec(&in.ImageSpec, &out.ImageSpec, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage is an autogenerated conversion function.
func Convert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage(in *navigator.ElasticsearchPilotImage, out *ElasticsearchPilotImage, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchPilotImage_To_v1alpha1_ElasticsearchPilotImage(in, out, s)
}

func autoConvert_v1alpha1_ElasticsearchPilotStatus_To_navigator_ElasticsearchPilotStatus(in *ElasticsearchPilotStatus, out *navigator.ElasticsearchPilotStatus, s conversion.Scope) error {
	out.Documents = (*int64)(unsafe.Pointer(in.Documents))
	return nil
}

// Convert_v1alpha1_ElasticsearchPilotStatus_To_navigator_ElasticsearchPilotStatus is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchPilotStatus_To_navigator_ElasticsearchPilotStatus(in *ElasticsearchPilotStatus, out *navigator.ElasticsearchPilotStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchPilotStatus_To_navigator_ElasticsearchPilotStatus(in, out, s)
}

func autoConvert_navigator_ElasticsearchPilotStatus_To_v1alpha1_ElasticsearchPilotStatus(in *navigator.ElasticsearchPilotStatus, out *ElasticsearchPilotStatus, s conversion.Scope) error {
	out.Documents = (*int64)(unsafe.Pointer(in.Documents))
	return nil
}

// Convert_navigator_ElasticsearchPilotStatus_To_v1alpha1_ElasticsearchPilotStatus is an autogenerated conversion function.
func Convert_navigator_ElasticsearchPilotStatus_To_v1alpha1_ElasticsearchPilotStatus(in *navigator.ElasticsearchPilotStatus, out *ElasticsearchPilotStatus, s conversion.Scope) error {
	return autoConvert_navigator_ElasticsearchPilotStatus_To_v1alpha1_ElasticsearchPilotStatus(in, out, s)
}

func autoConvert_v1alpha1_ImageSpec_To_navigator_ImageSpec(in *ImageSpec, out *navigator.ImageSpec, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	return nil
}

// Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec is an autogenerated conversion function.
func Convert_v1alpha1_ImageSpec_To_navigator_ImageSpec(in *ImageSpec, out *navigator.ImageSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ImageSpec_To_navigator_ImageSpec(in, out, s)
}

func autoConvert_navigator_ImageSpec_To_v1alpha1_ImageSpec(in *navigator.ImageSpec, out *ImageSpec, s conversion.Scope) error {
	out.Repository = in.Repository
	out.Tag = in.Tag
	out.PullPolicy = in.PullPolicy
	return nil
}

// Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec is an autogenerated conversion function.
func Convert_navigator_ImageSpec_To_v1alpha1_ImageSpec(in *navigator.ImageSpec, out *ImageSpec, s conversion.Scope) error {
	return autoConvert_navigator_ImageSpec_To_v1alpha1_ImageSpec(in, out, s)
}

func autoConvert_v1alpha1_Pilot_To_navigator_Pilot(in *Pilot, out *navigator.Pilot, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PilotSpec_To_navigator_PilotSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PilotStatus_To_navigator_PilotStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Pilot_To_navigator_Pilot is an autogenerated conversion function.
func Convert_v1alpha1_Pilot_To_navigator_Pilot(in *Pilot, out *navigator.Pilot, s conversion.Scope) error {
	return autoConvert_v1alpha1_Pilot_To_navigator_Pilot(in, out, s)
}

func autoConvert_navigator_Pilot_To_v1alpha1_Pilot(in *navigator.Pilot, out *Pilot, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_navigator_PilotSpec_To_v1alpha1_PilotSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_navigator_PilotStatus_To_v1alpha1_PilotStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_navigator_Pilot_To_v1alpha1_Pilot is an autogenerated conversion function.
func Convert_navigator_Pilot_To_v1alpha1_Pilot(in *navigator.Pilot, out *Pilot, s conversion.Scope) error {
	return autoConvert_navigator_Pilot_To_v1alpha1_Pilot(in, out, s)
}

func autoConvert_v1alpha1_PilotCondition_To_navigator_PilotCondition(in *PilotCondition, out *navigator.PilotCondition, s conversion.Scope) error {
	out.Type = navigator.PilotConditionType(in.Type)
	out.Status = navigator.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_PilotCondition_To_navigator_PilotCondition is an autogenerated conversion function.
func Convert_v1alpha1_PilotCondition_To_navigator_PilotCondition(in *PilotCondition, out *navigator.PilotCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_PilotCondition_To_navigator_PilotCondition(in, out, s)
}

func autoConvert_navigator_PilotCondition_To_v1alpha1_PilotCondition(in *navigator.PilotCondition, out *PilotCondition, s conversion.Scope) error {
	out.Type = PilotConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_navigator_PilotCondition_To_v1alpha1_PilotCondition is an autogenerated conversion function.
func Convert_navigator_PilotCondition_To_v1alpha1_PilotCondition(in *navigator.PilotCondition, out *PilotCondition, s conversion.Scope) error {
	return autoConvert_navigator_PilotCondition_To_v1alpha1_PilotCondition(in, out, s)
}

func autoConvert_v1alpha1_PilotElasticsearchSpec_To_navigator_PilotElasticsearchSpec(in *PilotElasticsearchSpec, out *navigator.PilotElasticsearchSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_PilotElasticsearchSpec_To_navigator_PilotElasticsearchSpec is an autogenerated conversion function.
func Convert_v1alpha1_PilotElasticsearchSpec_To_navigator_PilotElasticsearchSpec(in *PilotElasticsearchSpec, out *navigator.PilotElasticsearchSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PilotElasticsearchSpec_To_navigator_PilotElasticsearchSpec(in, out, s)
}

func autoConvert_navigator_PilotElasticsearchSpec_To_v1alpha1_PilotElasticsearchSpec(in *navigator.PilotElasticsearchSpec, out *PilotElasticsearchSpec, s conversion.Scope) error {
	return nil
}

// Convert_navigator_PilotElasticsearchSpec_To_v1alpha1_PilotElasticsearchSpec is an autogenerated conversion function.
func Convert_navigator_PilotElasticsearchSpec_To_v1alpha1_PilotElasticsearchSpec(in *navigator.PilotElasticsearchSpec, out *PilotElasticsearchSpec, s conversion.Scope) error {
	return autoConvert_navigator_PilotElasticsearchSpec_To_v1alpha1_PilotElasticsearchSpec(in, out, s)
}

func autoConvert_v1alpha1_PilotList_To_navigator_PilotList(in *PilotList, out *navigator.PilotList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]navigator.Pilot)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PilotList_To_navigator_PilotList is an autogenerated conversion function.
func Convert_v1alpha1_PilotList_To_navigator_PilotList(in *PilotList, out *navigator.PilotList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PilotList_To_navigator_PilotList(in, out, s)
}

func autoConvert_navigator_PilotList_To_v1alpha1_PilotList(in *navigator.PilotList, out *PilotList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Pilot)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_navigator_PilotList_To_v1alpha1_PilotList is an autogenerated conversion function.
func Convert_navigator_PilotList_To_v1alpha1_PilotList(in *navigator.PilotList, out *PilotList, s conversion.Scope) error {
	return autoConvert_navigator_PilotList_To_v1alpha1_PilotList(in, out, s)
}

func autoConvert_v1alpha1_PilotSpec_To_navigator_PilotSpec(in *PilotSpec, out *navigator.PilotSpec, s conversion.Scope) error {
	out.Elasticsearch = (*navigator.PilotElasticsearchSpec)(unsafe.Pointer(in.Elasticsearch))
	return nil
}

// Convert_v1alpha1_PilotSpec_To_navigator_PilotSpec is an autogenerated conversion function.
func Convert_v1alpha1_PilotSpec_To_navigator_PilotSpec(in *PilotSpec, out *navigator.PilotSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PilotSpec_To_navigator_PilotSpec(in, out, s)
}

func autoConvert_navigator_PilotSpec_To_v1alpha1_PilotSpec(in *navigator.PilotSpec, out *PilotSpec, s conversion.Scope) error {
	out.Elasticsearch = (*PilotElasticsearchSpec)(unsafe.Pointer(in.Elasticsearch))
	return nil
}

// Convert_navigator_PilotSpec_To_v1alpha1_PilotSpec is an autogenerated conversion function.
func Convert_navigator_PilotSpec_To_v1alpha1_PilotSpec(in *navigator.PilotSpec, out *PilotSpec, s conversion.Scope) error {
	return autoConvert_navigator_PilotSpec_To_v1alpha1_PilotSpec(in, out, s)
}

func autoConvert_v1alpha1_PilotStatus_To_navigator_PilotStatus(in *PilotStatus, out *navigator.PilotStatus, s conversion.Scope) error {
	out.LastCompletedPhase = navigator.PilotPhase(in.LastCompletedPhase)
	out.Conditions = *(*[]navigator.PilotCondition)(unsafe.Pointer(&in.Conditions))
	out.Elasticsearch = (*navigator.ElasticsearchPilotStatus)(unsafe.Pointer(in.Elasticsearch))
	return nil
}

// Convert_v1alpha1_PilotStatus_To_navigator_PilotStatus is an autogenerated conversion function.
func Convert_v1alpha1_PilotStatus_To_navigator_PilotStatus(in *PilotStatus, out *navigator.PilotStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PilotStatus_To_navigator_PilotStatus(in, out, s)
}

func autoConvert_navigator_PilotStatus_To_v1alpha1_PilotStatus(in *navigator.PilotStatus, out *PilotStatus, s conversion.Scope) error {
	out.LastCompletedPhase = PilotPhase(in.LastCompletedPhase)
	out.Conditions = *(*[]PilotCondition)(unsafe.Pointer(&in.Conditions))
	out.Elasticsearch = (*ElasticsearchPilotStatus)(unsafe.Pointer(in.Elasticsearch))
	return nil
}

// Convert_navigator_PilotStatus_To_v1alpha1_PilotStatus is an autogenerated conversion function.
func Convert_navigator_PilotStatus_To_v1alpha1_PilotStatus(in *navigator.PilotStatus, out *PilotStatus, s conversion.Scope) error {
	return autoConvert_navigator_PilotStatus_To_v1alpha1_PilotStatus(in, out, s)
}
