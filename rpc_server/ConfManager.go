package rpc_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type ConfManager struct {
	jsonBody	map[string]interface{}
	projectPath	string
}

func NewConfManager() (*ConfManager, error) {
	c := &ConfManager{
		jsonBody:make(map[string]interface{}),
	}
	str, err := GetProjectPath()
	if err != nil {
		return nil, err
	}
	c.projectPath = str
	err = c.JsonConfInit()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetProjectPath() (string, error) {
	cmd := exec.Command("pwd")
	data, err := cmd.Output()
	if err != nil {
		return "", errors.New(fmt.Sprintln("command pwd fail", err))
	}
	if len(data) <= 1 {
		return "", errors.New(fmt.Sprintln("path read may error"))
	}
	path := string(data[:len(data) - 1])
	fmt.Println(path)
	return path, nil
}

func (c *ConfManager) GetConf(key string) (string, error) {
	if v, ok := c.jsonBody[key]; ok {
		return v.(string), nil
	}
	return "", errors.New(fmt.Sprintln("key = " + key + "not found"))
}

func (c *ConfManager) GetProjectPath() string {
	return c.projectPath
}

func (c *ConfManager) analysisJsonType(node map[string]interface{}) {

}

func (c *ConfManager) JsonConfInit() error {
	if c.GetProjectPath() == "" {
		return errors.New("project path empty")
	}
	strFileName := c.GetProjectPath() + "/conf/conf.json"
	strJson, err := ioutil.ReadFile(strFileName)
	if err != nil {
		return errors.New(fmt.Sprintln("read file error", err))
	}
	err = json.Unmarshal(strJson, &c.jsonBody)
	if err != nil {
		return errors.New(fmt.Sprintln("json unmarshal error ", err))
	}
	for k, v := range c.jsonBody {
		c.jsonBody[k] = v
		fmt.Println(k, " ", v)
	}
	return nil
}
