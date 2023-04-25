package guesser

import (
	"embed"
	"io"
	"os"
	"path/filepath"
)

var (
	//go:embed model/*
	savedModel embed.FS
)

// as far as I know, there is no way to load a SavedModel/graph (with variables) from memory with the go tf library
// https://github.com/galeone/tfgo/issues/44#issuecomment-841806254
func writeModelToTempDir() (string, error) {
	modelPath, err := os.MkdirTemp("", "guesslang-go")
	if err != nil {
		return "", err
	}

	savedModelFile, err := savedModel.Open("model/saved_model.pb")
	if err != nil {
		return "", err
	}
	defer savedModelFile.Close()

	savedModelFileDisk, err := os.Create(modelPath + "/saved_model.pb")
	if err != nil {
		return "", err
	}
	defer savedModelFileDisk.Close()

	if _, err := io.Copy(savedModelFileDisk, savedModelFile); err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Join(modelPath, "variables"), 0755)
	if err != nil {
		return "", err
	}

	variablesFiles, err := savedModel.ReadDir("model/variables")
	if err != nil {
		return "", err
	}

	for _, file := range variablesFiles {
		variablesFile, err := savedModel.Open("model/variables/" + file.Name())
		if err != nil {
			return "", err
		}
		defer variablesFile.Close()

		variablesFileDisk, err := os.Create(filepath.Join(modelPath, "variables", file.Name()))
		if err != nil {
			return "", err
		}
		defer variablesFileDisk.Close()

		if _, err := io.Copy(variablesFileDisk, variablesFile); err != nil {
			return "", err
		}
	}

	return modelPath, nil
}
