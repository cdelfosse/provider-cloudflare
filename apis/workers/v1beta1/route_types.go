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

// RouteParameters are the configurable fields of a DNS Route.
type RouteParameters struct {
	// Pattern is the URL pattern of the route.
	Pattern string `json:"pattern"`

	// Script is the name of the worker script.
	// +optional
	Script *string `json:"script,omitempty"`

	// ZoneID this Worker Route is managed on.
	// +immutable
	// +optional
	Zone *string `json:"zone,omitempty"`

	// ZoneRef references the Zone object this Worker Route is managed on.
	// +immutable
	// +optional
	ZoneRef *xpv1.Reference `json:"zoneRef,omitempty"`

	// ZoneSelector selects the Zone object this Worker Route is managed on.
	// +immutable
	// +optional
	ZoneSelector *xpv1.Selector `json:"zoneSelector,omitempty"`
}

// RouteObservation is the observable fields of a Worker Route.
type RouteObservation struct{}

// A RouteSpec defines the desired state of a Worker Route.
type RouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteParameters `json:"forProvider"`
}

// A RouteStatus represents the observed state of a Worker Route.
type RouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Route represents a single Worker Route managed on a Zone.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="PATTERN",type="string",JSONPath=".spec.forProvider.pattern"
// +kubebuilder:printcolumn:name="SCRIPT",type="string",JSONPath=".spec.forProvider.script"
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,cloudflare}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Worker Route objects
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}