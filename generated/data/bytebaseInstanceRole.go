package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstanceRole = `{
  "block": {
    "attributes": {
      "attribute": {
        "computed": true,
        "description": "The attribute for the role.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bypass_rls": "bool",
              "can_login": "bool",
              "create_db": "bool",
              "create_role": "bool",
              "no_inherit": "bool",
              "replication": "bool",
              "super_user": "bool"
            }
          ]
        ]
      },
      "connection_limit": {
        "computed": true,
        "description": "Connection count limit for role",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The instance resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The role unique name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "valid_until": {
        "computed": true,
        "description": "It sets a date and time after which the role's password is no longer valid.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The instance role data source.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstanceRole), &result)
	return &result
}
