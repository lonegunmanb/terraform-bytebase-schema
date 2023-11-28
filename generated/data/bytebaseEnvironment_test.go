package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-bytebase-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestBytebaseEnvironmentSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.BytebaseEnvironmentSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
