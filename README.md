Post-release tests for The Economist in Go.

INSTALLATION:

* Install web driver of choice:
* `brew install phantomjs`.
* `brew install chromedriver`.
* `brew install selenium-server-standalone`.
* Copy `.env-sample` to `.env` and configure `DRIVER` to `phantomjs`, `chromedriver` or `selenium` (the default if left empty).
* Grab dependencies: `godep restore`.
* Test away: `go test`.

[agouti DSL](http://agouti.org/#the-agouti-dsl)