#用于拉取镜像
FROM golang:1.22.1 As build

#D
RUN mkdir /app

COPY ./ /app
#进到目录
WORKDIR /app

RUN go mod tidy
#打包二进制文件
RUN go build main.go
##拉取精简镜像
#FROM scratch
##复制上层镜像打包出来的二进制文件
#COPY --from=build /app/main /main

CMD ["./main"]