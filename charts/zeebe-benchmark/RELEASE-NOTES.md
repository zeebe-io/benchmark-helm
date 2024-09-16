The changelog is automatically generated using [git-chglog](https://github.com/git-chglog/git-chglog)
and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.


<a name="zeebe-benchmark-0.3.0"></a>
## [zeebe-benchmark-0.3.0](https://github.com/camunda/camunda-platform-helm/compare/zeebe-benchmark-0.2.4...zeebe-benchmark-0.3.0) (2024-09-16)

### Docs

* cadd better example

### Feat

* add example values file
* add publish message configuration
* configure starter app
* add payloadPath and completionDelay as config
* extend worker configuration

### Fix

* correctly reset benchmark app
* configure starter correctly
* set extraBpmnModels correct in starter app
* scrape correct endpoint
* separate deployments

### Refactor

* remove worker
* introduce workers values
* replace defaults

### Test

* run workers test separate
* add test for workers
* generate-golden files
* remove capability as we run in priviliged mode
* update golden files

### Reverts

* fix: scrape correct endpoint 

### Pull Requests

* Merge pull request [#180](https://github.com/camunda/camunda-platform-helm/issues/180) from zeebe-io/release

