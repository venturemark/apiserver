package handler

import "google.golang.org/grpc"

type Interface interface {
	Attach(g *grpc.Server)
}
