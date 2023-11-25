package options

import (
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/component-base/pkg/json"
	"github.com/wangzhen94/iam/internal/authzserver/analytics"
	genericoptions "github.com/wangzhen94/iam/internal/pkg/options"
	"github.com/wangzhen94/iam/pkg/log"
)

type Options struct {
	RPCServer               string                                 `json:"rpcserver"      mapstructure:"rpcserver"`
	ClientCA                string                                 `json:"client-ca-file" mapstructure:"client-ca-file"`
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"         mapstructure:"server"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure"       mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"         mapstructure:"secure"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"          mapstructure:"redis"`
	Log                     *log.Options                           `json:"log"        mapstructure:"log"`
	AnalyticsOptions        *analytics.AnalyticsOptions            `json:"analytics"            mapstructure:"analytics"`
}

func NewOptions() *Options {
	return &Options{
		RPCServer:               "127.0.0.1:8081",
		ClientCA:                "",
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		RedisOptions:            genericoptions.NewRedisOptions(),
		Log:                     log.NewOptions(),
		AnalyticsOptions:        analytics.NewAnalyticsOptions(),
	}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.AnalyticsOptions.AddFlags(fss.FlagSet("analytics"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))

	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs := fss.FlagSet("misc")
	fs.StringVar(&o.RPCServer, "rpcserver", o.RPCServer, "The address of iam rpc server. "+
		"The rpc server can provide all the secrets and policies to use.")
	fs.StringVar(&o.ClientCA, "client-ca-file", o.ClientCA, ""+
		"If set, any request presenting a client certificate signed by one of "+
		"the authorities in the client-ca-file is authenticated with an identity "+
		"corresponding to the CommonName of the client certificate.")

	return fss
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Complete set default Options.
func (o *Options) Complete() error {
	return o.SecureServing.Complete()
}
