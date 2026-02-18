FROM golang:latest AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o app .


FROM alpine:latest AS production-stage
WORKDIR /app
COPY --from=build-stage /app/app .
COPY --from=build-stage /app/config.yaml .
EXPOSE 8080
# 可选：安装 bash 或其他工具
#RUN apk add --no-cache bash
ENTRYPOINT ["./app"]
