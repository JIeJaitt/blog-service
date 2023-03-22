package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine 创建一个新的数据库引擎
// 它接受一个指向setting包中DatabaseSettingS结构体的指针作为参数
// 它根据数据库的配置信息，创建并返回一个gorm.DB类型的对象
// 它还返回一个error类型的值，表示是否有错误发生
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// 定义一个字符串s，用来存储数据库连接字符串的格式
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	// 使用gorm.Open方法打开数据库连接，并传入数据库类型和连接字符串
	// 连接字符串使用fmt.Sprintf方法根据s和databaseSetting中的字段进行格式化
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		// 如果有错误就返回nil和错误
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		// 如果服务器运行模式是debug模式，就设置db为LogMode(true)，这样可以输出更多日志信息
		db.LogMode(true)
	}
	// 设置db为SingularTable(true)，这样可以让表名使用单数形式而不是复数形式
	db.SingularTable(true)
	// 设置db底层数据库连接池的最大空闲连接数和最大打开连接数
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	// 返回db和nil（表示没有错误）
	return db, nil
}
