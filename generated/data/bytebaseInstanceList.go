package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstanceList = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_sources": [
                "list",
                [
                  "object",
                  {
                    "database": "string",
                    "host": "string",
                    "port": "string",
                    "title": "string",
                    "type": "string",
                    "username": "string"
                  }
                ]
              ],
              "engine": "string",
              "environment": "string",
              "external_link": "string",
              "resource_id": "string",
              "title": "string"
            }
          ]
        ]
      },
      "show_deleted": {
        "description": "Including removed instance in the response.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The instance data source list.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstanceList), &result)
	return &result
}
