package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebaseInstance = `{
  "block": {
    "attributes": {
      "engine": {
        "description": "The instance engine. Support MYSQL, POSTGRES, TIDB, SNOWFLAKE, CLICKHOUSE, MONGODB, SQLITE, REDIS, ORACLE, SPANNER, MSSQL, REDSHIFT, MARIADB, OCEANBASE.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment": {
        "description": "The environment resource id for your instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "external_link": {
        "description": "The external console URL managing this instance (e.g. AWS RDS console, your in-house DB instance console)",
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
      "resource_id": {
        "description": "The instance unique resource id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "title": {
        "description": "The instance title.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "data_sources": {
        "block": {
          "attributes": {
            "database": {
              "description": "The database for the instance, you can set this if the engine type is POSTGRES.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host": {
              "description": "Host or socket for your instance, or the account name if the instance type is Snowflake.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "The connection user password used by Bytebase to perform DDL and DML operations.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "description": "The port for your instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ssl_ca": {
              "description": "The CA certificate. Optional, you can set this if the engine type is MYSQL, POSTGRES, TIDB or CLICKHOUSE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_cert": {
              "description": "The client certificate. Optional, you can set this if the engine type is MYSQL, POSTGRES, TIDB or CLICKHOUSE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_key": {
              "description": "The client key. Optional, you can set this if the engine type is MYSQL, POSTGRES, TIDB or CLICKHOUSE.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "title": {
              "description": "The unique data source name in this instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The data source type. Should be ADMIN or RO.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
              "description": "The connection user name used by Bytebase to perform DDL and DML operations.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The connection for the instance. You can configure read-only or admin connection account here.",
          "description_kind": "plain"
        },
        "max_items": 2,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description": "The instance resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebaseInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebaseInstance), &result)
	return &result
}
