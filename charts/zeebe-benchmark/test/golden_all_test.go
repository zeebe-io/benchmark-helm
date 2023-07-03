package test

import (
	"benchmark-helm/charts/zeebe-benchmark/test/golden"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenDefaults(t *testing.T) {
	// Test which allows to verify also parent chart templates
	// This makes sure that properties are correctly set
	// and if new deployments have been added OR
	// configurations have been changed

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "benchmark-test",
		Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
		GoldenFileName: "golden-all-template",
		Templates:      nil, // in order to include all templates even from camunda-platform
		SetValues:      map[string]string{"retentionPolicy.enabled": "true"},
	})
}
