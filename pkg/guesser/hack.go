package guesser

import (
	"path"
	"runtime"
)

/*

TODO(robherley): This is a hack to get the path to the model directory.
I am not familiar enough with TensorFlow to properly serialize the model (yet)

*/

func hackyPathToModel() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "model")
}
