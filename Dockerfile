
ARG BASE=scratch

FROM golang:1.22.4-alpine3.19 AS build

ARG COMMIT=
ARG VERSION=
ARG RELEASE=false

ENV GOBIN=/build/bin

WORKDIR /build

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN <<EOF
#!/usr/bin/env sh
set -o errexit
set -o errtrace
set -o nounset
set -o pipefail
# Uncomment for debugging purpose
# set -o xtrace

GO_LDFLAGS=""
if [ "$RELEASE" = "true" ]; then
  GO_LDFLAGS="$GO_LDFLAGS -s -w"
fi
if [[ -n "$COMMIT" ]]; then
  GO_LDFLAGS="$GO_LDFLAGS -X=main.GitCommit=$COMMIT"
fi
if [[ -n "$VERSION" ]]; then
  GO_LDFLAGS="$GO_LDFLAGS -X=main.Version=$VERSION"
fi

go install \
  -ldflags "$GO_LDFLAGS" \
  -trimpath \
  ./...
EOF

# Build the final image
FROM $BASE
USER 10000:10000

COPY --from=build /build/bin/todo /usr/bin/
ENV HTTP_SERVER__ADDR_PORT=0.0.0.0:8080
EXPOSE 8080

ENTRYPOINT ["/usr/bin/todo"]