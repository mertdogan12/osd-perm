FROM golang:latest as build-deps

ARG github_username
ARG github_token

RUN echo "machine github.com login $github_username password $github_token" | cat > /root/.netrc
RUN echo "machine github.com login $github_username password $github_token"

WORKDIR /go/src/github.com/mertdogan12/osd-perm

ENV GOPRIVATE=github.com/mertdogan12/osd
ENV GOPATH=/go

COPY . ./
RUN go get -d ./...
RUN go build .

RUN rm -rf /root/.ssh

FROM ubuntu:latest

WORKDIR /app

COPY --from=build-deps /go/src/github.com/mertdogan12/osd-perm/osd-perm ./

EXPOSE 80

CMD ["/app/osd-perm"]
