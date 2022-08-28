FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod .
COPY go.sum .
# RUN go get github.com/cespare/reflex
RUN go mod download

COPY . .

RUN go build -o kchat .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY . .
# 需要执行权限
RUN chmod +x ./wait-for-it.sh
#Copy executable from builder
COPY --from=builder /go/src/app/kchat .

EXPOSE 8080

CMD [ "./kchat" ]