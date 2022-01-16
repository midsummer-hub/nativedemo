FROM golang:1.17 AS build
WORKDIR /httpdemo/
COPY . .
ENV CGO_ENABLE=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \

RUN GOOS=linux go build -o httpdemo main.go

FROM busybox
COPY --from=build /httpdemo/httpdemo /httpdemo/httpdemo
EXPOSE 8080
WORKDIR /httpdemo/
ENTRYPOINT ["./httpdemo"]
