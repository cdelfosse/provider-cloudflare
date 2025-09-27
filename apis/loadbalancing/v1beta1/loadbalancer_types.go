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

// LoadBalancerParameters define the desired state of a Cloudflare Load Balancer
type LoadBalancerParameters struct {
	// Zone is the zone ID where this load balancer will be created.
	// Load balancers are zone-scoped resources.
	// +required
	Zone string `json:"zone"`

	// Name is the DNS name for this load balancer.
	// +optional
	Name *string `json:"name,omitempty"`

	// Description is a human-readable description of the load balancer.
	// +optional
	Description *string `json:"description,omitempty"`

	// TTL is the DNS TTL for the load balancer.
	// +optional
	TTL *int `json:"ttl,omitempty"`

	// FallbackPool is the pool ID to use when all other pools are unhealthy.
	// +optional
	FallbackPool *string `json:"fallbackPool,omitempty"`

	// DefaultPools is the list of pool IDs ordered by their failover priority.
	// +optional
	DefaultPools []string `json:"defaultPools,omitempty"`

	// RegionPools maps regions to pool lists for geo-steering.
	// +optional
	RegionPools map[string][]string `json:"regionPools,omitempty"`

	// Proxied indicates whether traffic should be proxied through Cloudflare.
	// +optional
	Proxied *bool `json:"proxied,omitempty"`

	// Enabled indicates whether the load balancer is enabled.
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// SessionAffinity controls session stickiness.
	// Valid values: "none", "cookie", "ip_cookie"
	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty"`

	// SessionAffinityTTL is the TTL for session affinity in seconds.
	// +optional
	SessionAffinityTTL *int `json:"sessionAffinityTtl,omitempty"`

	// SteeringPolicy determines the load balancing steering policy.
	// Valid values: "off", "geo", "dynamic_latency", "random", "proximity", "least_outstanding_requests"
	// +optional
	SteeringPolicy *string `json:"steeringPolicy,omitempty"`
}

// LoadBalancerObservation contains the observable fields of a LoadBalancer
type LoadBalancerObservation struct {
	// ID is the unique identifier for the load balancer
	ID string `json:"id,omitempty"`

	// Name is the DNS name for this load balancer
	Name string `json:"name,omitempty"`

	// CreatedOn indicates when this load balancer was created
	CreatedOn *metav1.Time `json:"createdOn,omitempty"`

	// ModifiedOn indicates when this load balancer was last modified
	ModifiedOn *metav1.Time `json:"modifiedOn,omitempty"`

	// Status indicates the status of the load balancer
	Status string `json:"status,omitempty"`
}

// A LoadBalancerSpec defines the desired state of a LoadBalancer.
type LoadBalancerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LoadBalancerParameters `json:"forProvider"`
}

// A LoadBalancerStatus represents the observed state of a LoadBalancer.
type LoadBalancerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A LoadBalancer provides load balancing and traffic steering for web applications
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="NAME",type="string",JSONPath=".status.atProvider.name"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={crossplane,managed,cloudflare}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}