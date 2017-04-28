# Go Report Card Lite

GoReportCard Lite, a slim version of original [Go Report Card](https://github.com/gojp/goreportcard/) by [Shawn Smith](https://twitter.com/shawnps) and [Herman Schaaf](https://twitter.com/ironzeb) which is catered to be integrated into CI pipelines and can generate reports using local/private repositories. 

A web application that generates a report on the quality of go project. It uses several measures, including `gofmt`, `go vet`, `go lint` and `gocyclo`.

### Installation

Assuming you already have a recent version of Go installed, pull down the code with `go get`:

```
go get github.com/nearmap/goreportcardlite
cd $GOPATH/src/github.com/nearmap/goreportcardlite
make install
```

Now run

```
make start
```

URL for report card
```
curl http://localhost:8000/report?repo=<full_path_to_local_repository>
```


### Reference 

```sh
docker build -t nearmap/goreportcard .
```

```sh
docker run -ti -v /Users/suneeta.mall/Documents/Wks/go/src/github.nearmap.com/hyperweb/authproxy.git:/go/src/github.nearmap.com/hyperweb/authproxy.git/ nearmap/goreportcard github.nearmap.com/hyperweb/authproxy.git
```

