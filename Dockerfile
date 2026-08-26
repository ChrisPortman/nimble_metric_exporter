ARG BUILDBASE=alpine
ARG RUNTIMEBASE=$BUILDBASE

FROM docker.io/golang:1.26-${BUILDBASE} AS build

ARG VERSION="develop"
ARG VERSION_PKG="github.com/ChrisPortman/nimble_metric_exporter/internal/version"

WORKDIR /src
RUN mkdir /build

COPY go.mod go.sum ./
RUN go mod download

RUN go env -w GOFLAGS=-buildvcs=false
RUN --mount=type=bind,target=. export DATE=$(date) && \
     export DATE=$(date) && \
     export GOVERSION=$(go version) && \
     export PREFIX="${VERSION_PKG}" && \
     export LDFLAGS="-s -w" && \
     export LDFLAGS="$LDFLAGS -X '${PREFIX}.Version=${VERSION}'" && \
     export LDFLAGS="$LDFLAGS -X '${PREFIX}.BuildDate=${DATE}'" && \
     export LDFLAGS="$LDFLAGS -X '${PREFIX}.GoVersion=${GOVERSION}'" && \
     echo "Version: ${VERSION}" && \
     echo ${LDFLAGS} && \
     go build -o /build -trimpath -v -ldflags "${LDFLAGS}" ./...

# The gateway image
FROM ${RUNTIMEBASE}:latest
COPY --from=build /build/nimble_exporter /
ENTRYPOINT ["/nimble_exporter"]
