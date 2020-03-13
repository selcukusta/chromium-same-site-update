FROM golang:alpine3.10 as builder
RUN apk add --update && \
    apk add git
WORKDIR $GOPATH/src/selcukusta/same-site-consumer/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/main .
COPY templates/*.html /go/bin/templates/
RUN ls -laR
FROM alpine:3.11.3 as final
RUN apk add --update && \
    apk add bash
COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/bin/templates /go/bin/templates
EXPOSE 80
CMD ["/go/bin/main"]