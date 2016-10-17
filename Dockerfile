FROM golang:latest

#RUN apt-get update && apt-get install -qy libsystemd-dev

WORKDIR /usr/src/bakery-agent
ENV "GOPATH=/go:/usr"
RUN mkdir -p /usr/src/bakery-agent
COPY *.go /usr/src/bakery-agent
RUN go get && \
rm -rf /usr/src/bakery-agent

VOLUME /usr/src/bakery-agent
CMD go build
