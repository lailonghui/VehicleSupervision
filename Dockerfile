# 构建前端
FROM node:12.19-alpine as node-build
WORKDIR /app
ENV SASS_BINARY_SITE https://npm.taobao.org/mirrors/node-sass/
COPY ./html/package.json .
RUN npm install --registry=https://registry.npm.taobao.org
COPY ./html .
RUN npm run build

# 构建golang
FROM golang:alpine AS go-build
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
# 删除静态文件并构建出可执行文件
RUN  rm -rf html && go build -o app


# 将golang可执行文件放到空的alpine
FROM alpine:latest AS production
# 设置时区为上海的时区
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
WORKDIR /app
# 复制可执行文件
COPY --from=go-build /app/app ./app
# 复制静态文件
COPY --from=node-build /app/dist ./html
EXPOSE 8080
ENTRYPOINT ["./app"]
