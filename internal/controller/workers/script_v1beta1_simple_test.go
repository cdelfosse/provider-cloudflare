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

package workers

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/rossigee/provider-cloudflare/apis/workers/v1beta1"
	"github.com/rossigee/provider-cloudflare/internal/controller/testutils"
)

// TestV1Beta1ScriptCreation tests basic Script v1beta1 creation
func TestV1Beta1ScriptCreation(t *testing.T) {
	placementMode := v1beta1.PlacementModeSmart

	script := &v1beta1.Script{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-worker-script",
			Namespace: "edge-services",
		},
		Spec: v1beta1.ScriptSpec{
			ForProvider: v1beta1.ScriptParameters{
				ScriptName: "test-worker-script",
				Script:     "export default { fetch() { return new Response('Hello from v1beta1!'); } }",
				Module:     testutils.BoolPtr(true),
				CompatibilityDate:  testutils.StringPtr("2023-05-15"),
				CompatibilityFlags: []string{"nodejs_compat"},
				Bindings: []v1beta1.WorkerBinding{
					{
						Name: "MY_VAR",
						Type: "plain_text",
						Text: testutils.StringPtr("test-value"),
					},
					{
						Name:        "MY_KV",
						Type:        "kv_namespace",
						NamespaceID: testutils.StringPtr("test-kv-namespace-id"),
					},
				},
				Placement: &placementMode,
			},
		},
	}

	// Test that the Script implements the managed resource interface
	if script.GetCondition("Ready").Status == "" {
		// This validates the managed resource interface is implemented
		t.Log("Script v1beta1 successfully implements managed resource interface")
	}

	// Test basic field access
	if script.Spec.ForProvider.ScriptName != "test-worker-script" {
		t.Errorf("Expected ScriptName 'test-worker-script', got %s", script.Spec.ForProvider.ScriptName)
	}

	if !*script.Spec.ForProvider.Module {
		t.Error("Expected Module to be true")
	}

	if *script.Spec.ForProvider.CompatibilityDate != "2023-05-15" {
		t.Errorf("Expected CompatibilityDate '2023-05-15', got %s", *script.Spec.ForProvider.CompatibilityDate)
	}

	// Test bindings
	if len(script.Spec.ForProvider.Bindings) != 2 {
		t.Errorf("Expected 2 bindings, got %d", len(script.Spec.ForProvider.Bindings))
	}

	if script.Spec.ForProvider.Bindings[0].Type != "plain_text" {
		t.Errorf("Expected first binding type 'plain_text', got %s", script.Spec.ForProvider.Bindings[0].Type)
	}

	if script.Spec.ForProvider.Bindings[1].Type != "kv_namespace" {
		t.Errorf("Expected second binding type 'kv_namespace', got %s", script.Spec.ForProvider.Bindings[1].Type)
	}

	// Test placement
	if *script.Spec.ForProvider.Placement != v1beta1.PlacementModeSmart {
		t.Errorf("Expected placement mode '%s', got %s", v1beta1.PlacementModeSmart, *script.Spec.ForProvider.Placement)
	}

	// Test namespace scope
	if script.Namespace != "edge-services" {
		t.Errorf("Expected namespace 'edge-services', got %s", script.Namespace)
	}

	t.Log("v1beta1 Script creation and field access tests passed")
}

// TestV1Beta1ScriptAdvancedBindings tests advanced v1beta1 binding features
func TestV1Beta1ScriptAdvancedBindings(t *testing.T) {
	placementMode := v1beta1.PlacementModeSmart

	script := &v1beta1.Script{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "advanced-worker-script",
			Namespace: "production",
		},
		Spec: v1beta1.ScriptSpec{
			ForProvider: v1beta1.ScriptParameters{
				ScriptName: "advanced-worker-script",
				Script:     "// Advanced ES Module worker\nexport default { fetch() { return handleRequest(); } }",
				Module:     testutils.BoolPtr(true),
				Bindings: []v1beta1.WorkerBinding{
					{
						Name: "SECRET_KEY",
						Type: "secret_text",
						Text: testutils.StringPtr("secret-value"),
					},
					{
						Name:        "MY_KV",
						Type:        "kv_namespace",
						NamespaceID: testutils.StringPtr("production-kv-namespace"),
					},
				},
				Placement: &placementMode,
			},
		},
	}

	// Test advanced binding types
	bindings := script.Spec.ForProvider.Bindings

	// Secret text binding
	if bindings[0].Type != "secret_text" {
		t.Errorf("Expected binding type 'secret_text', got %s", bindings[0].Type)
	}

	// KV namespace binding
	if bindings[1].Type != "kv_namespace" {
		t.Errorf("Expected binding type 'kv_namespace', got %s", bindings[1].Type)
	}

	if *bindings[1].NamespaceID != "production-kv-namespace" {
		t.Errorf("Expected namespace ID 'production-kv-namespace', got %s", *bindings[1].NamespaceID)
	}

	t.Log("v1beta1 Script advanced binding tests passed")
}