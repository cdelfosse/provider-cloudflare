/*
Copyright 2025 The Crossplane Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// ZoneParameters are the configurable fields of a Zone.
type ZoneParameters struct {
	// Name is the name of the Zone, which should be a valid
	// domain.
	// +kubebuilder:validation:Format=hostname
	// +kubebuilder:validation:MaxLength=253
	// +immutable
	Name string `json:"name"`

	// AccountID is the account ID under which this Zone will be
	// created.
	// +immutable
	// +optional
	AccountID *string `json:"accountId,omitempty"`

	// JumpStart enables attempting to import existing DNS records
	// when a new Zone is created.
	//
	// WARNING: When enabled, Cloudflare automatically creates DNS records
	// by scanning your domain's existing nameservers. These auto-created
	// records will NOT be managed by Crossplane and will exist only in
	// Cloudflare. To manage them with Crossplane, you must:
	// 1. Create corresponding Record resources with matching settings
	// 2. Import the external records using crossplane.io/external-name annotation
	//
	// Recommendation: Leave disabled (false) for new zones to maintain
	// full Crossplane control over DNS records.
	// +kubebuilder:default=false
	// +immutable
	// +optional
	JumpStart bool `json:"jumpStart"`

	// Paused indicates if the zone is only using Cloudflare DNS services.
	// +optional
	Paused *bool `json:"paused,omitempty"`

	// PlanID indicates the plan that this Zone will be subscribed
	// to.
	// +optional
	PlanID *string `json:"planId,omitempty"`

	// Type indicates the type of this zone - partial (partner-hosted
	// or CNAME only) or full.
	// +kubebuilder:validation:Enum=full;partial
	// +kubebuilder:default=full
	// +immutable
	// +optional
	Type *string `json:"type,omitempty"`
}

// ZoneObservation are the observable fields of a Zone.
type ZoneObservation struct {
	// AccountID is the account ID that this zone exists under
	AccountID string `json:"accountId,omitempty"`

	// AccountName is the account name that this zone exists under
	Account string `json:"accountName,omitempty"`

	// Status indicates if this zone is active or pending
	Status string `json:"status,omitempty"`

	// Plan indicates the Cloudflare plan type for this zone
	Plan string `json:"plan,omitempty"`

	// NameServers lists the nameservers for this Zone
	NameServers []string `json:"nameServers,omitempty"`

	// CreatedOn indicates when this zone was created
	// on Cloudflare.
	CreatedOn *metav1.Time `json:"createdOn,omitempty"`

	// ModifiedOn indicates when this zone was modified
	// on Cloudflare.
	ModifiedOn *metav1.Time `json:"modifiedOn,omitempty"`
}

// A ZoneSpec defines the desired state of a Zone.
type ZoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ZoneParameters `json:"forProvider"`
}

// A ZoneStatus represents the observed state of a Zone.
type ZoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Zone is a set of common settings applied to one or more domains.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="ACCOUNT",type="string",JSONPath=".status.atProvider.accountId"
// +kubebuilder:printcolumn:name="PLAN",type="string",JSONPath=".status.atProvider.plan"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,cloudflare}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zone objects.
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}