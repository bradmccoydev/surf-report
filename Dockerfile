FROM golang:1.20.4-alpine3.16 AS builder

ENV CGO_ENABLED=0
ARG REVISION
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /workspace/surf-report -ldflags "-X github.com/bradmccoydev/surf-report/pkg/version.REVISION=${REVISION}" ./ 

FROM gcr.io/distroless/static AS production

LABEL org.opencontainers.image.source="https://github.com/bradmccoydev/surf-report" \
    org.opencontainers.image.url="https://avatars.githubusercontent.com/u/25817813?v=4" \
    org.opencontainers.image.title="Surf Report" \
    org.opencontainers.image.vendor='bradmccoydev' \
    org.opencontainers.image.licenses='Apache-2.0' \
    org.opencontainers.image.description='Surf Report'

WORKDIR /
COPY --from=builder /workspace/surf-report .

USER 65532:65532

ENTRYPOINT ["/surf-report"]