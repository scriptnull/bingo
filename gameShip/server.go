package main

import (
	"net"

	"github.com/scriptnull/bingo/gameShip/action"
	"github.com/scriptnull/bingo/gameShip/gameShipRpc"
	"github.com/scriptnull/bingo/gameShip/store"
	"github.com/uber-go/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {

	var logger = zap.New(
		zap.NewTextEncoder(),
		zap.InfoLevel,
	)

	var server = &action.GameShipRPCServer{
		Logger:    logger,
		GameStore: &store.LocalGameStore{},
	}

	logger.Info("Booting gameShip")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("Failed to listen:", zap.Error(err))
	}

	s := grpc.NewServer()
	gameShipRpc.RegisterGameShipRpcServer(s, server)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve:", zap.Error(err))
	}
}
