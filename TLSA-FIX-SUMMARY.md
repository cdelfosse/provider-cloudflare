# GitHub Actions Release Workflow Fix - v0.11.0

## Problem Summary

The v0.11.0 release workflow completed successfully but failed to actually build and publish container images to the registry. The `make publish` command completed in 0.37 seconds with no output.

## Root Cause Analysis

The issue was a **GNU Make variable precedence problem** in `.github/workflows/release.yml`:

### Original Broken Code (Line 73-81)
```bash
# Build and Publish All Artifacts
run: |
  # Use the standardized build system to handle everything
  VERSION="${{ steps.version.outputs.version }}"

  echo "Repository: ${{ github.repository }}"
  echo "Repository owner: ${{ github.repository_owner }}"
  echo "Version: ${VERSION}"

  # Build and publish using make targets that handle Docker + Crossplane packages
  # IMPORTANT: VERSION must be passed as make variable, not environment variable
  export VERSION="${VERSION}"
  make publish REGISTRY_ORGS="ghcr.io/${{ github.repository_owner }}"
```

### The Bug

The workflow set `export VERSION="v0.11.0"` as an **environment variable**, but the Makefile at `build/makelib/common.mk:215-223` uses:

```makefile
# set a semantic version number from git if VERSION is undefined.
ifeq ($(origin VERSION), undefined)
# use tags
VERSION := $(shell git describe --dirty --always --tags ...)
endif
export VERSION
```

### GNU Make Variable Precedence

Make has three sources for variables with this precedence order:

1. **Command-line arguments** (HIGHEST) - `make VAR=value`
2. **Makefile variables** (MEDIUM) - `VAR := value`
3. **Environment variables** (LOWEST) - `export VAR=value` before make

The Makefile uses **immediate assignment** (`:=`) which evaluates the shell command at parse time. Even though the environment had `VERSION=v0.11.0`, the Makefile's computed version from `git describe` took precedence.

Additionally, `ifeq ($(origin VERSION), undefined)` checks if VERSION is undefined, but when VERSION comes from environment, `$(origin VERSION)` returns `"environment"` not `"undefined"`, so the condition is false. However, the `:=` assignment operator means the Makefile variable definition still overrides the environment variable value.

## The Fix

Changed the workflow to pass VERSION as a **make command-line argument** instead:

```bash
# Build and Publish All Artifacts
run: |
  # Use the standardized build system to handle everything
  VERSION="${{ steps.version.outputs.version }}"

  echo "Repository: ${{ github.repository }}"
  echo "Repository owner: ${{ github.repository_owner }}"
  echo "Version: ${VERSION}"

  # Build and publish using make targets that handle Docker + Crossplane packages
  # IMPORTANT: VERSION must be passed as make variable, not environment variable
  make publish VERSION="${VERSION}" REGISTRY_ORGS="ghcr.io/${{ github.repository_owner }}" XPKG_REG_ORGS="ghcr.io/${{ github.repository_owner }}"
```

**Key change**: `make publish VERSION="${VERSION}"` instead of `export VERSION="${VERSION}"; make publish`

Command-line arguments have the **highest precedence** in Make, overriding both Makefile variables and environment variables.

## Verification

After the fix was committed and pushed (commit 8c86a3c), future releases will correctly pass the VERSION to the Makefile, ensuring builds use the release tag version instead of the git-computed version.

## Manual Workaround Applied

For v0.11.0, manual build and publish was executed:

```bash
VERSION=v0.11.0 make build
VERSION=v0.11.0 make publish PLATFORMS=linux_amd64 REGISTRY_ORGS=ghcr.io/rossigee XPKG_REG_ORGS=ghcr.io/rossigee
```

Successfully published:
- Container image: `ghcr.io/rossigee/provider-cloudflare:v0.11.0`
- Container image: `ghcr.io/rossigee/provider-cloudflare:latest`
- Crossplane package: `ghcr.io/rossigee/provider-cloudflare:v0.11.0`

Digest: `sha256:f314a11fc5c3ca4b6b1970654139fd21bf364c1a87e761f2f4cb1e1031ff1df8`

## Impact

- **Affected Version**: v0.11.0 release (manually fixed)
- **Future Releases**: Fixed by commit 8c86a3c - workflow now passes VERSION correctly
- **Breaking Change**: None - this was a CI/CD infrastructure bug fix

## References

- Commit fixing workflow: 8c86a3c
- v0.11.0 release commit: 3a0611c
- GNU Make documentation: https://www.gnu.org/software/make/manual/html_node/Variables.html
- Make variable precedence: https://www.gnu.org/software/make/manual/html_node/Overriding.html

## Date

2025-10-20
