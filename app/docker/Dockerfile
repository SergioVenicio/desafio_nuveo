FROM golang:1.18

WORKDIR /usr/src/app
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

ENTRYPOINT ["go", "run", "main.go"]