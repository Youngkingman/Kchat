FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/cespare/reflex

COPY . .

RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8000
CMD ["./run"]