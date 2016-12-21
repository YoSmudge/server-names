FROM golang:1.7

ENV GOPATH=/app
RUN mkdir -p /app/src/github.com/YoSmudge/server-names &&\
  apt-get update && apt-get install -y wordnet &&\
  wget -q -O /tmp/glide.tar.gz https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-linux-amd64.tar.gz &&\
  cd /tmp && tar -xvf glide.tar.gz && mv linux-amd64/glide /usr/bin && glide -v &&\
  go get -u github.com/jteeuwen/go-bindata/...
WORKDIR /app/src/github.com/YoSmudge/server-names
COPY . .
