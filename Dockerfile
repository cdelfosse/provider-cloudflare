# Build the provider binary
FROM golang:1.25.3 AS build
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY apis/ apis/
COPY cmd/ cmd/
COPY internal/ internal/
COPY hack/ hack/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o provider ./cmd/provider/

# Stage to carry CRDs (robust with buildx)
FROM busybox AS crds
WORKDIR /crds
COPY package/crds/ ./   # copies all CRD yamls

# Final minimal runtime image + package metadata
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=build /workspace/provider /provider
COPY package/crossplane.yaml /package.yaml
COPY --from=crds /crds/ /crds/
USER 65532:65532
ENTRYPOINT ["/provider"]