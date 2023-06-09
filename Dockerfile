FROM golang:alpine as build
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get ./...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -tags timetzdata -o gg main.go

FROM scratch
COPY --from=build /go/src/app/gg /usr/bin/gg
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 1848
#vangogh artifacts: checksums, images, metadata, recycle_bin, videos
VOLUME /var/lib/vangogh

ENTRYPOINT ["/usr/bin/gg"]
CMD ["serve","-port", "1848", "-stderr"]