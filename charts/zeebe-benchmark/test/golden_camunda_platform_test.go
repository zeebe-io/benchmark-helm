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

func TestGoldenCamundaPlatformDefaults(t *testing.T) {
	// Test which allows to verify also parent chart templates
	// This makes sure that properties are correctly set
	// OR configurations have been changed

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	templateNames := []string{"curator-cronjob", "curator-configmap", "service-monitor"}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: name,
			Templates:      []string{"charts/camunda-platform/templates/" + name + ".yaml"},
			SetValues:      map[string]string{"retentionPolicy.enabled": "true"},
		})
	}
}
