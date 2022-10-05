/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BootDiskInitializeParamsObservation struct {
}

type BootDiskInitializeParamsParameters struct {

	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceFromTemplateAdvancedMachineFeaturesObservation struct {
}

type InstanceFromTemplateAdvancedMachineFeaturesParameters struct {

	// +kubebuilder:validation:Optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty" tf:"enable_nested_virtualization,omitempty"`

	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type InstanceFromTemplateAttachedDiskObservation struct {
}

type InstanceFromTemplateAttachedDiskParameters struct {

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`

	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRaw *string `json:"diskEncryptionKeyRaw,omitempty" tf:"disk_encryption_key_raw"`

	// +kubebuilder:validation:Optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256"`

	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source"`
}

type InstanceFromTemplateBootDiskObservation struct {
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type InstanceFromTemplateBootDiskParameters struct {

	// Default is 6 minutes.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InitializeParams []BootDiskInitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type InstanceFromTemplateConfidentialInstanceConfigObservation struct {
}

type InstanceFromTemplateConfidentialInstanceConfigParameters struct {

	// +kubebuilder:validation:Required
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute,omitempty"`
}

type InstanceFromTemplateGuestAcceleratorObservation struct {
}

type InstanceFromTemplateGuestAcceleratorParameters struct {

	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type InstanceFromTemplateNetworkInterfaceObservation struct {

	// +kubebuilder:validation:Optional
	IPv6AccessConfig []NetworkInterfaceIPv6AccessConfigObservation `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InstanceFromTemplateNetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	AccessConfig []NetworkInterfaceAccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// +kubebuilder:validation:Optional
	AliasIPRange []NetworkInterfaceAliasIPRangeParameters `json:"aliasIpRange,omitempty" tf:"alias_ip_range,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AccessConfig []NetworkInterfaceIPv6AccessConfigParameters `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// +kubebuilder:validation:Optional
	QueueCount *float64 `json:"queueCount,omitempty" tf:"queue_count,omitempty"`

	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type InstanceFromTemplateObservation struct {

	// +kubebuilder:validation:Optional
	BootDisk []InstanceFromTemplateBootDiskObservation `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	CPUPlatform *string `json:"cpuPlatform,omitempty" tf:"cpu_platform,omitempty"`

	CurrentStatus *string `json:"currentStatus,omitempty" tf:"current_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []InstanceFromTemplateNetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type InstanceFromTemplateParameters struct {

	// +kubebuilder:validation:Optional
	AdvancedMachineFeatures []InstanceFromTemplateAdvancedMachineFeaturesParameters `json:"advancedMachineFeatures,omitempty" tf:"advanced_machine_features,omitempty"`

	// Default is 6 minutes.
	// +kubebuilder:validation:Optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// +kubebuilder:validation:Optional
	AttachedDisk []InstanceFromTemplateAttachedDiskParameters `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`

	// +kubebuilder:validation:Optional
	BootDisk []InstanceFromTemplateBootDiskParameters `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// +kubebuilder:validation:Optional
	CanIPForward *bool `json:"canIpForward,omitempty" tf:"can_ip_forward,omitempty"`

	// +kubebuilder:validation:Optional
	ConfidentialInstanceConfig []InstanceFromTemplateConfidentialInstanceConfigParameters `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredStatus *string `json:"desiredStatus,omitempty" tf:"desired_status,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDisplay *bool `json:"enableDisplay,omitempty" tf:"enable_display,omitempty"`

	// +kubebuilder:validation:Optional
	GuestAccelerator []InstanceFromTemplateGuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`

	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []InstanceFromTemplateNetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ReservationAffinity []InstanceFromTemplateReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// +kubebuilder:validation:Optional
	ResourcePolicies []*string `json:"resourcePolicies,omitempty" tf:"resource_policies,omitempty"`

	// +kubebuilder:validation:Optional
	Scheduling []InstanceFromTemplateSchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// +kubebuilder:validation:Optional
	ScratchDisk []InstanceFromTemplateScratchDiskParameters `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccount []InstanceFromTemplateServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []InstanceFromTemplateShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// Name or self link of an instance
	// template to create the instance based on.
	// +crossplane:generate:reference:type=InstanceTemplate
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceInstanceTemplate *string `json:"sourceInstanceTemplate,omitempty" tf:"source_instance_template,omitempty"`

	// Reference to a InstanceTemplate to populate sourceInstanceTemplate.
	// +kubebuilder:validation:Optional
	SourceInstanceTemplateRef *v1.Reference `json:"sourceInstanceTemplateRef,omitempty" tf:"-"`

	// Selector for a InstanceTemplate to populate sourceInstanceTemplate.
	// +kubebuilder:validation:Optional
	SourceInstanceTemplateSelector *v1.Selector `json:"sourceInstanceTemplateSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone that the machine should be created in. If not
	// set, the provider zone is used.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceFromTemplateReservationAffinityObservation struct {
}

type InstanceFromTemplateReservationAffinityParameters struct {

	// +kubebuilder:validation:Optional
	SpecificReservation []ReservationAffinitySpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceFromTemplateSchedulingObservation struct {
}

type InstanceFromTemplateSchedulingParameters struct {

	// +kubebuilder:validation:Optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`

	// +kubebuilder:validation:Optional
	MinNodeCpus *float64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus,omitempty"`

	// +kubebuilder:validation:Optional
	NodeAffinities []SchedulingNodeAffinitiesParameters `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`

	// +kubebuilder:validation:Optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`

	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningModel *string `json:"provisioningModel,omitempty" tf:"provisioning_model,omitempty"`
}

type InstanceFromTemplateScratchDiskObservation struct {
}

type InstanceFromTemplateScratchDiskParameters struct {

	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface"`
}

type InstanceFromTemplateServiceAccountObservation struct {
}

type InstanceFromTemplateServiceAccountParameters struct {

	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email"`

	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes"`
}

type InstanceFromTemplateShieldedInstanceConfigObservation struct {
}

type InstanceFromTemplateShieldedInstanceConfigParameters struct {

	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

type NetworkInterfaceAccessConfigObservation struct {
}

type NetworkInterfaceAccessConfigParameters struct {

	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip"`

	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name"`
}

type NetworkInterfaceAliasIPRangeObservation struct {
}

type NetworkInterfaceAliasIPRangeParameters struct {

	// +kubebuilder:validation:Optional
	IPCidrRange *string `json:"ipCidrRange,omitempty" tf:"ip_cidr_range"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name"`
}

type NetworkInterfaceIPv6AccessConfigObservation struct {
	ExternalIPv6 *string `json:"externalIpv6,omitempty" tf:"external_ipv6,omitempty"`

	ExternalIPv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty" tf:"external_ipv6_prefix_length,omitempty"`
}

type NetworkInterfaceIPv6AccessConfigParameters struct {

	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`

	// A unique name for the resource, required by GCE.
	// Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type ReservationAffinitySpecificReservationObservation struct {
}

type ReservationAffinitySpecificReservationParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SchedulingNodeAffinitiesObservation struct {
}

type SchedulingNodeAffinitiesParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// InstanceFromTemplateSpec defines the desired state of InstanceFromTemplate
type InstanceFromTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceFromTemplateParameters `json:"forProvider"`
}

// InstanceFromTemplateStatus defines the observed state of InstanceFromTemplate.
type InstanceFromTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceFromTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFromTemplate is the Schema for the InstanceFromTemplates API. Manages a VM instance resource within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceFromTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceFromTemplateSpec   `json:"spec"`
	Status            InstanceFromTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceFromTemplateList contains a list of InstanceFromTemplates
type InstanceFromTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceFromTemplate `json:"items"`
}

// Repository type metadata.
var (
	InstanceFromTemplate_Kind             = "InstanceFromTemplate"
	InstanceFromTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceFromTemplate_Kind}.String()
	InstanceFromTemplate_KindAPIVersion   = InstanceFromTemplate_Kind + "." + CRDGroupVersion.String()
	InstanceFromTemplate_GroupVersionKind = CRDGroupVersion.WithKind(InstanceFromTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceFromTemplate{}, &InstanceFromTemplateList{})
}