# osd-perm

OSD permission backend

## Building

```bash
# You need access to the repo: mertdogan12/osd

# Go build
go build .

# Docker build
docker build -t ghcr.io/mertdogan12/osd-perm:latest --build-arg github_username=$GIHUB_USERNAME --build-arg github_token=$GIHUB_TOKEN .
```

## Install

```bash
go install github.com/mertdogan12/osd-perm:latest
```

## Use it

```bash
# Default port 80
osu-perm -p <port>
```

## [API Dokumentation](https://documenter.getpostman.com/view/14220165/UyrBhvb9)
