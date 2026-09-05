package initialize

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"dev-portfolio-api/models"
	"dev-portfolio-api/models/profile"
	"dev-portfolio-api/pkg/global"

	sqlDriver "github.com/go-sql-driver/mysql"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// InitDatabase 初始化数据库 (支持 MySQL 和 SQLite)
func InitDatabase() {
	dbType := global.Conf.System.DbType
	if dbType == "" {
		dbType = "sqlite" // 默认为 SQLite，适合轻量级部署
	}

	logLevel := gormlogger.Info
	if global.Conf.System.RunMode == "prd" {
		logLevel = gormlogger.Warn
	}

	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(logLevel),
	}

	var db *gorm.DB
	var err error

	switch dbType {
	case "mysql":
		db, err = initMysql()
		if err != nil {
			global.Log.Error(fmt.Sprintf("初始化 MySQL 失败: %v", err))
			panic(err)
		}
		// MySQL 特有的连接池配置
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100) // 轻量级实例调低一点
		sqlDB.SetConnMaxLifetime(time.Hour)
	default:
		db, err = initSqlite(gormConfig)
		if err != nil {
			global.Log.Error(fmt.Sprintf("初始化 SQLite 失败: %v", err))
			panic(err)
		}
	}

	// 设置全局 DB
	global.DB = db

	// AutoMigrate
	migrate()

	global.Log.Info(fmt.Sprintf("数据库初始化完成 [%s]", dbType))
}

// initMysql 初始化 MySQL
func initMysql() (*gorm.DB, error) {
	if global.Conf.Mysql == nil {
		return nil, fmt.Errorf("mysql 配置为空")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database,
		global.Conf.Mysql.Query,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		if e, ok := err.(*sqlDriver.MySQLError); ok && e.Number == 1049 {
			// 数据库不存在，尝试创建
			if err := createMySqlDatabase(); err != nil {
				return nil, err
			}
			// 重试连接
			return gorm.Open(mysql.Open(dsn), &gorm.Config{})
		}
		return nil, err
	}
	return db, nil
}

// createMySqlDatabase 创建 MySQL 数据库
func createMySqlDatabase() error {
	cfg := global.Conf.Mysql
	createsql := fmt.Sprintf("CREATE DATABASE `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci;", cfg.Database)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(createsql)
	if err != nil {
		return err
	}
	global.Log.Info(fmt.Sprintf("数据库 %s 创建成功", cfg.Database))
	return nil
}

// initSqlite 初始化 SQLite
func initSqlite(cfg *gorm.Config) (*gorm.DB, error) {
	file := global.Conf.Sqlite.File
	if file == "" {
		file = "data/portfolio.db"
	}

	// 确保目录存在
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("创建 SQLite 目录失败: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(file), cfg)
	if err != nil {
		return nil, err
	}

	// 开启 WAL 模式提升并发性能
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA synchronous=NORMAL;")

	return db, nil
}

// migrate 自动迁移表结构
func migrate() {
	global.Log.Info("开始自动迁移表结构...")
	err := global.DB.AutoMigrate(
		&models.User{},
		&models.ProfileInfo{},
		&models.ProfileSocial{},
		&models.ProfileNavBar{},
		&profile.ProfileSkillGroup{},
		&profile.ProfileSkill{},
		&models.Project{},
		&models.BlogPost{},
		&models.BlogAttachment{},
	)
	if err != nil {
		global.Log.Error(fmt.Sprintf("AutoMigrate 失败: %v", err))
		panic(err)
	}
	global.Log.Info("表结构迁移完成")
}
