package storage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/audience"
	"github.com/venturemark/apiserver/pkg/storage/message"
	"github.com/venturemark/apiserver/pkg/storage/metric"
	"github.com/venturemark/apiserver/pkg/storage/metupd"
	"github.com/venturemark/apiserver/pkg/storage/texupd"
	"github.com/venturemark/apiserver/pkg/storage/timeline"
	"github.com/venturemark/apiserver/pkg/storage/update"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type Storage struct {
	Audience *audience.Audience
	Message  *message.Message
	Metric   *metric.Metric
	MetUpd   *metupd.MetUpd
	TexUpd   *texupd.TexUpd
	Timeline *timeline.Timeline
	Update   *update.Update
}

func New(config Config) (*Storage, error) {
	var err error

	var audienceStorage *audience.Audience
	{
		c := audience.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		audienceStorage, err = audience.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var messageStorage *message.Message
	{
		c := message.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		messageStorage, err = message.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var metricStorage *metric.Metric
	{
		c := metric.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		metricStorage, err = metric.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var metupdStorage *metupd.MetUpd
	{
		c := metupd.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		metupdStorage, err = metupd.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var texupdStorage *texupd.TexUpd
	{
		c := texupd.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		texupdStorage, err = texupd.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var timelineStorage *timeline.Timeline
	{
		c := timeline.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		timelineStorage, err = timeline.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var updateStorage *update.Update
	{
		c := update.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		updateStorage, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Storage{
		Audience: audienceStorage,
		Message:  messageStorage,
		Metric:   metricStorage,
		MetUpd:   metupdStorage,
		TexUpd:   texupdStorage,
		Timeline: timelineStorage,
		Update:   updateStorage,
	}

	return s, nil
}
