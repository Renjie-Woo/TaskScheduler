package tool

import (
	"encoding/json"
	"os"
	"os/exec"
)

func ReadFile(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}

func ReadStruct(filepath string, want interface{}) error {
	var bytes, err = ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &want)
	return err
}

func RunPythonScript(scriptPath string, params []string) ([]byte, error) {
	params = append([]string{scriptPath}, params...)
	cmd := exec.Command("python3", params...)
	return cmd.CombinedOutput()
}

func IsFileExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateFile(filepath string) {

}
