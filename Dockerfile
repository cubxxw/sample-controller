FROM golang
WORKDIR /go/src/k8s.io/sample-controller
COPY . .
CMD ["/go/bin/sample-controller"]
