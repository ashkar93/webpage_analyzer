FROM golang:alpine as build-env

RUN mkdir /api
WORKDIR /api
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .

# Unit Tests
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v ./services
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v ./controllers

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/main

FROM alpine 
COPY --from=build-env /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"]