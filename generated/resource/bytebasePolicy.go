package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const bytebasePolicy = `{
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
      "inherit_from_parent": {
        "description": "Decide if the policy should inherit from the parent.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "instance": {
        "description": "The instance resource id for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The policy name",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description": "The project resource id for the policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The policy type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "access_control_policy": {
        "block": {
          "block_types": {
            "disallow_rules": {
              "block": {
                "attributes": {
                  "all_databases": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "backup_plan_policy": {
        "block": {
          "attributes": {
            "retention_duration": {
              "description": "The minimum allowed seconds that backup data is kept for databases in an environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schedule": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "deployment_approval_policy": {
        "block": {
          "attributes": {
            "default_strategy": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "deployment_approval_strategies": {
              "block": {
                "attributes": {
                  "approval_group": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "approval_strategy": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "deployment_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sensitive_data_policy": {
        "block": {
          "block_types": {
            "sensitive_data": {
              "block": {
                "attributes": {
                  "column": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mask_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schema": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sql_review_policy": {
        "block": {
          "attributes": {
            "title": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "rules": {
              "block": {
                "attributes": {
                  "engine": {
                    "description": "The engine for this rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "payload": {
                    "block": {
                      "attributes": {
                        "format": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "list": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "max_length": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "number": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "required": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The SQL review rules. Please check the doc for details: https://www.bytebase.com/docs/sql-review/review-policy/supported-rules",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "description": "The policy resource.",
    "description_kind": "plain"
  },
  "version": 0
}`

func BytebasePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(bytebasePolicy), &result)
	return &result
}
