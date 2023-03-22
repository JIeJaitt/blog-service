package setting

import "time"

// ServerSettingS 服务器配置结构体
type ServerSettingS struct {
	RunMode      string        // 运行模式
	HttpPort     string        // HTTP 端口
	ReadTimeout  time.Duration // 读取超时时间
	WriteTimeout time.Duration // 写入超时时间
}

type AppSettingS struct {
	DefaultPageSize int    // 默认页面大小
	MaxPageSize     int    // 最大页面大小
	LogSavePath     string // 日志保存路径
	LogFileName     string // 日志文件名
	LogFileExt      string // 日志文件扩展名
}

type DatabaseSettingS struct {
	DBType       string // 数据库类型
	UserName     string // 用户名
	Password     string // 密码
	Host         string // 主机名
	DBName       string // 数据库名
	TablePrefix  string // 数据表前缀
	Charset      string // 字符集
	ParseTime    bool   // 是否解析时间
	MaxIdleConns int    // 最大空闲连接数
	MaxOpenConns int    // 最大打开连接数
}

// ReadSection 是Setting结构体的一个方法
// 它接受一个字符串k和一个接口v作为参数
// 它从配置文件中读取k对应的部分
// 并将其解析为v类型的值
// 它返回一个error类型的值，表示是否有错误发生
func (s *Setting) ReadSection(k string, v interface{}) error {
	// 使用viper的UnmarshalKey方法将部分解析为v
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		// 如果有错误就返回错误
		return err
	}
	// 如果没有错误就返回nil
	return nil
}
