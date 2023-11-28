package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-bytebase-schema/generated/data"
	"github.com/lonegunmanb/terraform-bytebase-schema/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["bytebase_environment"] = resource.BytebaseEnvironmentSchema()  
	resources["bytebase_instance"] = resource.BytebaseInstanceSchema()  
	resources["bytebase_instance_role"] = resource.BytebaseInstanceRoleSchema()  
	resources["bytebase_policy"] = resource.BytebasePolicySchema()  
	resources["bytebase_project"] = resource.BytebaseProjectSchema()  
	dataSources["bytebase_environment"] = data.BytebaseEnvironmentSchema()  
	dataSources["bytebase_environment_list"] = data.BytebaseEnvironmentListSchema()  
	dataSources["bytebase_instance"] = data.BytebaseInstanceSchema()  
	dataSources["bytebase_instance_list"] = data.BytebaseInstanceListSchema()  
	dataSources["bytebase_instance_role"] = data.BytebaseInstanceRoleSchema()  
	dataSources["bytebase_instance_role_list"] = data.BytebaseInstanceRoleListSchema()  
	dataSources["bytebase_policy"] = data.BytebasePolicySchema()  
	dataSources["bytebase_policy_list"] = data.BytebasePolicyListSchema()  
	dataSources["bytebase_project"] = data.BytebaseProjectSchema()  
	dataSources["bytebase_project_list"] = data.BytebaseProjectListSchema()  
	Resources = resources
	DataSources = dataSources
}