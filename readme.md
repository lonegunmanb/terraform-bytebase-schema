# Terraform Bytebase Provider Schema Repository

This repository contains the generated Go files for the Bytebase provider schemas, which are based on the Terraform AWS Provider. These schema files can be used as a reference when writing tools, such as TFLint plugins, that interact with the AWS provider.

## Repository Structure

Each tag version of the Terraform Bytebase Provider has a corresponding tag in this repository. You can find the schema files for each provider version under the respective tag.

e.g.: to use `bytebase`'s `0.0.9` schema, you could:

```shell
$ go get github.com/lonegunmanb/terraform-bytebase-schema@v0.0.9
```

Then you can read schemas like this:

```go
import (
"testing"

"github.com/lonegunmanb/terraform-bytebase-schema/generated"
"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
assert.NotEmpty(t, generated.Resources)
assert.NotEmpty(t, generated.DataSources)
}
```

## Generating Schema Files

The schema files are generated using the terraform provider schema -json command. This command retrieves the schema information and converts it into JSON format. The JSON files are then converted into Go files, which can be found in this repository.

If you encounter any issues or would like to contribute to this repository, please submit an issue or a pull request on GitHub.

## License

[MIT](LICENSE)
