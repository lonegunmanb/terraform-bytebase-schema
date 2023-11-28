package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstanceRole = `{
  "block": {
    "attributes": {
      "connection_limit": {
        "description": "Connection count limit for role",
        "description_kind": "plain",
        "optional": true,
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
      "password": {
        "computed": true,
        "description": "The password.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "valid_until": {
        "description": "It sets a date and time after which the role's password is no longer valid. Should be a timestamp in \"2006-01-02T15:04:05+08:00\" format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "attribute": {
        "block": {
          "attributes": {
            "bypass_rls": {
              "description": "Set the ` + "`" + `BYPASSRLS` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "can_login": {
              "description": "Set the ` + "`" + `LOGIN` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_db": {
              "description": "Set the ` + "`" + `CREATEDB` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_role": {
              "description": "Set the ` + "`" + `CREATEROLE` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "no_inherit": {
              "description": "Set the ` + "`" + `NOINHERIT` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "replication": {
              "description": "Set the ` + "`" + `REPLICATION` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "super_user": {
              "description": "Set the ` + "`" + `SUPERUSER` + "`" + ` attribute for the role. Default ` + "`" + `false` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "The attribute for the role.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "The role resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstanceRole), &result)
	return &result
}
