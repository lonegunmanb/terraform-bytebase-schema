package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseProject = `{
  "block": {
    "attributes": {
      "databases": {
        "computed": true,
        "description": "The databases in the project.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instance": "string",
              "labels": [
                "map",
                "string"
              ],
              "name": "string",
              "schema_version": "string",
              "successful_sync_time": "string",
              "sync_state": "string"
            }
          ]
        ]
      },
      "db_name_template": {
        "computed": true,
        "description": "The project db name template.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "The project key.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "description": "The project unique resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_change": {
        "computed": true,
        "description": "The project schema change type.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_version": {
        "computed": true,
        "description": "The project schema version.",
        "description_kind": "plain",
        "type": "string"
      },
      "tenant_mode": {
        "computed": true,
        "description": "The project tenant mode.",
        "description_kind": "plain",
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The project title.",
        "description_kind": "plain",
        "type": "string"
      },
      "visibility": {
        "computed": true,
        "description": "The project visibility.",
        "description_kind": "plain",
        "type": "string"
      },
      "workflow": {
        "computed": true,
        "description": "The project workflow.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The project data source.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseProject), &result)
	return &result
}
