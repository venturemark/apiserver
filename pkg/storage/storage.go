package storage

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/audience"
	"github.com/venturemark/apiserver/pkg/storage/message"
	"github.com/venturemark/apiserver/pkg/storage/role"
	"github.com/venturemark/apiserver/pkg/storage/texupd"
	"github.com/venturemark/apiserver/pkg/storage/timeline"
	"github.com/venturemark/apiserver/pkg/storage/update"
	"github.com/venturemark/apiserver/pkg/storage/user"
	"github.com/venturemark/apiserver/pkg/storage/venture"
)

type Config struct {
	Logger     logger.Interface
	Permission permission.Gateway
	Redigo     redigo.Interface
	Rescue     rescue.Interface
}

type Storage struct {
	Audience *audience.Audience
	Message  *message.Message
	Role     *role.Role
	TexUpd   *texupd.TexUpd
	Timeline *timeline.Timeline
	Update   *update.Update
	User     *user.User
	Venture  *venture.Venture
}

func New(config Config) (*Storage, error) {
	var err error

	var audienceStorage *audience.Audience
	{
		c := audience.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		audienceStorage, err = audience.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var messageStorage *message.Message
	{
		c := message.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		messageStorage, err = message.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var roleStorage *role.Role
	{
		c := role.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		roleStorage, err = role.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var texupdStorage *texupd.TexUpd
	{
		c := texupd.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		texupdStorage, err = texupd.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var timelineStorage *timeline.Timeline
	{
		c := timeline.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		timelineStorage, err = timeline.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var updateStorage *update.Update
	{
		c := update.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		updateStorage, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var userStorage *user.User
	{
		c := user.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		userStorage, err = user.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var ventureStorage *venture.Venture
	{
		c := venture.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		ventureStorage, err = venture.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Storage{
		Audience: audienceStorage,
		Message:  messageStorage,
		Role:     roleStorage,
		TexUpd:   texupdStorage,
		Timeline: timelineStorage,
		Update:   updateStorage,
		User:     userStorage,
		Venture:  ventureStorage,
	}

	return s, nil
}
