ARG GO_VERSION=1.21.4
ARG TARGETOS TARGETARCH
FROM golang:${GO_VERSION}-alpine AS go-plugin
LABEL stage=go-plugin

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk add curl make
RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod \
    PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
    curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip && \
    unzip -o protoc-25.1-linux-x86_64.zip -d /go/bin && \
    go install github.com/zeromicro/go-zero/tools/goctl@v1.6.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1  && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1 && \
    go install github.com/envoyproxy/protoc-gen-validate@v1.0.4

ENV PATH $PATH:/go/bin/bin


FROM go-plugin AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux


WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  go mod download
COPY . .
COPY ./etc /app/etc
RUN pwd && ls -l
#RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  which goctl
#RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  which protoc
#RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  ls -al /go/bin/
#RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  ls -al /go/bin/bin
#RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  ls -al /go/bin/bin/protoc
RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  make gen
RUN --mount=type=cache,target=/go/bin --mount=type=cache,target=/go/pkg/mod  go build -ldflags="-s -w" -o /app/app-oracleX oracleX.go


FROM alpine
RUN apk update --no-cache && apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
RUN ls -l
COPY --from=builder /app/app-oracleX /app/app-oracleX
COPY --from=builder /app/etc /app/etc
RUN ls -l
CMD ["./app-oracleX", "-f", "etc/oracleX.yaml"]
