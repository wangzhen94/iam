package apiserver

import (
	"github.com/wangzhen94/iam/internal/apiserver/config"
	"github.com/wangzhen94/iam/internal/apiserver/options"
	"github.com/wangzhen94/iam/pkg/app"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		//app.WithDefaultValidArgs(),
		//app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		//log.Init(opts.Log)
		//defer log.Flush()
		//
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
