Post-release tests for The Economist in Go.

INSTALLATION:

* Install web driver of choice:
* `brew install phantomjs`.
* `brew install chromedriver`.
* `brew install selenium-server-standalone`.
* Copy `.env-sample` to `.env` and configure `DRIVER` to `phantomjs`, `chromedriver` or `selenium` (the default if left empty).
* Grab dependencies: `godep restore`.
* Test away: `ginkgo` or `go test`.

This code makes use of:

* [agouti web driver](http://agouti.org/#the-agouti-dsl)
* [gingko BDD package](http://onsi.github.io/ginkgo)
* [gomega matcher package](http://onsi.github.io/gomega)