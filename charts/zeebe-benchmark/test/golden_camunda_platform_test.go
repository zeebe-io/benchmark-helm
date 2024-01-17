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
	templateNames := []string{"zeebe-service-monitor", "zeebe-gateway-service-monitor", "operate-service-monitor"}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: "c8-" + name,
			Templates:      []string{"charts/camunda-platform/templates/service-monitor/" + name + ".yaml"},
			SetValues:      map[string]string{"camunda-platform.operate.enabled": "true"},
		})
	}
}

func TestGoldenCamundaPlatformOperateDefaults(t *testing.T) {
	// Test which allows to verify also parent chart templates
	// This makes sure that properties are correctly set
	// OR configurations have been changed

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	templateNames := []string{"deployment", "service", "configmap"}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: "operate-" + name,
			Templates:      []string{"charts/camunda-platform/templates/operate/" + name + ".yaml"},
			SetValues:      map[string]string{"camunda-platform.operate.enabled": "true"},
		})
	}
}
