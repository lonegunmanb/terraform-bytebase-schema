package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseEnvironment = `{
  "block": {
    "attributes": {
      "environment_tier_policy": {
        "computed": true,
        "description": "If marked as PROTECTED, developers cannot execute any query on this environment's databases using SQL Editor by default.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "order": {
        "computed": true,
        "description": "The environment sorting order.",
        "description_kind": "plain",
        "type": "number"
      },
      "resource_id": {
        "description": "The environment unique resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The environment unique name.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The environment data source.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseEnvironment), &result)
	return &result
}
