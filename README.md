# Go Report Card Lite

GoReportCard Lite, a slim version of original [Go Report Card](https://github.com/gojp/goreportcard/) by [Shawn Smith](https://twitter.com/shawnps) and [Herman Schaaf](https://twitter.com/ironzeb) which is catered to be integrated into CI pipelines and can generate reports using local/private repositories. 

A web application that generates a report on the quality of go project. It uses several measures, including `gofmt`, `go vet`, `go lint` and `gocyclo`.

### Quick start
If you have Docker installed and want to analyze a local project, just navigate to that project and run:

```sh
curl -s https://raw.githubusercontent.com/nearmap/goreportcardlite/master/analyze_current_directory.sh | bash
```

or, you can save [this script](analyze_current_directory.sh) to your PATH, and run it, still from the project you want to analyze.

### Installation

Assuming you already have a recent version of Go installed, pull down the code with `go get`:

```sh
go get github.com/nearmap/goreportcardlite
cd $GOPATH/src/github.com/nearmap/goreportcardlite
make install
## To run
make start
```


URL for report card
```sh
curl http://localhost:8000/report?repo=<full_path_to_local_repository>
```


### Docker 
Build docker image
```sh
docker build -t nearmap/goreportcard .
```

Generate report with docker, run the following which will generate 'goreportcard.htm' report card and place it in mounted volume i.e. github.com/nearmap/goreportcardlite/goreportcard.htm:
```sh
docker run -ti -v /Users/suneeta.mall/Documents/Wks/go/src/github.com/nearmap/goreportcardlite:/go/src/github.com/nearmap/goreportcardlite/ nearmap/goreportcard github.com/nearmap/goreportcardlite
```

Public docker image is avaialble @ dockerhub 
```
#https://hub.docker.com/r/nearmap/goreportcard/
docker pull nearmap/goreportcard
```
