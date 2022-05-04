package service

import (
	"context"
	"golang-with-kafka/internal/entity"
	"log"
)

type Service interface {
	Call(context.Context, *entity.Request) (interface{}, error)
}

type service struct {
	log log.Logger
}

func NewService() {
	return
}

func (s *service) Call(ctx context.Context, request *entity.Request) (interface{}, error) {
	return nil, nil
}
