package properties_test

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"

	"kratos-sms/pkg/properties"
)

type ATest struct {
	log log.Logger
}

var tca = &ATest{
	log: nil,
}

func TestIsStructPtr(t *testing.T) {
	assert.True(t, properties.IsStructPtr(tca))

	var b int
	assert.False(t, properties.IsStructPtr(&b))
}

func TestContainsFieldPtr(t *testing.T) {
	assert.True(t, properties.ContainsFieldPtr(tca, "log"))
	assert.False(t, properties.ContainsFieldPtr(tca, "foo"))
}
