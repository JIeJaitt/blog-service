package global

import "github.com/go-programming-tour-book/blog-service/pkg/setting"

var (
	// ServerSetting 是一个指向setting包中ServerSettingS结构体的指针
	// 它存储了服务器的相关配置信息，比如端口、运行模式等
	ServerSetting *setting.ServerSettingS
	// AppSetting 是一个指向setting包中AppSettingS结构体的指针
	// 它存储了应用的相关配置信息，比如日志路径、JWT密钥等
	AppSetting *setting.AppSettingS
	// DatabaseSetting 是一个指向setting包中DatabaseSettingS结构体的指针
	// 它存储了数据库的相关配置信息，比如驱动名、连接字符串等
	DatabaseSetting *setting.DatabaseSettingS
)
