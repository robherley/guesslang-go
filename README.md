# guesslang-go 🔍

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/robherley/guesslang-go)
[![Test](https://github.com/robherley/guesslang-go/actions/workflows/test.yml/badge.svg)](https://github.com/robherley/guesslang-go/actions/workflows/test.yml)

Go port of [yoeo/guesslang](https://github.com/yoeo/guesslang). Detects programming language of source code using a deep learning model.

## Setup

### Dependencies

Requires [`libtensorflow`](https://www.tensorflow.org/install/lang_c) C API.

On macOS, it can be installed with homebrew:

```
brew install libtensorflow
```

Alternatively, for Linux-based systems:

```
script/install-libtensorflow
```

### Install

```
go get github.com/robherley/guesslang-go
```

See example usage in [`examples/main.go`](/example/main.go)

## Caveats

To work around some of the limitations of the Go TensorFlow bindings (and the wrapper library)[^1], the [SavedModel](https://www.tensorflow.org/guide/saved_model) is embeded in the binary and
when a [`Guesser`](https://pkg.go.dev/github.com/robherley/guesslang-go/pkg/guesser#Guesser) is initialized, it temporarily writes the model to a directory (and removes it after).

So, in order to use this package, you must at least have a writeable temporary directory that aligns with Go's [`os.TempDir()`](https://pkg.go.dev/os@go1.20.3#TempDir).

[^1]: https://github.com/galeone/tfgo/issues/44#issuecomment-841806254

## Acknowledgements

Powered by:

- [yoeo/guesslang](https://github.com/yoeo/guesslang): language model
- [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow): TensorFlow
- [galeone/tfgo](https://github.com/galeone/tfgo): TensorFlow in Go

Inspired by:

- [microsoft/vscode-languagedetection](https://github.com/microsoft/vscode-languagedetection)
- [hieplpvip/guesslang-js](https://github.com/hieplpvip/guesslang-js)
