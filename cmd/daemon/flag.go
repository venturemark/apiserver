package daemon

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Host string
	Port string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.Host, "host", "", "127.0.0.1", "The host for binding the grpc server to.")
	cmd.Flags().StringVarP(&f.Port, "port", "", "7777", "The port for binding the grpc server to.")
}

func (f *flag) Validate() error {
	{
		if f.Host == "" {
			return tracer.Maskf(invalidFlagError, "--host must not be empty")
		}
		if f.Port == "" {
			return tracer.Maskf(invalidFlagError, "--port must not be empty")
		}
	}

	return nil
}
