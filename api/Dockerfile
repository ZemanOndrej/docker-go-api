FROM golang

ADD . /go/src/docker_api/src/api
WORKDIR /go/src/docker_api/src/api
RUN go get
RUN go install

ENTRYPOINT /go/bin/api

EXPOSE 8080