# 编译镜像
FROM golang:1.18 as builder

WORKDIR /go/src/test

COPY . .

RUN go build -mod=vendor -o test test.go

# 运行镜像
FROM alpine

WORKDIR /root

COPY --from=builder /go/src/test/test  .

EXPOSE 8080

ENTRYPOINT ["./test"]

