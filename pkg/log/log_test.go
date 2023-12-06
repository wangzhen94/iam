package log_test

import (
	"github.com/likexian/gokit/assert"
	"github.com/spf13/pflag"
	"github.com/wangzhen94/iam/pkg/log"
	"testing"
)

func TestWithName(t *testing.T) {
	defer log.Flush()

	logger := log.WithName("logName")
	logger.Infow("hello world!", "foo", "juice")
}

func TestWithValues(t *testing.T) {
	defer log.Flush()

	logger := log.WithValues("name", "wangx")
	logger.Info("hello, world!")
	logger.Info("hello, world!")
}

func Test_V(t *testing.T) {
	defer log.Flush()

	log.V(0).Infow("hello world!", "name", "wangx")
	log.V(2).Infow("hello world!", "age", 18)
}

func Test_Option(t *testing.T) {
	fs := pflag.NewFlagSet("testName", pflag.ExitOnError)
	opt := log.NewOptions()
	opt.AddFlags(fs)

	args := []string{"--log.level=error"}
	err := fs.Parse(args)
	assert.Nil(t, err)

	assert.Equal(t, "error", opt.Level)
}
