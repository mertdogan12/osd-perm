FROM golang:latest as build-deps

ARG ssh_prv_key
ARG ssh_pub_key

RUN apt update; \
    apt install -y \
        git \
        openssh-server

# Authorize SSH Host
RUN mkdir -p /root/.ssh; \
    chmod 0700 /root/.ssh; \
    ssh-keyscan -t rsa github.com >> /root/.ssh/known_hosts

# Add the keys and set permissions
RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa; \
    echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub; \
    chmod 600 /root/.ssh/id_rsa; \
    chmod 600 /root/.ssh/id_rsa.pub

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

WORKDIR /go/src/github.com/mertdogan12/osd-perm

ENV GOPRIVATE=github.com/mertdogan12/osd
ENV GOPATH=/go

COPY . ./
RUN go get -d ./...
RUN go build .

RUN rm -rf /root/.ssh

FROM ubuntu:latest

WORKDIR /app

COPY --from=build-deps /go/src/github.com/mertdogan12/osd-perm ./

EXPOSE 80
