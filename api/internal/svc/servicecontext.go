package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformer"
)

type ServiceContext struct {
	Config    config.Config
	Transform transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Transform: transformer.NewTransformer(zrpc.MustNewClient(c.Transformer)),
	}
}
