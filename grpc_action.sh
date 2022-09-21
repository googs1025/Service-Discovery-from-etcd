
echo "生成 server端 grpc代码"

OUT=./server/rpc

protoc --go_out=${OUT} --go-grpc_out=${OUT} --go-grpc_opt=require_unimplemented_servers=false protocol.proto

echo "生成 client grpc代码"

OUT=./client/rpc

protoc --go_out=${OUT} --go-grpc_out=${OUT} --go-grpc_opt=require_unimplemented_servers=false protocol.proto