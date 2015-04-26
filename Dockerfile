FROM golang:1.4.2-onbuild
RUN ln /go/bin/app /go/bin/media-controller
EXPOSE 2222
CMD media-controller -l :2222 server
