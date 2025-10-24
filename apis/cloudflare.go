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

// Package apis contains Kubernetes API for the Template provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	cachev1beta1 "github.com/rossigee/provider-cloudflare/apis/cache/v1beta1"
	dnsv1beta1 "github.com/rossigee/provider-cloudflare/apis/dns/v1beta1"
	emailroutingv1beta1 "github.com/rossigee/provider-cloudflare/apis/emailrouting/v1beta1"
	firewallv1beta1 "github.com/rossigee/provider-cloudflare/apis/firewall/v1beta1"
	loadbalancingv1beta1 "github.com/rossigee/provider-cloudflare/apis/loadbalancing/v1beta1"
	logpushv1beta1 "github.com/rossigee/provider-cloudflare/apis/logpush/v1beta1"
	originsslv1beta1 "github.com/rossigee/provider-cloudflare/apis/originssl/v1beta1"
	r2v1beta1 "github.com/rossigee/provider-cloudflare/apis/r2/v1beta1"
	rulesetsv1beta1 "github.com/rossigee/provider-cloudflare/apis/rulesets/v1beta1"
	securityv1beta1 "github.com/rossigee/provider-cloudflare/apis/security/v1beta1"
	spectrumv1beta1 "github.com/rossigee/provider-cloudflare/apis/spectrum/v1beta1"
	sslv1beta1 "github.com/rossigee/provider-cloudflare/apis/ssl/v1beta1"
	sslsaasv1beta1 "github.com/rossigee/provider-cloudflare/apis/sslsaas/v1beta1"
	transformv1beta1 "github.com/rossigee/provider-cloudflare/apis/transform/v1beta1"
	cloudflarev1beta1 "github.com/rossigee/provider-cloudflare/apis/v1beta1"
	workersv1beta1 "github.com/rossigee/provider-cloudflare/apis/workers/v1beta1"
	zonev1beta1 "github.com/rossigee/provider-cloudflare/apis/zone/v1beta1"

	"github.com/pkg/errors"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		cloudflarev1beta1.SchemeBuilder.AddToScheme,
		cachev1beta1.SchemeBuilder.AddToScheme,
		dnsv1beta1.SchemeBuilder.AddToScheme,
		emailroutingv1beta1.SchemeBuilder.AddToScheme,
		sslsaasv1beta1.SchemeBuilder.AddToScheme,
		originsslv1beta1.SchemeBuilder.AddToScheme,
		spectrumv1beta1.SchemeBuilder.AddToScheme,
		zonev1beta1.SchemeBuilder.AddToScheme,
		firewallv1beta1.SchemeBuilder.AddToScheme,
		workersv1beta1.SchemeBuilder.AddToScheme,
		transformv1beta1.SchemeBuilder.AddToScheme,
		rulesetsv1beta1.SchemeBuilder.AddToScheme,
		securityv1beta1.SchemeBuilder.AddToScheme,
		sslv1beta1.SchemeBuilder.AddToScheme,
		loadbalancingv1beta1.SchemeBuilder.AddToScheme,
		logpushv1beta1.SchemeBuilder.AddToScheme,
		r2v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

// VerifySchemeRegistration tests that all API types are properly registered
// This helps catch scheme registration issues that would cause panics during provider startup
func VerifySchemeRegistration() error {
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		return err
	}

	// Test that we can create instances of key types from all packages
	// This would fail if types weren't registered properly
	testTypes := []runtime.Object{
		// Core provider types
		&cloudflarev1beta1.ProviderConfig{},
		&cloudflarev1beta1.ProviderConfigList{},
		&cloudflarev1beta1.ProviderConfigUsage{},
		&cloudflarev1beta1.ProviderConfigUsageList{},

		// Zone and DNS
		&zonev1beta1.Zone{},
		&zonev1beta1.ZoneList{},
		&dnsv1beta1.Record{},
		&dnsv1beta1.RecordList{},

		// Load balancing
		&loadbalancingv1beta1.LoadBalancer{},
		&loadbalancingv1beta1.LoadBalancerList{},
		&loadbalancingv1beta1.LoadBalancerMonitor{},
		&loadbalancingv1beta1.LoadBalancerMonitorList{},
		&loadbalancingv1beta1.LoadBalancerPool{},
		&loadbalancingv1beta1.LoadBalancerPoolList{},

		// Security and firewall
		&securityv1beta1.BotManagement{},
		&securityv1beta1.BotManagementList{},
		&securityv1beta1.RateLimit{},
		&securityv1beta1.RateLimitList{},
		&securityv1beta1.Turnstile{},
		&securityv1beta1.TurnstileList{},

		// SSL and certificates
		&sslv1beta1.CertificatePack{},
		&sslv1beta1.CertificatePackList{},
		&sslv1beta1.TotalTLS{},
		&sslv1beta1.TotalTLSList{},
		&sslv1beta1.UniversalSSL{},
		&sslv1beta1.UniversalSSLList{},
		&originsslv1beta1.Certificate{},
		&originsslv1beta1.CertificateList{},

		// Applications and services
		&spectrumv1beta1.Application{},
		&spectrumv1beta1.ApplicationList{},
		&sslsaasv1beta1.CustomHostname{},
		&sslsaasv1beta1.CustomHostnameList{},
		&sslsaasv1beta1.FallbackOrigin{},
		&sslsaasv1beta1.FallbackOriginList{},

		// Rules and transforms
		&rulesetsv1beta1.Ruleset{},
		&rulesetsv1beta1.RulesetList{},
		&transformv1beta1.Rule{},
		&transformv1beta1.RuleList{},

		// Workers and edge computing
		&workersv1beta1.CronTrigger{},
		&workersv1beta1.CronTriggerList{},
		&workersv1beta1.Domain{},
		&workersv1beta1.DomainList{},
		&workersv1beta1.KVNamespace{},
		&workersv1beta1.KVNamespaceList{},
		&workersv1beta1.Route{},
		&workersv1beta1.RouteList{},
		&workersv1beta1.Script{},
		&workersv1beta1.ScriptList{},
		&workersv1beta1.Subdomain{},
		&workersv1beta1.SubdomainList{},

		// Cache and performance
		&cachev1beta1.CacheRule{},
		&cachev1beta1.CacheRuleList{},

		// Email and logging
		&emailroutingv1beta1.Rule{},
		&emailroutingv1beta1.RuleList{},
		&logpushv1beta1.Job{},
		&logpushv1beta1.JobList{},

		// Storage
		&r2v1beta1.Bucket{},
		&r2v1beta1.BucketList{},
	}

	for _, obj := range testTypes {
		gvk, err := SchemeGroupVersionKind(obj, scheme)
		if err != nil {
			return err
		}
		if gvk.Group == "" || gvk.Version == "" || gvk.Kind == "" {
			return errors.Errorf("type %T has incomplete GVK: %v", obj, gvk)
		}
	}

	return nil
}

// SchemeGroupVersionKind returns the GroupVersionKind for the given object using the scheme
func SchemeGroupVersionKind(obj runtime.Object, scheme *runtime.Scheme) (schema.GroupVersionKind, error) {
	gvks, _, err := scheme.ObjectKinds(obj)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	if len(gvks) == 0 {
		return schema.GroupVersionKind{}, errors.New("no GVKs found")
	}
	return gvks[0], nil
}
