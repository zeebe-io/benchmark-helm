The changelog is automatically generated using [git-chglog](https://github.com/git-chglog/git-chglog)
and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.


<a name="zeebe-benchmark-0.3.4"></a>
## [zeebe-benchmark-0.3.4](https://github.com/camunda/camunda-platform-helm/compare/zeebe-benchmark-0.3.3...zeebe-benchmark-0.3.4) (2024-10-29)

### Build

* update chart version

### Feat

* introduce worker threads
* remove repeated completion delay, and only add this if not set.

### Fix

* regen threads
* increase worker threads
* use larger disks to avoid throttling

### Test

* update golden files
* adjust golded file for new disk size

### Pull Requests

* Merge pull request [#186](https://github.com/camunda/camunda-platform-helm/issues/186) from zeebe-io/ls/increase-disk-size
* Merge pull request [#196](https://github.com/camunda/camunda-platform-helm/issues/196) from zeebe-io/release

