package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"
)

// MySQL 链接结构体
type MySQLConnector struct {
	options *MysqlOptions
	tables  []interface{}
	Db      *xorm.Engine // xorm 框架实例
}

// MySQL 链接配置信息
type MysqlOptions struct {
	Hostname           string // 数据库服务器域名
	Port               string // 端口
	User               string // 数据库用户
	Password           string // 数据库密码
	DbName             string // 数据库名称
	TablePrefix        string // 数据库表前缀
	MaxOpenConnections int    // 数据库最大连接数
	MaxIdleConnections int    // 数据库最大空闲连接数
	ConnMaxLifetime    int    // 空闲链接空闲多久被回收，单位秒
}

// tables 是表格的结构体实例数组
func NewMqSQLConnector(options *MysqlOptions, tables []interface{}) MySQLConnector {
	var connector MySQLConnector
	connector.options = options
	connector.tables = tables
	// 设置数据量链接 url
	url := ""
	if options.Hostname == "" || options.Hostname == "127.0.0.1" {
		url = fmt.Sprintf(
			"%s:%s@/%s?charset=utf8&parseTime=True",
			options.User, options.Password, options.DbName)
	} else {
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			options.User, options.Password, options.Hostname, options.Port, options.DbName)
	}

	db, err := xorm.NewEngine("mysql", url) // 以 MySQL 数据库连接
	if err != nil {
		panic(fmt.Errorf("数据库初始化失败 %s", err.Error()))
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, options.TablePrefix)
	db.SetTableMapper(tbMapper)
	db.DB().SetConnMaxLifetime(time.Duration(options.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(options.MaxIdleConnections)
	db.DB().SetMaxOpenConns(options.MaxOpenConnections)
	// db.ShowSQL(true) // 是否开启打印 sql 日志到控制台
	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("数据库连接失败 %s", err.Error()))
	}

	connector.Db = db

	// 创建表，策略是不存在则创建
	if err := connector.createTables(); err != nil {
		panic(fmt.Errorf("创建数据表失败 %s", err.Error()))
	}

	return connector
}

// 创建表，策略是不存在则创建
func (s *MySQLConnector) createTables() error {
	if err := s.Db.CreateTables(s.tables...); err != nil {
		return fmt.Errorf("create mysql table error:%s", err.Error())
	}
	// 同步表的修改
	if err := s.Db.Sync2(s.tables...); err != nil {
		return fmt.Errorf("sync table error:%s", err.Error())
	}
	return nil
}
