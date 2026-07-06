# syntax=docker/dockerfile:1

ARG GO_VERSION=1.26.5

# ---- build stage ----
# BUILDPLATFORM pins this stage to the host arch so the Go compiler itself
# always runs natively — no QEMU emulation needed even when cross-compiling.
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS builder

ARG VCS_REF
ARG VERSION=dev
ARG TARGETOS
ARG TARGETARCH

# Install mise (pinned tool versions + codegen tasks)
RUN curl -fsSL https://mise.run | sh
ENV PATH="/root/.local/bin:/root/.local/share/mise/shims:${PATH}"

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY mise.toml ./
RUN mise trust && mise install

COPY . .

RUN mise run generate

# Cross-compile for whatever platform buildx is targeting
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -trimpath \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${VCS_REF}" \
    -o /out/server \
    ./cmd/server

# ---- final stage ----
FROM alpine:3.20

RUN apk add --no-cache ca-certificates \
    && addgroup -S app && adduser -S app -G app \
    && mkdir -p /data/conf && chown -R app:app /data/conf

WORKDIR /app
COPY --from=builder --chown=app:app /out/server /app/server

ARG VCS_REF
ARG VERSION=dev
LABEL org.opencontainers.image.revision=${VCS_REF} \
      org.opencontainers.image.version=${VERSION}

USER app

EXPOSE 8000 9000
VOLUME /data/conf

ENTRYPOINT ["./server"]
CMD ["-conf", "/data/conf"]
