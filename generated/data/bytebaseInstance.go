package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstance = `{
  "block": {
    "attributes": {
      "data_sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
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
        ]
      },
      "engine": {
        "computed": true,
        "description": "The instance engine. Support MYSQL, POSTGRES, TIDB, SNOWFLAKE, CLICKHOUSE, MONGODB, SQLITE, REDIS, ORACLE, SPANNER, MSSQL, REDSHIFT, MARIADB, OCEANBASE.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "The environment name for your instance in \"environments/{resource id}\" format.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_link": {
        "computed": true,
        "description": "The external console URL managing this instance (e.g. AWS RDS console, your in-house DB instance console)",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "description": "The instance unique resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The instance title.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The instance data source.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstance), &result)
	return &result
}
