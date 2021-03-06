// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsAppsyncDatasourceInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_appsync_datasource" "foo" {
	type = "AMAZON_SIMPLEDB"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsAppsyncDatasourceInvalidTypeRule(),
					Message: `"AMAZON_SIMPLEDB" is an invalid value as type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_appsync_datasource" "foo" {
	type = "AWS_LAMBDA"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsAppsyncDatasourceInvalidTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
