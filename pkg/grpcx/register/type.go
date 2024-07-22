package register

import "google.golang.org/grpc"

type GRPCxServer interface {
	Register(registrar grpc.ServiceRegistrar)
}
