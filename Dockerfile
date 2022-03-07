FROM golang:1.17

WORKDIR /app
COPY . .
ARG GIT_ACCESS_TOKEN
ENV GOPRIVATE=gitlab.booking.com/go
RUN git config --global url."https://gitlab-ci-token:$GIT_ACCESS_TOKEN@gitlab.booking.com/".insteadOf "https://gitlab.booking.com/"
RUN http_proxy=http://webproxy:3128 https_proxy=http://webproxy:3128 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o promviz ./cmd/promviz/main.go

FROM alpine:3.12

RUN http_proxy=http://webproxy:3128 https_proxy=http://webproxy:3128 apk add --no-cache ca-certificates && update-ca-certificates

ADD promviz.yaml .
COPY --from=0 /app/promviz .

EXPOSE 8080

CMD ["./promviz", "--config.file=./promviz.yaml", "--storage.path=/tmp", "--storage.retention=1h", "--api.port=8080"]
