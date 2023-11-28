package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseEnvironmentList = `{
  "block": {
    "attributes": {
      "environments": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "environment_tier_policy": "string",
              "order": "number",
              "resource_id": "string",
              "title": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "show_deleted": {
        "description": "Including removed instance in the response.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The environment data source list.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseEnvironmentListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseEnvironmentList), &result)
	return &result
}
