# todo-go

This is todo site sample for backend.


## ■ Using with
- docker
- go >= 1.20.0
- proto v3
- gRPC


## ■ Install Requirements
With the assumption that Go and Docker is installed.

### ① Install protobuf
```
% brew install protobuf
```

### ② Install some requirements referred with [here](https://grpc.io/docs/languages/go/quickstart/)
```
% go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
% go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

% vi ~/.zprofile 
--------
# Add this
export PATH="$PATH:$(go env GOPATH)/bin"% 
--------
```

### ③ Build Docker containers
```
% make up
```

### ④ If you want to run all services on each containers, run below
```
% make run-all
```


## ■ If you change some proto files, run below
```
% make proto-gen
```

## ■ Following is build for appmixer that is used as api-gateway
```
% make proto-gw
```

## ■ What I referred as study and constructed this app
- [grpc](https://dev.classmethod.jp/articles/golang-grpc-sample-project/)
- [protoc cmd](https://qiita.com/maaaashin324/items/b8d3c5c016203dce2d6a)
- [gRPC-gateway](https://note.com/leslesnoa/n/n1888381e068b)
  - めっちゃはまったけど、解決してくれた記事
    - [記事①](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/)
    - [記事②](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/#:~:text=%E3%81%A7%E3%82%B9%E3%82%BF%E3%83%96%E3%82%92%E7%94%9F%E6%88%90%E3%81%99%E3%82%8B%E5%89%8D%E3%81%ABprotoc%E3%80%81%E3%81%84%E3%81%8F%E3%81%A4%E3%81%8B%E3%81%AE%E4%BE%9D%E5%AD%98%E9%96%A2%E4%BF%82%E3%82%92%20proto%20%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E6%A7%8B%E9%80%A0%E3%81%AB%E3%82%B3%E3%83%94%E3%83%BC%E3%81%99%E3%82%8B%E5%BF%85%E8%A6%81%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82%E3%81%AE%E3%82%B5%E3%83%96%E3%82%BB%E3%83%83%E3%83%88%E3%82%92%E5%85%AC%E5%BC%8F%E3%83%AA%E3%83%9D%E3%82%B8%E3%83%88%E3%83%AAgoogleapis%E3%81%8B%E3%82%89%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E3%81%AE%20proto%20%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E6%A7%8B%E9%80%A0%E3%81%AB%E3%82%B3%E3%83%94%E3%83%BC%E3%81%97%E3%81%BE%E3%81%99%E3%80%82%E3%81%9D%E3%81%AE%E5%BE%8C%E3%80%81%E6%AC%A1%E3%81%AE%E3%82%88%E3%81%86%E3%81%AB%E3%81%AA%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)