# Gather build dependencies
FROM golang:1.13-stretch as builder_base

WORKDIR /go/src/github.com/bpmericle/go-webservice

# Force the go compiler to use modules 
ENV GO111MODULE=on
# Disable cgo - this makes static binaries that will work on an Alpine image
ENV CGO_ENABLED=0
# Use a proxy, instead of pulling direct from source
#ENV GOPROXY=https://proxy.golang.org,direct
ENV GOPROXY=https://gocenter.io,direct

# Handle go modules setup
COPY go.mod .
COPY go.sum .
RUN go mod download

# Unit test and intall the application
FROM builder_base AS builder

# Copy in source
COPY . .

# Run the tests
RUN go test -failfast -short -v ./...

# Install binary
RUN go install

# Final small image
FROM alpine:3.10

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/go-webservice /bin/go-webservice

ENTRYPOINT ["/bin/go-webservice"]