FROM golang:1.4.2-onbuild
EXPOSE 2222
CMD media-controller -s :2222 server
