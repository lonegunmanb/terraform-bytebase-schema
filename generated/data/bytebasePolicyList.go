package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebasePolicyList = `{
  "block": {
    "attributes": {
      "database": {
        "description": "The database name for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment": {
        "description": "The environment resource id for the policy.",
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
      "instance": {
        "description": "The instance resource id for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_control_policy": [
                "list",
                [
                  "object",
                  {
                    "disallow_rules": [
                      "list",
                      [
                        "object",
                        {
                          "all_databases": "bool"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "backup_plan_policy": [
                "list",
                [
                  "object",
                  {
                    "retention_duration": "number",
                    "schedule": "string"
                  }
                ]
              ],
              "deployment_approval_policy": [
                "list",
                [
                  "object",
                  {
                    "default_strategy": "string",
                    "deployment_approval_strategies": [
                      "list",
                      [
                        "object",
                        {
                          "approval_group": "string",
                          "approval_strategy": "string",
                          "deployment_type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "inherit_from_parent": "bool",
              "name": "string",
              "sensitive_data_policy": [
                "list",
                [
                  "object",
                  {
                    "sensitive_data": [
                      "list",
                      [
                        "object",
                        {
                          "column": "string",
                          "mask_type": "string",
                          "schema": "string",
                          "table": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "sql_review_policy": [
                "list",
                [
                  "object",
                  {
                    "rules": [
                      "list",
                      [
                        "object",
                        {
                          "engine": "string",
                          "level": "string",
                          "payload": [
                            "list",
                            [
                              "object",
                              {
                                "format": "string",
                                "list": [
                                  "list",
                                  "string"
                                ],
                                "max_length": "number",
                                "number": "number",
                                "required": "bool"
                              }
                            ]
                          ],
                          "type": "string"
                        }
                      ]
                    ],
                    "title": "string"
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "project": {
        "description": "The project resource id for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "show_deleted": {
        "description": "Including removed policy in the response.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "The policy data source list.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebasePolicyListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebasePolicyList), &result)
	return &result
}
