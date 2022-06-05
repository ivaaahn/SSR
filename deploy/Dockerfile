FROM golang:1.18-alpine as builder
WORKDIR /build

COPY go.mod .
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 go build -o ssr ./cmd/app/main.go

FROM scratch

COPY --from=builder /build/ssr /ssr
COPY --from=builder /build/config/config.yml /config.yml

ENV DEPLOY_MODE=PROD
ENTRYPOINT ["/ssr"]
