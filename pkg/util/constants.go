/*
Copyright 2020 The Karmada Authors.

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

package util

import "time"

// Define labels used by karmada system.
const (
	// ServiceNamespaceLabel is added to work object, which is report by member cluster, to specify service namespace associated with EndpointSlice.
	ServiceNamespaceLabel = "endpointslice.karmada.io/namespace"

	// ServiceNameLabel is added to work object, which is report by member cluster, to specify service name associated with EndpointSlice.
	ServiceNameLabel = "endpointslice.karmada.io/name"

	// MultiClusterServiceNamespaceLabel is added to work object, represents the work is managed by the corresponding MultiClusterService
	// This label indicates the namespace.
	MultiClusterServiceNamespaceLabel = "multiclusterservice.karmada.io/namespace"

	// MultiClusterServiceNameLabel is added to work object, represents the work is managed by the corresponding MultiClusterService
	// This label indicates the name.
	MultiClusterServiceNameLabel = "multiclusterservice.karmada.io/name"

	// PropagationInstruction is used to mark a resource(like Work) propagation instruction.
	// Valid values includes:
	// - suppressed: indicates that the resource should not be propagated.
	//
	// Note: This instruction is intended to set on Work objects to indicate the Work should be ignored by
	// execution controller. The instruction maybe deprecated once we extend the Work API and no other scenario want this.
	PropagationInstruction = "propagation.karmada.io/instruction"

	// FederatedResourceQuotaNamespaceLabel is added to Work to specify associated FederatedResourceQuota's namespace.
	FederatedResourceQuotaNamespaceLabel = "federatedresourcequota.karmada.io/namespace"

	// FederatedResourceQuotaNameLabel is added to Work to specify associated FederatedResourceQuota's name.
	FederatedResourceQuotaNameLabel = "federatedresourcequota.karmada.io/name"

	// ManagedByKarmadaLabel is a reserved karmada label to indicate whether resources are managed by karmada controllers.
	ManagedByKarmadaLabel = "karmada.io/managed"

	// EndpointSliceDispatchControllerLabelValue indicates the endpointSlice are controlled by Karmada
	EndpointSliceDispatchControllerLabelValue = "endpointslice-dispatch-controller.karmada.io"

	// RetainReplicasLabel is a reserved label to indicate whether the replicas should be retained. e.g:
	// resourcetemplate.karmada.io/retain-replicas: true   // with value `true` indicates retain
	// resourcetemplate.karmada.io/retain-replicas: false  // with value `false` and others, indicates not retain
	RetainReplicasLabel = "resourcetemplate.karmada.io/retain-replicas"

	// ResourceTemplateClaimedByLabel is added to the ResourceTemplate, indicating which resource is in charge of propagating the ResourceTemplate.
	ResourceTemplateClaimedByLabel = "resourcetemplate.karmada.io/claimed-by"
)

const (
	// ManagedByKarmadaLabelValue indicates that resources are managed by karmada controllers.
	ManagedByKarmadaLabelValue = "true"

	// RetainReplicasValue is an optional value of RetainReplicasLabel, indicating retain
	RetainReplicasValue = "true"

	// PropagationInstructionSuppressed indicates that the resource should not be propagated.
	PropagationInstructionSuppressed = "suppressed"
)

// Define annotations used by karmada system.
const (
	// PolicyPlacementAnnotation is the annotation of a policy's placement.
	// It is intended to set on ResourceBinding or ClusterResourceBinding objects to record applied placement declaration.
	// The placement could be either PropagationPolicy's or ClusterPropagationPolicy's.
	PolicyPlacementAnnotation = "policy.karmada.io/applied-placement"

	// AppliedOverrides is the annotation which used to record override items an object applied.
	// It is intended to set on Work objects to record applied overrides.
	// The overrides items should be sorted alphabetically in ascending order by OverridePolicy's name.
	AppliedOverrides = "policy.karmada.io/applied-overrides"

	// AppliedClusterOverrides is the annotation which used to record override items an object applied.
	// It is intended to set on Work objects to record applied overrides.
	// The overrides items should be sorted alphabetically in ascending order by ClusterOverridePolicy's name.
	AppliedClusterOverrides = "policy.karmada.io/applied-cluster-overrides"

	// EndPointSliceProvisionClusterAnnotation is added to work of the dispatch EndpointSlice in consumption clusters's namespace.
	EndpointSliceProvisionClusterAnnotation = "endpointslice.karmada.io/provision-cluster"
)

// Define finalizers used by karmada system.
const (
	// ClusterControllerFinalizer is added to Cluster to ensure Work as well as the
	// execution space (namespace) is deleted before itself is deleted.
	ClusterControllerFinalizer = "karmada.io/cluster-controller"

	// ExecutionControllerFinalizer is added to Work to ensure manifests propagated to member cluster
	// is deleted before Work itself is deleted.
	ExecutionControllerFinalizer = "karmada.io/execution-controller"

	// BindingControllerFinalizer is added to ResourceBinding to ensure related Works are deleted
	// before ResourceBinding itself is deleted.
	BindingControllerFinalizer = "karmada.io/binding-controller"

	// MCSEndpointSliceCollectControllerFinalizer is added to mcs to ensure related Works in provider clusters are deleted
	MCSEndpointSliceCollectControllerFinalizer = "karmada.io/mcs-endpointslice-collect-controller"

	// MCSEndpointSliceDispatchControllerFinalizer is added to mcs to ensure related Works in consumption clusters are deleted
	MCSEndpointSliceDispatchControllerFinalizer = "karmada.io/mcs-endpointslice-dispatch-controller"

	// ClusterResourceBindingControllerFinalizer is added to ClusterResourceBinding to ensure related Works are deleted
	// before ClusterResourceBinding itself is deleted.
	ClusterResourceBindingControllerFinalizer = "karmada.io/cluster-resource-binding-controller"

	// MCSControllerFinalizer is added to Cluster to ensure service work is deleted before itself is deleted.
	MCSControllerFinalizer = "karmada.io/multiclusterservice-controller"
)

const (
	// ProviderField indicates the 'provider' field of a cluster
	ProviderField = "provider"
	// RegionField indicates the 'region' field of a cluster
	RegionField = "region"
	// ZoneField indicates the 'zone' field of a cluster
	ZoneField = "zone"
)

// Define resource kind.
const (
	// DeploymentKind indicates the target resource is a deployment
	DeploymentKind = "Deployment"
	// ServiceKind indicates the target resource is a service
	ServiceKind = "Service"
	// IngressKind indicates the target resource is a ingress
	IngressKind = "Ingress"
	// CronJobKind indicates the target resource is a cronjob
	CronJobKind = "CronJob"
	// JobKind indicates the target resource is a job
	JobKind = "Job"
	// PodKind indicates the target resource is a pod
	PodKind = "Pod"
	// ServiceAccountKind indicates the target resource is a serviceaccount
	ServiceAccountKind = "ServiceAccount"
	// ReplicaSetKind indicates the target resource is a replicaset
	ReplicaSetKind = "ReplicaSet"
	// StatefulSetKind indicates the target resource is a statefulset
	StatefulSetKind = "StatefulSet"
	// DaemonSetKind indicates the target resource is a daemonset
	DaemonSetKind = "DaemonSet"
	// EndpointSliceKind indicates the target resource is a endpointslice
	EndpointSliceKind = "EndpointSlice"
	// PersistentVolumeClaimKind indicates the target resource is a persistentvolumeclaim
	PersistentVolumeClaimKind = "PersistentVolumeClaim"
	// PersistentVolumeKind indicates the target resource is a persistentvolume
	PersistentVolumeKind = "PersistentVolume"
	// HorizontalPodAutoscalerKind indicates the target resource is a horizontalpodautoscaler
	HorizontalPodAutoscalerKind = "HorizontalPodAutoscaler"
	// PodDisruptionBudgetKind indicates the target resource is a poddisruptionbudget
	PodDisruptionBudgetKind = "PodDisruptionBudget"
	// ClusterRoleKind indicates the target resource is a clusterrole
	ClusterRoleKind = "ClusterRole"
	// ClusterRoleBindingKind indicates the target resource is a clusterrolebinding
	ClusterRoleBindingKind = "ClusterRoleBinding"
	// CRDKind indicates the target resource is a CustomResourceDefinition
	CRDKind = "CustomResourceDefinition"

	// ServiceExportKind indicates the target resource is a serviceexport crd
	ServiceExportKind = "ServiceExport"
	// ServiceImportKind indicates the target resource is a serviceimport crd
	ServiceImportKind = "ServiceImport"

	// MultiClusterServiceKind indicates the target resource is a MultiClusterService
	MultiClusterServiceKind = "MultiClusterService"
)

// Define resource filed
const (
	// SpecField indicates the 'spec' field of a resource
	SpecField = "spec"
	// ReplicasField indicates the 'replicas' field of a resource
	ReplicasField = "replicas"
	// ReadyReplicasField indicates the 'readyReplicas' field of a resource status
	ReadyReplicasField = "readyReplicas"
	// ParallelismField indicates the 'parallelism' field of a job
	ParallelismField = "parallelism"
	// CompletionsField indicates the 'completions' field of a job
	CompletionsField = "completions"
)

const (
	// NamespaceKarmadaSystem is the karmada system namespace.
	NamespaceKarmadaSystem = "karmada-system"
)

// ContextKey is the key of context.
type ContextKey string

const (
	// ContextKeyObject is the context value key of a resource.
	ContextKeyObject ContextKey = "object"
)

const (
	// CacheSyncTimeout refers to the time limit set on waiting for cache to sync
	CacheSyncTimeout = 2 * time.Minute
)
