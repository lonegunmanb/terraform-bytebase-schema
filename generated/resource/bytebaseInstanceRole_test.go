package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-bytebase-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestBytebaseInstanceRoleSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.BytebaseInstanceRoleSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
