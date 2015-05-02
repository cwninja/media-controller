FROM golang:1.4.2-onbuild
RUN ln /go/bin/app /go/bin/media-controller
EXPOSE 2222
EXPOSE 8081
CMD media-controller --listen :2222 --listen-web :8081 server
