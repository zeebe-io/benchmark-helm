The changelog is automatically generated using [git-chglog](https://github.com/git-chglog/git-chglog)
and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.


<a name="zeebe-benchmark-0.3.3"></a>
## [zeebe-benchmark-0.3.3](https://github.com/camunda/camunda-platform-helm/compare/zeebe-benchmark-0.3.2...zeebe-benchmark-0.3.3) (2024-09-26)

### Feat

* support credentials on starter
* support credentials on Workers
* define credentials secret
* add helper to find correct secret name
* define new values for SaaS credentials

### Fix

* scale workers down
* reduce capacity of workers
* use previously completionDelay default

### Refactor

* optimize workers

### Test

* verify existing secret behavior
* add golden test for credentials
* generate golden

### Pull Requests

* Merge pull request [#195](https://github.com/camunda/camunda-platform-helm/issues/195) from zeebe-io/ck-saas-support
* Merge pull request [#194](https://github.com/camunda/camunda-platform-helm/issues/194) from zeebe-io/ck-improve-worker-configs
* Merge pull request [#192](https://github.com/camunda/camunda-platform-helm/issues/192) from zeebe-io/release

