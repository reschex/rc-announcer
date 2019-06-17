FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/rc-announcer
COPY rc-announcer/*.go ./
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/rc-announcer


FROM scratch
WORKDIR /
COPY --from=builder /go/bin/rc-announcer .
EXPOSE 8080:8080

ENTRYPOINT ["./rc-announcer"]