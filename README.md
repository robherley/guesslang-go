# guesslang-go 🔍

Go port of [yoeo/guesslang](https://github.com/yoeo/guesslang). Detects programming language of source code using a deep learning model.

## Setup

### Dependencies

Requires [`libtensorflow`](https://www.tensorflow.org/install/lang_c) C API.

On macOS, it can be installed with homebrew:

```
brew install libtensorflow
```

### Install

```
go get github.com/robherley/guesslang-go
```

See example usage in [`examples/main.go`](/example/main.go)

## Acknowledgements

Powered by:

- [yoeo/guesslang](https://github.com/yoeo/guesslang): language model
- [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow): TensorFlow
- [galeone/tfgo](https://github.com/galeone/tfgo): TensorFlow in Go

Inspired by:

- [microsoft/vscode-languagedetection](https://github.com/microsoft/vscode-languagedetection)
- [hieplpvip/guesslang-js](https://github.com/hieplpvip/guesslang-js)
