docker build -t okteto/golang:1.17-arm -f okteto-golang-arm.dockerfile .
kind load docker-image okteto/golang:1.17-arm