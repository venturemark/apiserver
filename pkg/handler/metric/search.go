package metric

import (
	"context"
	"net"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *metric.SearchI) (*metric.SearchO, error) {
	{
		var err error
		var res []string

		{
			host := "google.com"

			h.logger.Log(ctx, "level", "info", "message", "resolving host", "host", host)
			res, err = net.LookupHost(host)
			if err != nil {
				h.logger.Log(ctx, "level", "error", "message", "resolving failed", "stack", tracer.JSON(err), "host", host)
			}

			for _, a := range res {
				h.logger.Log(ctx, "level", "info", "message", "resolved host", "host", host, "address", a)
			}
		}
		{
			host := "www.google.com"

			h.logger.Log(ctx, "level", "info", "message", "resolving host", "host", host)
			res, err = net.LookupHost(host)
			if err != nil {
				h.logger.Log(ctx, "level", "error", "message", "resolving failed", "stack", tracer.JSON(err), "host", host)
			}

			for _, a := range res {
				h.logger.Log(ctx, "level", "info", "message", "resolved host", "host", host, "address", a)
			}
		}
		{
			host := "redis-master.infra.svc.cluster.local"

			h.logger.Log(ctx, "level", "info", "message", "resolving host", "host", host)
			res, err = net.LookupHost(host)
			if err != nil {
				h.logger.Log(ctx, "level", "error", "message", "resolving failed", "stack", tracer.JSON(err), "host", host)
			}

			for _, a := range res {
				h.logger.Log(ctx, "level", "info", "message", "resolved host", "host", host, "address", a)
			}
		}
		{
			host := "cluster.local"

			h.logger.Log(ctx, "level", "info", "message", "resolving host", "host", host)
			res, err = net.LookupHost(host)
			if err != nil {
				h.logger.Log(ctx, "level", "error", "message", "resolving failed", "stack", tracer.JSON(err), "host", host)
			}

			for _, a := range res {
				h.logger.Log(ctx, "level", "info", "message", "resolved host", "host", host, "address", a)
			}
		}
	}

	// Search for any metric associated with the given updates. One or many
	// update IDs may be provided.
	{
		ok, err := h.storage.Metric.Search.Non.Timeline.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Metric.Search.Non.Timeline.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
