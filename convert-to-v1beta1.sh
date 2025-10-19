#!/bin/bash
set -e

# API groups that need conversion from v1alpha1 to v1beta1
API_GROUPS="firewall logpush originssl r2 spectrum ssl sslsaas transform"

BASE_DIR="/home/rossg/src/crossplane-providers/provider-cloudflare/apis"

for GROUP in $API_GROUPS; do
    echo "Converting $GROUP to v1beta1..."

    # Create v1beta1 directory
    mkdir -p "$BASE_DIR/$GROUP/v1beta1"

    # Get the API group name (handle special cases)
    case "$GROUP" in
        sslsaas)
            GROUP_NAME="sslsaas"
            ;;
        originssl)
            GROUP_NAME="originssl"
            ;;
        *)
            GROUP_NAME="$GROUP"
            ;;
    esac

    # Copy all non-generated Go files from v1alpha1 to v1beta1
    for file in "$BASE_DIR/$GROUP/v1alpha1"/*.go; do
        filename=$(basename "$file")

        # Skip generated files
        if [[ "$filename" == zz_generated* ]]; then
            continue
        fi

        # Read the file and convert it
        sed -e "s/package v1alpha1/package v1beta1/g" \
            -e "s/v1alpha1/${GROUP}v1alpha1/g" \
            -e "s/+groupName=${GROUP_NAME}.cloudflare.crossplane.io/+groupName=${GROUP_NAME}.cloudflare.m.crossplane.io/g" \
            -e "s/Group   = \"${GROUP_NAME}.cloudflare.crossplane.io\"/Group   = \"${GROUP_NAME}.cloudflare.m.crossplane.io\"/g" \
            -e "s/CRDGroup   = \"${GROUP_NAME}.cloudflare.crossplane.io\"/CRDGroup   = \"${GROUP_NAME}.cloudflare.m.crossplane.io\"/g" \
            -e "s/+versionName=v1alpha1/+versionName=v1beta1/g" \
            -e "s/Version = \"v1alpha1\"/Version = \"v1beta1\"/g" \
            -e "s/CRDVersion = \"v1alpha1\"/CRDVersion = \"v1beta1\"/g" \
            -e "s/scope=Cluster/scope=Namespaced/g" \
            "$file" > "$BASE_DIR/$GROUP/v1beta1/$filename"
    done

    echo "  ✓ Created v1beta1 for $GROUP"
done

echo ""
echo "All API groups converted to v1beta1!"
echo "Next steps:"
echo "1. Review the generated files"
echo "2. Run 'make generate' to create CRDs"
echo "3. Update controllers to use v1beta1"
