FROM golang:alpine as build-env

RUN mkdir -p /api
WORKDIR /api
COPY go.mod . 
COPY go.sum .

RUN apk add --no-cache git  \
    && go mod download \
    && apk del git 

COPY src/ .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /api/main

FROM scratch 
COPY --from=build-env /api/main /api/main
ENTRYPOINT ["/api/main"]
