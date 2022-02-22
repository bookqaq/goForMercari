package mercarigo

import (
	"fmt"
	"os/exec"
)

const (
	EXEC_PATH = "C:\\Users\\bookq\\Desktop\\WorkSpace\\Go\\goForMercari\\mercarigo\\executor\\"
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
