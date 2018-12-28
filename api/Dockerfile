FROM golang:alpine as build-env

RUN mkdir -p /api
WORKDIR /api
COPY go.mod . 
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api

FROM scratch 
COPY --from=build-env /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]
