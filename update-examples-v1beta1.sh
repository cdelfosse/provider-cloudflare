#!/bin/bash
set -e

echo "Updating all examples to use v1beta1 namespaced APIs..."

# Find all example YAML files and update them
find /home/rossg/src/crossplane-providers/provider-cloudflare/examples -name "*.yaml" -type f | while read -r file; do
    # Skip if file doesn't contain API version references
    if ! grep -q "\.crossplane\.io/" "$file"; then
        continue
    fi

    echo "  Updating $(basename $file)..."

    # Update API groups to use .m. (namespaced) and v1beta1
    sed -i 's|cache\.cloudflare\.crossplane\.io/v1alpha1|cache.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|dns\.cloudflare\.crossplane\.io/v1alpha1|dns.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|emailrouting\.cloudflare\.crossplane\.io/v1alpha1|emailrouting.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|firewall\.cloudflare\.crossplane\.io/v1alpha1|firewall.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|loadbalancing\.cloudflare\.crossplane\.io/v1alpha1|loadbalancing.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|logpush\.cloudflare\.crossplane\.io/v1alpha1|logpush.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|originssl\.cloudflare\.crossplane\.io/v1alpha1|originssl.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|r2\.cloudflare\.crossplane\.io/v1alpha1|r2.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|rulesets\.cloudflare\.crossplane\.io/v1alpha1|rulesets.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|security\.cloudflare\.crossplane\.io/v1alpha1|security.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|spectrum\.cloudflare\.crossplane\.io/v1alpha1|spectrum.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|ssl\.cloudflare\.crossplane\.io/v1alpha1|ssl.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|sslsaas\.cloudflare\.crossplane\.io/v1alpha1|sslsaas.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|transform\.cloudflare\.crossplane\.io/v1alpha1|transform.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|cloudflare\.crossplane\.io/v1alpha1|cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|workers\.cloudflare\.crossplane\.io/v1alpha1|workers.cloudflare.m.crossplane.io/v1beta1|g' "$file"
    sed -i 's|zone\.cloudflare\.crossplane\.io/v1alpha1|zone.cloudflare.m.crossplane.io/v1beta1|g' "$file"

    # Add namespace field if it doesn't exist (after metadata: line, before name:)
    # Only for resources that are not ProviderConfig
    if ! grep -q "kind: ProviderConfig" "$file" && ! grep -q "namespace:" "$file"; then
        sed -i '/^metadata:/a\  namespace: default' "$file"
    fi
done

echo "✓ All examples updated to use v1beta1 namespaced APIs"
