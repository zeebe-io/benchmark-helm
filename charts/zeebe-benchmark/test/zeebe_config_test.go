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

func TestGoldenCuratorDefaults(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	templateNames := []string{"zeebe-config"}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: name,
			Templates:      []string{"templates/" + name + ".yaml"},
			SetValues:      map[string]string{"retentionPolicy.enabled": "true"},
		})
	}
}
