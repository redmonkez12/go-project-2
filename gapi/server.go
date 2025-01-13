package gapi

import (
	"fmt"

	db "github.com/redmonkez12/go-project-2/db/sqlc"
	"github.com/redmonkez12/go-project-2/pb"
	"github.com/redmonkez12/go-project-2/token"
	"github.com/redmonkez12/go-project-2/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maket %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}


	return server, nil
}
