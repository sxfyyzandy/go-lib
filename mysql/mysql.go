package mysql

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type MysqlCfg struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Database       string
	ConnectTimeout time.Duration //单个连接超时时间
	MaxConnectNum  int           //最大连接数
	MaxIdleNum     int           //最大连接池空闲数
	ShowSQL        bool          //是否显示sql语句
	Checktime      time.Duration //ping间隔时间
}

type MysqlPool struct {
	*xorm.Engine
}

var mysqlPool *MysqlPool

func GetMysqlPool() *MysqlPool {
	return mysqlPool
}

func InitMysqlPool(cfg *MysqlCfg) {
	p, err := newMysqlPool(cfg)
	if err != nil {
		panic(err)
	}
	mysqlPool = p
}

func newMysqlPool(cfg *MysqlCfg) (*MysqlPool, error) {
	if cfg == nil {
		return nil, errors.New("config is invalid")
	}

	driverName := "mysql"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)

	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Create DB engine for Mysql failed: %v", err.Error()))
	}
	engine.SetMaxOpenConns(cfg.MaxConnectNum)
	engine.SetMaxIdleConns(cfg.MaxIdleNum)
	if cfg.ShowSQL {
		engine.ShowSQL(true)
	} else {
		engine.SetLogger(xorm.DiscardLogger{})
	}

	p := &MysqlPool{
		Engine: engine,
	}

	return p, nil
}
