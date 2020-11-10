package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := mainE(context.Background())
	if err != nil {
		tracer.Panic(err)
	}
}

func mainE(ctx context.Context) error {
	var err error

	var l net.Listener
	{
		l, err = net.Listen("tcp", fmt.Sprintf(":%d", 7777))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var a metric.APIServer
	{
		g := grpc.NewServer()
		reflection.Register(g)

		a = &API{}
		metric.RegisterAPIServer(g, a)

		err := g.Serve(l)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

// -------------------------------------------------------------------------- //

type API struct {
	metric.UnimplementedAPIServer
}

func (a *API) Create(ctx context.Context, cre *metric.CreateI) (*metric.CreateO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.CreateO{}, nil
}

func (a *API) Delete(ctx context.Context, del *metric.DeleteI) (*metric.DeleteO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.DeleteO{}, nil
}

func (a *API) Search(ctx context.Context, sea *metric.SearchI) (*metric.SearchO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.SearchO{}, nil
}

func (a *API) Update(ctx context.Context, upd *metric.UpdateI) (*metric.UpdateO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.UpdateO{}, nil
}
