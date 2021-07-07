package service

import (
	db "github.com/like-org/common/app/service/account/store"
)

type Server struct {
	store db.Store
}

func NewServer(store db.Store) *Server {
	return &Server{store}
}
