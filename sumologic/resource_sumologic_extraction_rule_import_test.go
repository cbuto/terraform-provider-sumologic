// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestSumologicFieldExtractionRule_import(t *testing.T) {
	var fieldextractionrule FieldExtractionRule
	testName := FieldsMap["FieldExtractionRule"]["name"]
	testScope := FieldsMap["FieldExtractionRule"]["scope"]
	testParseExpression := FieldsMap["FieldExtractionRule"]["parseExpression"]
	testEnabled, _ := strconv.ParseBool(FieldsMap["FieldExtractionRule"]["enabled"])

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFieldExtractionRuleDestroy(fieldextractionrule),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicFieldExtractionRuleConfigImported(testName, testScope, testParseExpression, testEnabled),
			},
			{
				ResourceName:      "sumologic_field_extraction_rule.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSumologicFieldExtractionRuleConfigImported(name string, scope string, parseExpression string, enabled bool) string {
	return fmt.Sprintf(`
resource "sumologic_field_extraction_rule" "foo" {
      name = "%s"
      scope = "%s"
      parse_expression = "%s"
      enabled = %t
}
`, name, scope, parseExpression, enabled)
}
