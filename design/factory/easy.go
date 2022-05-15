package factory

import "fmt"

// 简单工厂模式 场景: 生成不同的配置类

type ConfigTool interface {
	Get(val string) interface{}
}

type FileConfig struct {
}

func (this *FileConfig) Get(val string) interface{} {
	return fmt.Sprintf("file val: %s", val)
}

type YamlConfig struct {
}

func (this *YamlConfig) Get(val string) interface{} {
	return fmt.Sprintf("yaml val: %s", val)
}

// 工厂方法
func NewConfigTool(t string) ConfigTool {
	switch t {
	case "file":
		return &FileConfig{}
	case "yaml":
		return &YamlConfig{}
	}
	return nil
}
