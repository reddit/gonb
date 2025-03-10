# GoNB Changelog

## 0.7.7 -- 2023/08/08

* Added `DisplayMarkdown` and `UpdateMarkdown`.
* Changed `%help` to use markdown.
* `init_*` functions: 
  * Fixed duplicate rendering.
  * Added section about it in `tutorial.ipynb`
* Updated tutorial to use `%rm` to remove no longer wanted definitions.

## 0.7.6 -- 2023/07/28

* Issue #43:
  * %reset now also resets `go.mod`.
  * Added `%reset go.mod` which only resets `go.mod` but not the Go definitions memorized.

## 0.7.5 -- 2023/07/28

* Issue #30 (cont):
  * Added GONB_DIR and GONB_TMP_DIR even with the directories being used by GONB.

## 0.7.4 -- 2023/07/20

* Issue #38:
  * `%with_inputs` and `%with_password` now wait 200 milliseconds each time (a constant), before 
    prompting user with an input in the Jupyter Notebook.
  * Added `gonbui.RequestInput`, that will prompt the user with a text field in the notebook.

## 0.7.3 - 2023/07/14

* Issue #35: Fixed installation (--install): it now uses the absolute path to the gonb binary
  (as opposed to simply `os.Args[0]`).
  Also added check that it can find the "go" binary.
* Workaround for `go get` not working with `go.work`: parse `go get` errors, and if it's complaining about
  a missing package that is defined in one of the `go.work` "use" paths, it will add a suggestion the user
  add a `go.mod` replace rule.
* Added `%goworkfix` to add `use` clauses as `replace` clauses in `go.mod`.

## 0.7.2 - 2023/07/08

* Fixed bug crashing command "%cd" with no argument.
* Fixed error parsing: matching line number mix up with lines starting with 1 (instead of 0).
* Cleaned up logs: moved more logging to `klog`: most is disabled by default, but can be enabled
  for debugging passing the flags `--logtostderr --vmodule=...` (they work with `--install`).
* Fixed bug where #bytes written of parsed stderr was reported wrong, which lead to truncated errors.

## 0.7.1 - 2023/07/03

* Added support for tracking `go.work`, which allows auto-complete and contextual help
  to work with the local modules configured. It also requires `gopls` **v0.12.4** or newer to work.
* Fixed auto-complete bug when no `main` function (or no `%%`) was present in cell.
* Added special command `%cd` to chance current directory.
* Commands `%cd` and `%env` prints results of its execution.

## v0.7.0 - 2023/05/29

* Added "%ls" and "%rm" to manage memorized definitions directly.
* More contextual help and auto-complete improvements:
  * Tracking of files follows through symbolic links.

## v0.6.5 - 2023/05/23

* More contextual help and auto-complete improvements:
  * Added tracking of files in development (`%track`, `%untrack`), for usage with `gopls`.
  * Auto-track `replace` directives in `go.mod` pointing to local filesystem.

## v0.6.4 - 2023/05/22

* More InspectRequest improvements:
  * Search for identifier preceding the cursor if cursor is under a non-identifier.
  * If cursor under a ",", search for preceding function name identifier.
  * Handle case where cell is not parseable: like with auto-complete before.
* Fixed a bug where updates to `go.mod` and `go.sum` were not being notified to `gopls`.

## v0.6.3 - 2023/05/18

* Handle auto-complete case where cell is not parseable: now `gopls` is also called, and memorized
  definitions are instead saved on a second `other.go` file, for `gopls` to pick content from.
  (Issues #21 and #23).

## v0.6.2 - 2023/05/17

* Issue #23: Fixed support for generic types.

## v0.6.1

* Issue #21: Added call to `goimports` and `go get` before trying to get contextual information or auto-complete, 
  fixing many of the issues with those.

## v0.6.0

* Issue #16: Added package `cache`: implements a convenient cache of values for things that
  are expensive or slow to regenerate at each execution.
* Issue #13 and #16: Dockerfile and updated documentation.

## v0.5.1

* Fixed specialcmd_test.

## v0.5

* Improved error reporting, including indication of line number in cell.
* Parse error output of the execution of a cell, and if it contains a stack-trace, add a reference to the cell
  code (cell id and line number).
* Cleaned up, improved code documentation and testing for `goexec` package.

## v0.4.1

* Added support for Mac installation.

## v0.4.0

* "%%" or "%main" now set the program arguments as well. This may reset previously configured parameters
  given by "%args", which breaks compatibility is some cases, hence the version number bump.
* Added "UpdateHTML" and "UniqueID", to allow dynamically updated HTML content on the page.
* Fixed crash when auto-complete returns a nil structure.

## v0.3.9

* Small Go Report Card fixes (https://goreportcard.com/report/github.com/janpfeifer/gonb)

## v0.3.8

* Fixed CSS for VSCode/Github Codespaces -- it doesn't play very well with Jupyter CSS.

## v0.3.7

* Use standard Jupyter CSS color coding for error context -- should work on different themes (See #3).

## v0.3.6

* Better handling of gopls dying.
* Cleaned up and improved cursor mapping to generated Go file.
* Better handling of "didOpen" and "didChange" language server protocol with gopls.
* Monitor changes in files contents (for files being edited locally in parallel) 
  for gopls contextual help.
* Started instrumenting logs using `github.com/golang/glog`

## v0.3.5

* Display parsing errors that were disabled accidentally displaying.

## v0.3.4

* Added auto-complete.

## v0.3.3

* Fixed support of variables declared only with type but no value.
* Invoke `gopls` as a service, and talk LanguageServiceProtocol with it, to get inspection
  of symbol -- and upcoming auto-complete.
* Improved handling of cursor position: Jupyter sends UTF16 based positions (as opposed to bytes 
  or unicode runes). Still not perfect: regeneration of the Go code may get the cursor shifted.

## v0.3.2

* Added mybinder.org configuration

## v0.3.1

* Improved error message (in contextual help side-bar) if `gopls` is not installed.
* Added `--force` flag to allow installation even if `goimports` or `gopls` 
  are missing.

## v0.3.0

* Added support for **Contextual Help** (`control+I` in Jupyter), by servicing message `inpect_request`.
* Fixed keys for function receivers: when redefining a receiver as a pointer (from by value)
  they wouldn't be overwritten, and the presence of both would conflict. Special case of #1.

## v0.2.0, v0.2.1

* Added support for pointer receivers when defining methods of a type.
* Added `%env` to set environment variables from the kernel.
