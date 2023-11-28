package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstanceRoleList = `{
  "block": {
    "attributes": {
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
      "roles": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attribute": [
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
              ],
              "connection_limit": "number",
              "name": "string",
              "valid_until": "string"
            }
          ]
        ]
      }
    },
    "description": "The instance role data source list.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceRoleListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstanceRoleList), &result)
	return &result
}
