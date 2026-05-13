package gapi

import (
	"fmt"

	db "github.com/dekalor/simple-bank/db/sqlc"
	"github.com/dekalor/simple-bank/pb"
	"github.com/dekalor/simple-bank/token"
	"github.com/dekalor/simple-bank/utils"
)

// Server serves gRPC request for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
