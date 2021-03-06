# Builder

FROM golang:1.18-alpine AS builder
RUN apk add --update make git curl

ARG GITHUB_PATH=github.com/inqast/saga-order

COPY Makefile /home/${GITHUB_PATH}/Makefile
COPY go.mod /home/${GITHUB_PATH}/go.mod
COPY go.sum /home/${GITHUB_PATH}/go.sum

WORKDIR /home/${GITHUB_PATH}

COPY . /home/${GITHUB_PATH}
RUN make build-reserve

# Service

FROM alpine:latest as server
ARG GITHUB_PATH=github.com/inqast/saga-order
LABEL org.opencontainers.image.source=https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/service .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/reserve/ ./migrations/reserve/

RUN chown root:root service

CMD ["./service"]
