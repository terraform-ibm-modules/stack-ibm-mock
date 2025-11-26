package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
)

const resourceGroup = "geretain-test-resources"

func TestProjectsFullTest(t *testing.T) {
	t.Parallel()

	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing:            t,
		Prefix:             "rag",
		CatalogProductName: "Mock-Stack",
		CatalogFlavorName:  "standard",
	})

	options.StackInputs = map[string]interface{}{
		"prefix":              options.Prefix,
		"resource_group_name": resourceGroup,
	}

	err := options.RunProjectsTest()
	if assert.NoError(t, err) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}
