package setting

import "github.com/spf13/viper"

// Setting 结构体用于存储应用程序配置信息
type Setting struct {
	vp *viper.Viper
}

// NewSetting 函数用于创建并初始化Setting实例
func NewSetting() (*Setting, error) {
	// 创建一个viper实例，用于读取配置文件
	vp := viper.New()
	// 设置要读取的配置文件名为"config"
	vp.SetConfigName("config")
	// 设置要读取的配置文件路径为"configs/"
	vp.AddConfigPath("configs/")
	// 设置要读取的配置文件类型为yaml格式
	vp.SetConfigType("yaml")
	// 读取配置文件，并将配置信息存储在viper实例中
	err := vp.ReadInConfig()
	if err != nil {
		// 如果读取配置文件失败，则返回nil和错误信息
		return nil, err
	}

	// 返回一个包含viper实例的Setting实例
	return &Setting{vp}, nil
}
