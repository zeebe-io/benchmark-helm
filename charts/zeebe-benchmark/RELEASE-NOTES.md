The changelog is automatically generated using [git-chglog](https://github.com/git-chglog/git-chglog)
and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.


<a name="zeebe-benchmark-0.2.2"></a>
## [zeebe-benchmark-0.2.2](https://github.com/camunda/camunda-platform-helm/compare/zeebe-benchmark-0.2.1...zeebe-benchmark-0.2.2) (2024-08-08)

### Docs

* explain spring config

### Refactor

* load config as additional spring config
* set tag on all components
* change zeebe-gateway to zeebeGateway
* disable identityKeycloak
* rename camunda/zeebe repo

### Test

* use separate golden files for gw and zeebe
* regenerate golden files
* use right template file
* generate golden files
* add missing golden file test for zeebe deployments

### Pull Requests

* Merge pull request [#144](https://github.com/camunda/camunda-platform-helm/issues/144) from zeebe-io/release

