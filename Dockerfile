FROM golang:alpine
ADD . /iamgone
WORKDIR /iamgone
ENV GOPROXY=https://goproxy.io
ENTRYPOINT ["go","run","."]