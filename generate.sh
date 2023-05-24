# 生成model部分代码
goctl model mysql ddl -src ./model/sql/user.sql -dir ./model


goctl rpc protoc rpc/user.proto --go_out=rpc --go-grpc_out=rpc --zrpc_out=rpc


goctl api go -api ./api/user.api -dir ./api