FROM golang:1.19.5-alpine AS build

WORKDIR /go/src/github.com/jonggulee/search-job
COPY . .

RUN go get -d -v ./... \
&& CGO_ENABLED=0 go build -a -installsuffix cgo -o searchjob .

ENTRYPOINT ./searchjob
