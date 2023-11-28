package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseEnvironment = `{
  "block": {
    "attributes": {
      "environment_tier_policy": {
        "description": "If marked as PROTECTED, developers cannot execute any query on this environment's databases using SQL Editor by default.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order": {
        "description": "The environment sorting order.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "resource_id": {
        "description": "The environment unique resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description": "The environment unique name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The environment resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseEnvironment), &result)
	return &result
}
