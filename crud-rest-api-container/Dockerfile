# Stage 1 - Build the REST API Go Binary
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /
COPY flinks-crud-rest-api/flinks-crud-rest-api.go /
RUN go get -d -v
RUN go build -o /go/bin/flinks-crud-rest-api

# Stage 2 - Copy the compiled binary into a clean Alpine container for greatly reduced container size
FROM alpine:3.12

COPY --from=builder /go/bin/flinks-crud-rest-api /flinks-crud-rest-api
ENTRYPOINT ["/flinks-crud-rest-api"]
