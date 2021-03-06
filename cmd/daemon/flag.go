package daemon

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	ApiServer struct {
		Host                   string
		Port                   string
		TerminationGracePeriod time.Duration
	}
	Metrics struct {
		Host string
		Port string
	}
	Redis struct {
		Host string
		Kind string
		Port string
	}
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.ApiServer.Host, "apiserver-host", "", "127.0.0.1", "The host for binding the grpc apiserver endpoints to.")
	cmd.Flags().StringVarP(&f.ApiServer.Port, "apiserver-port", "", "7777", "The port for binding the grpc apiserver endpoints to.")
	cmd.Flags().DurationVarP(&f.ApiServer.TerminationGracePeriod, "apiserver-termination-grace-period", "", 5*time.Second, "The time to wait before terminating the apiserver process.")

	cmd.Flags().StringVarP(&f.Metrics.Host, "metrics-host", "", "127.0.0.1", "The host for binding the http metrics endpoints to.")
	cmd.Flags().StringVarP(&f.Metrics.Port, "metrics-port", "", "8000", "The port for binding the http metrics endpoints to.")

	cmd.Flags().StringVarP(&f.Redis.Host, "redis-host", "", "127.0.0.1", "The host for connecting with redis.")
	cmd.Flags().StringVarP(&f.Redis.Kind, "redis-kind", "", "single", "The kind of redis to connect to, e.g. simple or sentinel.")
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
		if f.ApiServer.TerminationGracePeriod == 0 {
			return tracer.Maskf(invalidFlagError, "--apiserver-termination-grace-period must not be empty")
		}
	}

	{
		if f.Metrics.Host == "" {
			return tracer.Maskf(invalidFlagError, "--metrics-host must not be empty")
		}
		if f.Metrics.Port == "" {
			return tracer.Maskf(invalidFlagError, "--metrics-port must not be empty")
		}
	}

	{
		if f.Redis.Host == "" {
			return tracer.Maskf(invalidFlagError, "--redis-host must not be empty")
		}
		if f.Redis.Kind == "" {
			return tracer.Maskf(invalidFlagError, "--redis-kind must not be empty")
		}
		if f.Redis.Port == "" {
			return tracer.Maskf(invalidFlagError, "--redis-port must not be empty")
		}
	}

	return nil
}
