# Gather build dependencies
FROM golang:1.12.9-stretch as builder_base

WORKDIR /go/src/github.com/bpmericle/go-webservice

# Force the go compiler to use modules 
ENV GO111MODULE=on
# Disable cgo - this makes static binaries that will work on an Alpine image
ENV CGO_ENABLED=0

# Handle go modules setup
COPY go.mod .
COPY go.sum .
RUN go mod download

# Unit test and intall the application
FROM builder_base AS builder

# Copy in source
COPY . .

# Run the tests
RUN go test -short -v ./...

# Install binary
RUN go install

# Final small image
FROM alpine:3.10

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/go-webservice /bin/go-webservice
ENTRYPOINT ["/bin/go-webservice"]