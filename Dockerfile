#FROM hyperweb/golang

FROM golang:1.7

RUN go get golang.org/x/tools/go/vcs

VOLUME /go/src

COPY . $GOPATH/src/github.com/nearmap/goreportcardlite

RUN apt-get install -yqq curl

COPY generate_report.sh /

RUN chmod +755 /generate_report.sh

ENTRYPOINT ["/generate_report.sh"]