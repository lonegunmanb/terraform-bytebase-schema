package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseProject = `{
  "block": {
    "attributes": {
      "db_name_template": {
        "description": "The project db name template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key": {
        "description": "The project unique key.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "description": "The project unique resource id. Cannot change this after created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_change": {
        "description": "The project schema change type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tenant_mode": {
        "description": "The project tenant mode.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "description": "The project title.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "visibility": {
        "description": "The project visibility. Cannot change this after created",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workflow": {
        "description": "The project workflow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "databases": {
        "block": {
          "attributes": {
            "instance": {
              "description": "The instance resource id for the database.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description": "The  deployment and policy control labels.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "name": {
              "description": "The database name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schema_version": {
              "computed": true,
              "description": "The version of database schema.",
              "description_kind": "plain",
              "type": "string"
            },
            "successful_sync_time": {
              "computed": true,
              "description": "The latest synchronization time.",
              "description_kind": "plain",
              "type": "string"
            },
            "sync_state": {
              "computed": true,
              "description": "The existence of a database on latest sync.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "The databases in the project.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description": "The project resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseProject), &result)
	return &result
}
