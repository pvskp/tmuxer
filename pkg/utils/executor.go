package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(binary string, args ...string) (string, error) {
	var combinedOutput string = ""

	cmd := exec.Command(binary, args...)
	byteResponse, err := cmd.CombinedOutput()
	combinedOutput = string(byteResponse)
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			combinedOutput = string(byteResponse)
			return fmt.Sprintf(
				"Command %s failed with exit code %d:\n, Output: %s",
				binary,
				exitErr.ExitCode(),
				combinedOutput,
			), err
		}
	}

	if err != nil {
		return fmt.Sprintf("Command executed sucessfully but failed to combine output"), err
	}

	return combinedOutput, nil
}
