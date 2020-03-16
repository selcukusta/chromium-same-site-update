FROM golang:alpine3.10 as builder
RUN apk add --update && \
    apk add git
WORKDIR $GOPATH/src/selcukusta/same-site-provider/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/main .
COPY templates/*.html /go/bin/templates/
FROM alpine:3.11.3 as final
RUN apk add --update && \
    apk add bash
COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/bin/templates /go/bin/templates
CMD ["/go/bin/main"]
