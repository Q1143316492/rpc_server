package OnlineJudgeService

import (
	"errors"
	"fmt"
	"rpc_server/server_base"
	"rpc_server/tools"
)

func CreateSandBoxEnv(server serverbase.IServer) error {
	fmt.Println("CreateSandBoxEnv...")
	testingMachinePath, err := server.GetConfManager().GetConf("TESTING_MACHINE_PATH")
	if err != nil {
		return err
	}
	if !tools.IsExist(testingMachinePath) && !tools.IsFile(testingMachinePath) {
		return errors.New("TESTING_MACHINE_PATH value error")
	}
	sandboxPath, err := server.GetConfManager().GetConf("SANDBOX_DIR")
	if err != nil {
		return err
	}
	if !tools.IsExist(sandboxPath) && tools.IsFile(sandboxPath) {
		return errors.New("SANDBOX_DIR value error")
	}
	imagePath, err := server.GetConfManager().GetConf("SANDBOX_IMAGE_TAR")
	if err != nil {
		return err
	}
	if !tools.IsExist(imagePath) && tools.IsFile(imagePath) {
		return errors.New("SANDBOX_IMAGE_TAR value error")
	}

	fmt.Println("success")

	return nil
}