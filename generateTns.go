package main

import (
	"errors"
	"os"
	"os/exec"
)

func GenerateTnsDeepSpeech(filepath string) (string, error) {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	cmd := exec.Command("/mnt/space/conda/c/envs/py38/bin/deepspeech", "--model", "deepspeech-0.9.3-models.pbmm", "--scorer", "deepspeech-0.9.3-models.scorer", "--audio", filepath)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out[:]), nil
}
