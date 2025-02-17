FROM --platform=$BUILDPLATFORM golang:1.23 AS builder
ARG TARGETOS TARGETARCH

# Set some shell options for using pipes and such.
SHELL [ "/bin/bash", "-euo", "pipefail", "-c" ]

# Copy necessary 'go.mod' and 'go.sum' files for separate Go module downloads.
WORKDIR /go/src/go.jlucktay.dev/my-github-repos
COPY go.* .

# Download Go modules in a separate step before adding the source code, to prevent invalidation of cached Go modules if
# only our source code is changed and not any dependencies.
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  GOOS=$TARGETOS GOARCH=$TARGETARCH go mod download

# Copy in all of the source code.
COPY . .

# Compile! With the '--mount' flags below, Go's build cache is kept between builds.
# https://github.com/golang/go/issues/27719#issuecomment-514747274
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
  GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
  -ldflags="-X 'go.jlucktay.dev/version.builtBy=Docker'" -trimpath -v -o /bin/my-github-repos

FROM gcr.io/distroless/base:nonroot AS deployable
USER 65532

# Bring binary over.
COPY --from=builder /bin/my-github-repos /bin/my-github-repos

VOLUME /workdir
WORKDIR /workdir

ENTRYPOINT [ "/bin/my-github-repos" ]
