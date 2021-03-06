/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/kubernetes-sigs/federation-v2/pkg/apis/federation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	federationFederatedClusterStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedCluster,
		FederatedClusterSchemeFns{},
		func() runtime.Object { return &FederatedCluster{} },     // Register versioned resource
		func() runtime.Object { return &FederatedClusterList{} }, // Register versioned resource list
		&FederatedClusterStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedConfigMapStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedConfigMap,
		FederatedConfigMapSchemeFns{},
		func() runtime.Object { return &FederatedConfigMap{} },     // Register versioned resource
		func() runtime.Object { return &FederatedConfigMapList{} }, // Register versioned resource list
		&FederatedConfigMapStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedConfigMapOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedConfigMapOverride,
		FederatedConfigMapOverrideSchemeFns{},
		func() runtime.Object { return &FederatedConfigMapOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedConfigMapOverrideList{} }, // Register versioned resource list
		&FederatedConfigMapOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedConfigMapPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedConfigMapPlacement,
		FederatedConfigMapPlacementSchemeFns{},
		func() runtime.Object { return &FederatedConfigMapPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedConfigMapPlacementList{} }, // Register versioned resource list
		&FederatedConfigMapPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedDeploymentStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedDeployment,
		FederatedDeploymentSchemeFns{},
		func() runtime.Object { return &FederatedDeployment{} },     // Register versioned resource
		func() runtime.Object { return &FederatedDeploymentList{} }, // Register versioned resource list
		&FederatedDeploymentStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedDeploymentOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedDeploymentOverride,
		FederatedDeploymentOverrideSchemeFns{},
		func() runtime.Object { return &FederatedDeploymentOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedDeploymentOverrideList{} }, // Register versioned resource list
		&FederatedDeploymentOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedDeploymentPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedDeploymentPlacement,
		FederatedDeploymentPlacementSchemeFns{},
		func() runtime.Object { return &FederatedDeploymentPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedDeploymentPlacementList{} }, // Register versioned resource list
		&FederatedDeploymentPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedJobStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedJob,
		FederatedJobSchemeFns{},
		func() runtime.Object { return &FederatedJob{} },     // Register versioned resource
		func() runtime.Object { return &FederatedJobList{} }, // Register versioned resource list
		&FederatedJobStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedJobOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedJobOverride,
		FederatedJobOverrideSchemeFns{},
		func() runtime.Object { return &FederatedJobOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedJobOverrideList{} }, // Register versioned resource list
		&FederatedJobOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedJobPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedJobPlacement,
		FederatedJobPlacementSchemeFns{},
		func() runtime.Object { return &FederatedJobPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedJobPlacementList{} }, // Register versioned resource list
		&FederatedJobPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedNamespacePlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedNamespacePlacement,
		FederatedNamespacePlacementSchemeFns{},
		func() runtime.Object { return &FederatedNamespacePlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedNamespacePlacementList{} }, // Register versioned resource list
		&FederatedNamespacePlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedReplicaSetStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedReplicaSet,
		FederatedReplicaSetSchemeFns{},
		func() runtime.Object { return &FederatedReplicaSet{} },     // Register versioned resource
		func() runtime.Object { return &FederatedReplicaSetList{} }, // Register versioned resource list
		&FederatedReplicaSetStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedReplicaSetOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedReplicaSetOverride,
		FederatedReplicaSetOverrideSchemeFns{},
		func() runtime.Object { return &FederatedReplicaSetOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedReplicaSetOverrideList{} }, // Register versioned resource list
		&FederatedReplicaSetOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedReplicaSetPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedReplicaSetPlacement,
		FederatedReplicaSetPlacementSchemeFns{},
		func() runtime.Object { return &FederatedReplicaSetPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedReplicaSetPlacementList{} }, // Register versioned resource list
		&FederatedReplicaSetPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecret,
		FederatedSecretSchemeFns{},
		func() runtime.Object { return &FederatedSecret{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretList{} }, // Register versioned resource list
		&FederatedSecretStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretOverrideStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecretOverride,
		FederatedSecretOverrideSchemeFns{},
		func() runtime.Object { return &FederatedSecretOverride{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretOverrideList{} }, // Register versioned resource list
		&FederatedSecretOverrideStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedSecretPlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedSecretPlacement,
		FederatedSecretPlacementSchemeFns{},
		func() runtime.Object { return &FederatedSecretPlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedSecretPlacementList{} }, // Register versioned resource list
		&FederatedSecretPlacementStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedServiceStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedService,
		FederatedServiceSchemeFns{},
		func() runtime.Object { return &FederatedService{} },     // Register versioned resource
		func() runtime.Object { return &FederatedServiceList{} }, // Register versioned resource list
		&FederatedServiceStrategy{builders.StorageStrategySingleton},
	)
	federationFederatedServicePlacementStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalFederatedServicePlacement,
		FederatedServicePlacementSchemeFns{},
		func() runtime.Object { return &FederatedServicePlacement{} },     // Register versioned resource
		func() runtime.Object { return &FederatedServicePlacementList{} }, // Register versioned resource list
		&FederatedServicePlacementStrategy{builders.StorageStrategySingleton},
	)
	federationPropagatedVersionStorage = builders.NewApiResource( // Resource status endpoint
		federation.InternalPropagatedVersion,
		PropagatedVersionSchemeFns{},
		func() runtime.Object { return &PropagatedVersion{} },     // Register versioned resource
		func() runtime.Object { return &PropagatedVersionList{} }, // Register versioned resource list
		&PropagatedVersionStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("federation.k8s.io", "v1alpha1").WithResources(
		federationFederatedClusterStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedClusterStatus,
			FederatedClusterSchemeFns{},
			func() runtime.Object { return &FederatedCluster{} },     // Register versioned resource
			func() runtime.Object { return &FederatedClusterList{} }, // Register versioned resource list
			&FederatedClusterStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedConfigMapStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedConfigMapStatus,
			FederatedConfigMapSchemeFns{},
			func() runtime.Object { return &FederatedConfigMap{} },     // Register versioned resource
			func() runtime.Object { return &FederatedConfigMapList{} }, // Register versioned resource list
			&FederatedConfigMapStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedConfigMapOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedConfigMapOverrideStatus,
			FederatedConfigMapOverrideSchemeFns{},
			func() runtime.Object { return &FederatedConfigMapOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedConfigMapOverrideList{} }, // Register versioned resource list
			&FederatedConfigMapOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedConfigMapPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedConfigMapPlacementStatus,
			FederatedConfigMapPlacementSchemeFns{},
			func() runtime.Object { return &FederatedConfigMapPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedConfigMapPlacementList{} }, // Register versioned resource list
			&FederatedConfigMapPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedDeploymentStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedDeploymentStatus,
			FederatedDeploymentSchemeFns{},
			func() runtime.Object { return &FederatedDeployment{} },     // Register versioned resource
			func() runtime.Object { return &FederatedDeploymentList{} }, // Register versioned resource list
			&FederatedDeploymentStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedDeploymentOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedDeploymentOverrideStatus,
			FederatedDeploymentOverrideSchemeFns{},
			func() runtime.Object { return &FederatedDeploymentOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedDeploymentOverrideList{} }, // Register versioned resource list
			&FederatedDeploymentOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedDeploymentPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedDeploymentPlacementStatus,
			FederatedDeploymentPlacementSchemeFns{},
			func() runtime.Object { return &FederatedDeploymentPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedDeploymentPlacementList{} }, // Register versioned resource list
			&FederatedDeploymentPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedJobStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedJobStatus,
			FederatedJobSchemeFns{},
			func() runtime.Object { return &FederatedJob{} },     // Register versioned resource
			func() runtime.Object { return &FederatedJobList{} }, // Register versioned resource list
			&FederatedJobStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedJobOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedJobOverrideStatus,
			FederatedJobOverrideSchemeFns{},
			func() runtime.Object { return &FederatedJobOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedJobOverrideList{} }, // Register versioned resource list
			&FederatedJobOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedJobPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedJobPlacementStatus,
			FederatedJobPlacementSchemeFns{},
			func() runtime.Object { return &FederatedJobPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedJobPlacementList{} }, // Register versioned resource list
			&FederatedJobPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedNamespacePlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedNamespacePlacementStatus,
			FederatedNamespacePlacementSchemeFns{},
			func() runtime.Object { return &FederatedNamespacePlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedNamespacePlacementList{} }, // Register versioned resource list
			&FederatedNamespacePlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedReplicaSetStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedReplicaSetStatus,
			FederatedReplicaSetSchemeFns{},
			func() runtime.Object { return &FederatedReplicaSet{} },     // Register versioned resource
			func() runtime.Object { return &FederatedReplicaSetList{} }, // Register versioned resource list
			&FederatedReplicaSetStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedReplicaSetOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedReplicaSetOverrideStatus,
			FederatedReplicaSetOverrideSchemeFns{},
			func() runtime.Object { return &FederatedReplicaSetOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedReplicaSetOverrideList{} }, // Register versioned resource list
			&FederatedReplicaSetOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedReplicaSetPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedReplicaSetPlacementStatus,
			FederatedReplicaSetPlacementSchemeFns{},
			func() runtime.Object { return &FederatedReplicaSetPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedReplicaSetPlacementList{} }, // Register versioned resource list
			&FederatedReplicaSetPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretStatus,
			FederatedSecretSchemeFns{},
			func() runtime.Object { return &FederatedSecret{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretList{} }, // Register versioned resource list
			&FederatedSecretStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretOverrideStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretOverrideStatus,
			FederatedSecretOverrideSchemeFns{},
			func() runtime.Object { return &FederatedSecretOverride{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretOverrideList{} }, // Register versioned resource list
			&FederatedSecretOverrideStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedSecretPlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedSecretPlacementStatus,
			FederatedSecretPlacementSchemeFns{},
			func() runtime.Object { return &FederatedSecretPlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedSecretPlacementList{} }, // Register versioned resource list
			&FederatedSecretPlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedServiceStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedServiceStatus,
			FederatedServiceSchemeFns{},
			func() runtime.Object { return &FederatedService{} },     // Register versioned resource
			func() runtime.Object { return &FederatedServiceList{} }, // Register versioned resource list
			&FederatedServiceStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationFederatedServicePlacementStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalFederatedServicePlacementStatus,
			FederatedServicePlacementSchemeFns{},
			func() runtime.Object { return &FederatedServicePlacement{} },     // Register versioned resource
			func() runtime.Object { return &FederatedServicePlacementList{} }, // Register versioned resource list
			&FederatedServicePlacementStatusStrategy{builders.StatusStorageStrategySingleton},
		), federationPropagatedVersionStorage,
		builders.NewApiResource( // Resource status endpoint
			federation.InternalPropagatedVersionStatus,
			PropagatedVersionSchemeFns{},
			func() runtime.Object { return &PropagatedVersion{} },     // Register versioned resource
			func() runtime.Object { return &PropagatedVersionList{} }, // Register versioned resource list
			&PropagatedVersionStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// FederatedCluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedClusterSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedCluster `json:"items"`
}

//
// FederatedConfigMap Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedConfigMapSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedConfigMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedConfigMap `json:"items"`
}

//
// FederatedConfigMapOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedConfigMapOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedConfigMapOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedConfigMapOverride `json:"items"`
}

//
// FederatedConfigMapPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedConfigMapPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedConfigMapPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedConfigMapPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedConfigMapPlacement `json:"items"`
}

//
// FederatedDeployment Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedDeploymentSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedDeployment `json:"items"`
}

//
// FederatedDeploymentOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedDeploymentOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedDeploymentOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedDeploymentOverride `json:"items"`
}

//
// FederatedDeploymentPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedDeploymentPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedDeploymentPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedDeploymentPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedDeploymentPlacement `json:"items"`
}

//
// FederatedJob Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedJobSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedJobStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedJobStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedJob `json:"items"`
}

//
// FederatedJobOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedJobOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedJobOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedJobOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedJobOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedJobOverride `json:"items"`
}

//
// FederatedJobPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedJobPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedJobPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedJobPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedJobPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedJobPlacement `json:"items"`
}

//
// FederatedNamespacePlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedNamespacePlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedNamespacePlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedNamespacePlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedNamespacePlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedNamespacePlacement `json:"items"`
}

//
// FederatedReplicaSet Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSet `json:"items"`
}

//
// FederatedReplicaSetOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSetOverride `json:"items"`
}

//
// FederatedReplicaSetPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSetPlacement `json:"items"`
}

//
// FederatedSecret Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecret `json:"items"`
}

//
// FederatedSecretOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretOverrideSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecretOverride `json:"items"`
}

//
// FederatedSecretPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretPlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedSecretPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedSecretPlacement `json:"items"`
}

//
// FederatedService Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedServiceSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedServiceStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedServiceStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedService `json:"items"`
}

//
// FederatedServicePlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedServicePlacementSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type FederatedServicePlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedServicePlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedServicePlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedServicePlacement `json:"items"`
}

//
// PropagatedVersion Functions and Structs
//
// +k8s:deepcopy-gen=false
type PropagatedVersionSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type PropagatedVersionStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type PropagatedVersionStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PropagatedVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PropagatedVersion `json:"items"`
}
