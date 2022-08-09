# grpc
https://grpc.io/docs/languages/go/quickstart/
# Test
test 会改变工作区，导致配置目录找不到，根目录下CONF_PATH=$(pwd)/conf && go test 
当前目录： `go test`  
递归所有目录：`go test ./...`