package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseProjectList = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "projects": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "databases": [
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
              ],
              "db_name_template": "string",
              "key": "string",
              "resource_id": "string",
              "schema_change": "string",
              "schema_version": "string",
              "tenant_mode": "string",
              "title": "string",
              "visibility": "string",
              "workflow": "string"
            }
          ]
        ]
      },
      "show_deleted": {
        "description": "Including removed project in the response.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The project data source list.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseProjectListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseProjectList), &result)
	return &result
}
