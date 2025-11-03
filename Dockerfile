# Build the provider binary
FROM golang:1.25.3 AS build

WORKDIR /workspace

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY apis/ apis/
COPY cmd/ cmd/
COPY internal/ internal/
COPY hack/ hack/

# Build the provider binary (minimal version with DNS and Zone support only)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o provider ./cmd/provider/

# Use distroless as minimal base image
FROM gcr.io/distroless/static:nonroot
WORKDIR /

# Copy controller binary
COPY --from=build /workspace/provider /provider

# Copy package metadata to the image root so Crossplane can find it
COPY package/crossplane.yaml /package.yaml

# Copy CRDs into a dedicated folder that the package will expose
# IMPORTANT: use trailing slashes to copy directory contents
COPY package/crds/ /crds/

USER 65532:65532
ENTRYPOINT ["/provider"]