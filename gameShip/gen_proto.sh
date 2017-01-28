echo "gen go proto"
protoc -I gameShipRpc/  gameShipRpc/gameShipRpc.proto  --go_out=plugins=grpc:gameShipRpc
