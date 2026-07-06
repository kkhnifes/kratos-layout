// Package server wires HTTP and gRPC servers and registers services.
package server

import (
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"

	"github.com/kkhnifes/kratos-layout/internal/conf"
	"github.com/kkhnifes/kratos-layout/internal/service"
)

// New builds this layer's HTTP and gRPC servers from the registered services.
func New(c *conf.Server, todoSvc *service.TodoService) (*grpc.Server, *http.Server) {
	return NewGRPCServer(c, todoSvc), NewHTTPServer(c, todoSvc)
}
