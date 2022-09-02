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

type ProjectObservation struct {

	// an identifier for the resource with format projects/{{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The numeric identifier of the project.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`
}

type ProjectParameters struct {

	// Create the 'default' network automatically.  Default true.
	// If set to false, the default network will be deleted.  Note that, for quota purposes, you
	// will still need to have 1 network slot available to create the project successfully, even if
	// you set auto_create_network to false, since the network will exist momentarily.
	// +kubebuilder:validation:Optional
	AutoCreateNetwork *bool `json:"autoCreateNetwork,omitempty" tf:"auto_create_network,omitempty"`

	// The alphanumeric ID of the billing account this project
	// belongs to. The user or service account performing this operation with Terraform
	// must have at minimum Billing Account User privileges  on the billing account.
	// See Google Cloud Billing API Access Control
	// for more details.
	// +kubebuilder:validation:Optional
	BillingAccount *string `json:"billingAccount,omitempty" tf:"billing_account,omitempty"`

	// The numeric ID of the folder this project should be
	// created under. Only one of org_id or folder_id may be
	// specified. If the folder_id is specified, then the project is
	// created under the specified folder. Changing this forces the
	// project to be migrated to the newly specified folder.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the project.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The display name of the project.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The numeric ID of the organization this project belongs to.
	// Changing this forces a new project to be created.  Only one of
	// org_id or folder_id may be specified. If the org_id is
	// specified then the project is created at the top level. Changing
	// this forces the project to be migrated to the newly specified
	// organization.
	// The numeric ID of the organization this project belongs to.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// If true, the Terraform resource can be deleted
	// without deleting the Project via the Google API.
	// +kubebuilder:validation:Optional
	SkipDelete *bool `json:"skipDelete,omitempty" tf:"skip_delete,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API. Allows management of a Google Cloud Platform project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
