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

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	templateNames := []string{"clients-service", "leader-balancing-cron", "publisher", "starter", "timer", "worker", "benchmark-config", "workers"}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: name,
			Templates:      []string{"templates/" + name + ".yaml"},
			SetValues:      map[string]string{},
		})
	}
}

func TestGoldenWorkers(t *testing.T) {

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)
	templateNames := []string{"workers"}

	values := map[string]string{
		"workers.otherWorker.jobType":  "otherWorker-job",
		"workers.otherWorker.replicas": "1",
		"workers.otherWorker.capacity": "10",
	}

	for _, name := range templateNames {
		suite.Run(t, &golden.TemplateGoldenTest{
			ChartPath:      chartPath,
			Release:        "benchmark-test",
			Namespace:      "benchmark-" + strings.ToLower(random.UniqueId()),
			GoldenFileName: name,
			Templates:      []string{"templates/" + name + ".yaml"},
			SetValues:      values,
		})
	}
}
