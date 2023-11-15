package options

import (
	//cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	//"github.com/marmotedu/component-base/pkg/json"
	//"github.com/marmotedu/component-base/pkg/util/idutil"

	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	genericoptions "github.com/wangzhen94/iam/internal/pkg/options"
	//"github.com/wangzhen94/iam/internal/pkg/server"
	"github.com/wangzhen94/iam/pkg/log"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server"   mapstructure:"server"`
	//GRPCOptions             *genericoptions.GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	//InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	//SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql"    mapstructure:"mysql"`
	RedisOptions *genericoptions.RedisOptions `json:"redis"    mapstructure:"redis"`
	//JwtOptions              *genericoptions.JwtOptions             `json:"jwt"      mapstructure:"jwt"`
	Log *log.Options `json:"log"      mapstructure:"log"`
	//FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		//GRPCOptions:             genericoptions.NewGRPCOptions(),
		//InsecureServing:         genericoptions.NewInsecureServingOptions(),
		//SecureServing:           genericoptions.NewSecureServingOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
		//JwtOptions:              genericoptions.NewJwtOptions(),
		Log: log.NewOptions(),
		//FeatureOptions:          genericoptions.NewFeatureOptions(),
	}

	return &o
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	//o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	//o.JwtOptions.AddFlags(fss.FlagSet("jwt"))
	//o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	//o.FeatureOptions.AddFlags(fss.FlagSet("features"))
	//o.InsecureServing.AddFlags(fss.FlagSet("insecure serving"))
	//o.SecureServing.AddFlags(fss.FlagSet("secure serving"))
	o.Log.AddFlags(fss.FlagSet("logs"))

	return fss
}
