#!/bin/bash
# Script to fix crossplane-runtime MockClient missing Apply method

set -e

echo "Fixing crossplane-runtime MockClient Apply method..."

# Add MockApplyFn type after MockPatchFn
sed -i '/\/\/ A MockPatchFn is used to mock client\.Client'\''s Patch implementation\./a \
\/\/ A MockApplyFn is used to mock client.Client'\''s Apply implementation.\
type MockApplyFn func(ctx context.Context, obj runtime.ApplyConfiguration, opts ...client.ApplyOption) error' vendor/github.com/crossplane/crossplane-runtime/v2/pkg/test/fake.go

# Add MockApply field to MockClient struct
sed -i '/	MockPatch       MockPatchFn/a \
	MockApply       MockApplyFn' vendor/github.com/crossplane/crossplane-runtime/v2/pkg/test/fake.go

# Add NewMockApplyFn function
sed -i '/\/\/ NewMockPatchFn returns a MockPatchFn that returns the supplied error\./a \
\/\/ NewMockApplyFn returns a MockApplyFn that returns the supplied error.\
func NewMockApplyFn(err error) MockApplyFn {\
	return func(_ context.Context, _ runtime.ApplyConfiguration, _ ...client.ApplyOption) error {\
		return err\
	}\
}' vendor/github.com/crossplane/crossplane-runtime/v2/pkg/test/fake.go

# Add MockApply initialization in NewMockClient
sed -i '/		MockPatch:       NewMockPatchFn(nil),/a \
		MockApply:       NewMockApplyFn(nil),' vendor/github.com/crossplane/crossplane-runtime/v2/pkg/test/fake.go

# Add Apply method implementation
sed -i '/\/\/ Patch calls MockClient'\''s MockPatch function\./a \
\/\/ Apply calls MockClient'\''s MockApply function.\
func (c *MockClient) Apply(ctx context.Context, obj runtime.ApplyConfiguration, opts ...client.ApplyOption) error {\
	return c.MockApply(ctx, obj, opts...)\
}' vendor/github.com/crossplane/crossplane-runtime/v2/pkg/test/fake.go

echo "✅ Fixed crossplane-runtime MockClient Apply method"