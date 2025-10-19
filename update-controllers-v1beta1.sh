#!/bin/bash
set -e

echo "Updating all controllers to use v1beta1 APIs..."

# Find all controller files and update imports
find /home/rossg/src/crossplane-providers/provider-cloudflare/internal/controller -name "*.go" -type f | while read -r file; do
    # Skip if file doesn't contain v1alpha1 imports
    if ! grep -q "v1alpha1" "$file"; then
        continue
    fi

    echo "  Updating $(basename $(dirname $file))/$(basename $file)..."

    # Update imports - replace /v1alpha1 with /v1beta1 in import paths
    sed -i 's|/apis/cache/v1alpha1|/apis/cache/v1beta1|g' "$file"
    sed -i 's|/apis/dns/v1alpha1|/apis/dns/v1beta1|g' "$file"
    sed -i 's|/apis/emailrouting/v1alpha1|/apis/emailrouting/v1beta1|g' "$file"
    sed -i 's|/apis/firewall/v1alpha1|/apis/firewall/v1beta1|g' "$file"
    sed -i 's|/apis/loadbalancing/v1alpha1|/apis/loadbalancing/v1beta1|g' "$file"
    sed -i 's|/apis/logpush/v1alpha1|/apis/logpush/v1beta1|g' "$file"
    sed -i 's|/apis/originssl/v1alpha1|/apis/originssl/v1beta1|g' "$file"
    sed -i 's|/apis/r2/v1alpha1|/apis/r2/v1beta1|g' "$file"
    sed -i 's|/apis/rulesets/v1alpha1|/apis/rulesets/v1beta1|g' "$file"
    sed -i 's|/apis/security/v1alpha1|/apis/security/v1beta1|g' "$file"
    sed -i 's|/apis/spectrum/v1alpha1|/apis/spectrum/v1beta1|g' "$file"
    sed -i 's|/apis/ssl/v1alpha1|/apis/ssl/v1beta1|g' "$file"
    sed -i 's|/apis/sslsaas/v1alpha1|/apis/sslsaas/v1beta1|g' "$file"
    sed -i 's|/apis/transform/v1alpha1|/apis/transform/v1beta1|g' "$file"
    sed -i 's|/apis/v1alpha1|/apis/v1beta1|g' "$file"
    sed -i 's|/apis/workers/v1alpha1|/apis/workers/v1beta1|g' "$file"
    sed -i 's|/apis/zone/v1alpha1|/apis/zone/v1beta1|g' "$file"
done

echo "✓ All controllers updated to use v1beta1 APIs"
