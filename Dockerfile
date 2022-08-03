FROM golang:1.18-alpine

WORKDIR /go-utils

COPY ./ ./

CMD ["go","test","-v","./..."]