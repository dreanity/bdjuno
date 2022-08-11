FROM golang:1.17-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/dreanity/bdjuno
COPY . ./
RUN go mod download
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/dreanity/bdjuno/build/bdjuno /usr/bin/bdjuno
CMD [ "bdjuno" ]
