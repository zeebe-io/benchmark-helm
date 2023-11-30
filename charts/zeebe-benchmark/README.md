# Zeebe Benchmark Helm Chart


- [Requirements](#requirements)
- [Dependencies](#dependencies)
- [Installation](#installation)
    - [Local Kubernetes](#local-kubernetes)
- [Uninstalling Charts](#uninstalling-charts)
- [Configuration](#configuration)
    - [Global](#global)
    - [Camunda Platform](#camunda-platform)
    - [Retention Policy](#retention-policy)
    - [Worker](#worker)
    - [Starter](#publisher)
    - [Publisher](#starter)
    - [Timer](#timer)
    - [Leader Balancing](#leader-balancing)
    - [Zeebe](#zeebe)
- [Development](#development)
- [Releasing the Charts](#releasing-the-charts)


## Requirements

* [Helm](https://helm.sh/) >= 3.11.x
* Kubernetes >= 1.20+
* Minimum cluster requirements include the following to run this chart with default settings.
  All of these settings are configurable.
    * Three Kubernetes nodes to respect the default "hard" affinity settings
    * 2GB of RAM for the JVM heap

## Dependencies

Zeebe Benchmark Helm chart depends on the Camunda Platform 8 Helm chart, which has its own dependencies.

The dependency management is fully automated and managed by Helm itself;
however, it's good to understand the dependency structure. This third-party dependency is reflected in the Helm chart
as follows:

```
zeebe-benchmark
  |_camunda-platform
      |_ elasticsearch
      |_ identity
        |_ keycloak
          |_ postgresql
      |_ optimize
      |_ operate
      |_ tasklist
      |_ web-modeler
        |_ postgresql
      |_ zeebe
```

In the Zeebe Benchmark Helm chart, every sub-component from Camunda Platform 8 is disabled except Zeebe.
This can be changed via setting the right values.

## Installation

The first command adds the Zeebe Benchmark Helm charts repo, and the second installs the chart to your current 
Kubernetes context.

```shell
  helm repo add zeebe-benchmark https://zeebe-io.github.io/benchmark-helm/
  helm install this-is-a-benchmark zeebe-benchmark/zeebe-benchmark
```

### Local Kubernetes

We recommend using Helm on KIND for local environments, as the Helm configurations are battle-tested
and much closer to production systems.

For more details, follow the Camunda Platform 8
[local Kubernetes cluster guide](https://docs.camunda.io/docs/self-managed/platform-deployment/helm-kubernetes/guides/local-kubernetes-cluster/).

## Uninstalling Charts

You can remove these charts by running:

```sh
helm uninstall YOUR_RELEASE_NAME
```

> **Note**
>
> Notice that all the Services and Pods will be deleted, but not the PersistentVolumeClaims (PVC)
> which are used to hold the storage for the data generated by the cluster and Elasticsearch.

To free up the storage, you need to delete all the PVCs manually.

First, view the PVCs:

```sh
kubectl get pvc -l app.kubernetes.io/instance=YOUR_RELEASE_NAME
kubectl get pvc -l release=YOUR_RELEASE_NAME
```

Then delete the ones that you don't want to keep:

```sh
kubectl delete pvc -l app.kubernetes.io/instance=YOUR_RELEASE_NAME
kubectl delete pvc -l release=YOUR_RELEASE_NAME
```

Or you can delete the related Kubernetes namespace, which contains all PVCs.

## Configuration

The following sections contain the configuration values for the chart and each sub-chart. All of them can be overwritten
via a separate `values.yaml` file.

Check out the default [values.yaml](values.yaml) file, which contains the same content and documentation.

> **Note**
> For more details about deploying Camunda Platform 8 on Kubernetes, please visit the
> [Helm/Kubernetes installation instructions docs](https://docs.camunda.io/docs/self-managed/platform-deployment/helm-kubernetes/overview/).

### Global

| Section | Parameter | Description | Default |
|-|-|-|-|
| `global` | | Global variables which can be accessed by all sub charts | |
| | `image.repository` | Defines the repository from which to fetch the images. | "gcr.io/zeebe-io" |
| | `image.tag` | Defines the tag / version which should be used in the chart. | "SNAPSHOT" |
| | `image.pullPolicy` | Defines the [image pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) which should be used. | Always |

### Camunda Platform

Zeebe Benchmark Helm chart has a dependency on the [Camunda 8 Platform Helm chart](https://helm.camunda.io/). All variables related to Camunda Platform can be found in [camunda-platform/values.yaml](https://github.com/camunda/camunda-platform-helm/tree/main/charts/camunda-platform/values.yaml) and can be set under `camunda-platform`.

| Section | Parameter | Description | Default |
|-|-|-|-|
| `camunda-platform`| `enabled` | If true, enables Camunda Platform 8 deployment as part of the Helm chart | `true` |

**Example:**

```yaml
camunda-platform:
  enabled: true
```

Per default Zeebe and Zeebe-Gateway are enabled and all other components are disabled.

### Retention Policy

Allows to configure the elasticsearch index retention policies, which are much more fine granular in contrast to the ones provided by the parent helm chart (camunda-platform). 

| Section | Parameter | Description | Default |
|-|-|-|-|
| `retentionPolicy` | | Configuration to configure the elasticsearch index retention policies | |
| | `enabled` | If true, elasticsearch curator cronjob and configuration will be deployed. | `false` |
| | `schedule` |Defines how often/when the curator should run. | `"0 * * * *"` |
| | `policies` |Defines a list of policies, which allows to specify policies for specific indexes or globally (based on the pattern specified). | `"0 * * * *"` |
| | `image` |Configuration for the elasticsearch curator cronjob. ||


### Worker

Allows to configure the workers which can be deployed along the Zeebe Cluster. The worker code can be found [here](https://github.com/camunda/zeebe/blob/main/benchmarks/project/src/main/java/io/camunda/zeebe/Worker.java).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `worker` | | Configuration for the to be deployed worker application | |
| | `replicas` | Defines how many replicas of the application should be deployed. | `3` |
| | `capacity` | Defines how many jobs the worker should activate and work on. | `60` |


### Starter

Allows to configure the starter which can be deployed along the Zeebe Cluster. The start code can be found [here](https://github.com/camunda/zeebe/blob/main/benchmarks/project/src/main/java/io/camunda/zeebe/Starter.java).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `starter` | | Configuration for the to be deployed starter application | |
| | `replicas` | Defines how many replicas of the application should be deployed. | `1` |
| | `rate` | Defines with which rate process instances should be created by the starter. | `150` |


### Publisher

Allows to configure the publisher which can be deployed along the Zeebe Cluster. The start code can be found [here](https://github.com/camunda/zeebe/blob/main/benchmarks/project/src/main/java/io/camunda/zeebe/Starter.java).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `publisher` | | Configuration for the to be deployed publisher application | |
| | `replicas` | Defines how many replicas of the application should be deployed. | `0` |
| | `rate` | Defines with which rate message should be published. | `25` |

### Timer

Allows to configure the timer which can be deployed along the Zeebe Cluster. The start code can be found [here](https://github.com/camunda/zeebe/blob/main/benchmarks/project/src/main/java/io/camunda/zeebe/Starter.java).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `timer` | | Configuration for the to be deployed starter application | |
| | `replicas` | Defines how many replicas of the application should be deployed. | `0` |
| | `rate` | Defines with which rate process instances with timers should be created. | `25` |


### Leader Balancing

Allows to configure the auto leader balancing. For more details see [Rebalancing docs](https://docs.camunda.io/docs/self-managed/zeebe-deployment/operations/rebalancing/).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `leaderBalancing` | | Configuration for auto rebalancing feature | |
| | `enabled` | If true, enables the auto leader rebalancing. | `true` |
| | `schedule` | Defines the schedule of the auto leader rebalancing feature. | `"*/15 * * * *"` |


### Zeebe

For more information about Zeebe, visit [Zeebe Overview](https://docs.camunda.io/docs/components/zeebe/zeebe-overview/).

| Section | Parameter | Description | Default |
|-|-|-|-|
| `zeebe` | | Configuration to configure Zeebe and Gateway | |
| | `config` | Can be used to configure Zeebe Broker and Gateway additional without the need of overwriting all environment variables from the dependency chart. | `{}` |
| | `profiling` | Configuration for pyroscope profiling. | |
| | `profiling.enabled` | If true, enables the pyroscope profiling. | `false` |

## Development

For development purposes, you might want to deploy and test the charts without creating a new helm chart release.
To do this you can run the following:

```sh
 helm install YOUR_RELEASE_NAME --atomic --debug ./charts/camunda-platform
```

* `--atomic if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used`

* `--debug enable verbose output`

To generate the resources/manifests without really installing them, you can use:

* `--dry-run simulate an install`

If you see errors like:

```sh
Error: found in Chart.yaml, but missing in charts/ directory: elasticsearch
```

Then you need to download the dependencies first.

Run the following to add resolve the dependencies:

```sh
make helm.repos-add
```

After this, you can run: `make helm.dependency-update`, which will update and download the dependencies for all charts.

The execution should look like this:
```
$ make helm.dependency-update
helm dependency update charts/camunda-platform
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "camunda-platform" chart repository
...Successfully got an update from the "elastic" chart repository
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 6 charts
Dependency zeebe did not declare a repository. Assuming it exists in the charts directory
Dependency zeebe-gateway did not declare a repository. Assuming it exists in the charts directory
Dependency operate did not declare a repository. Assuming it exists in the charts directory
Dependency tasklist did not declare a repository. Assuming it exists in the charts directory
Dependency identity did not declare a repository. Assuming it exists in the charts directory
Downloading elasticsearch from repo https://helm.elastic.co
Deleting outdated charts
helm dependency update charts/camunda-platform/charts/identity
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "camunda-platform" chart repository
...Successfully got an update from the "elastic" chart repository
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
Saving 2 charts
Downloading keycloak from repo https://charts.bitnami.com/bitnami
Downloading common from repo https://charts.bitnami.com/bitnami
```

## Releasing the Charts

Please see the corresponding [release guide](../../RELEASE.md) to find out how to release the chart.