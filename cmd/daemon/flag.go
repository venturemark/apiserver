package daemon

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	ApiServer struct {
		Host string
		Port string
	}
	Redis struct {
		Host string
		Port string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.ApiServer.Host, "apiserver-host", "", "127.0.0.1", "The host for binding the grpc apiserver to.")
	cmd.Flags().StringVarP(&f.ApiServer.Port, "apiserver-port", "", "7777", "The port for binding the grpc apiserver to.")
	cmd.Flags().StringVarP(&f.Redis.Host, "redis-host", "", "127.0.0.1", "The host for connecting with redis.")
	cmd.Flags().StringVarP(&f.Redis.Port, "redis-port", "", "6379", "The port for connecting with redis.")
}

func (f *flag) Validate() error {
	{
		if f.ApiServer.Host == "" {
			return tracer.Maskf(invalidFlagError, "--apiserver-host must not be empty")
		}
		if f.ApiServer.Port == "" {
			return tracer.Maskf(invalidFlagError, "--apiserver-port must not be empty")
		}
	}

	{
		if f.ApiServer.Host == "" {
			return tracer.Maskf(invalidFlagError, "--redis-host must not be empty")
		}
		if f.ApiServer.Port == "" {
			return tracer.Maskf(invalidFlagError, "--redis-port must not be empty")
		}
	}

	return nil
}
