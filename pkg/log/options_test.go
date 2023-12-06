package log

import (
	"fmt"
	"github.com/likexian/gokit/assert"
	"github.com/marmotedu/log"
	"testing"
)

func TestOptions_Validate(t *testing.T) {
	opts := &log.Options{
		Level:  "test",
		Format: "test",
	}

	errors := opts.Validate()
	expect := `[unrecognized level: "test" not a valid log format: "test"]`
	assert.Equal(t, expect, fmt.Sprintf("%s", errors))
}
