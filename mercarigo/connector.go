package mercarigo

import (
	"fmt"
	"os/exec"
)

const (
	EXEC_PATH = "path/to/your/dPoP/generator/executable"
	FAIL_MSG  = "Run Failed."
)

func exec_func(name string, params interface{}) []byte {
	executor := exec.Command(EXEC_PATH + name)
	result, err := executor.Output()
	if err != nil {
		fmt.Println(err)
		return []byte(FAIL_MSG)
	}
	return result
}
