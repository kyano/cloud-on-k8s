# Build the manager binary
FROM golang:1.11 as builder

# Copy in the go src
WORKDIR /go/src/github.com/elastic/k8s-operators/operators
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o elastic-operator github.com/elastic/k8s-operators/operators/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o keystore-updater github.com/elastic/k8s-operators/operators/cmd/keystore-updater

# Copy the controller-manager into a thin image
FROM ubuntu:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/elastic/k8s-operators/operators/elastic-operator .
COPY --from=builder /go/src/github.com/elastic/k8s-operators/operators/keystore-updater .
ENTRYPOINT ["./elastic-operator"]
CMD ["manager"]