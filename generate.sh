# 生成model部分代码
goctl model mysql ddl -src ./model/sql/user.sql -dir ./model --home common/goctl


goctl rpc protoc rpc/user.proto -m --go_out=rpc --go-grpc_out=rpc --zrpc_out=rpc --home common/goctl


goctl api go -api ./api/user.api -dir ./api --home common/goctl