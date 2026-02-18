package initialize

import (
	"server/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//	func InitGorm() *gorm.DB {
//		pgsql := global.Config.Pgsql
//		db, err := gorm.Open(postgres.New(postgres.Config{
//			DSN:                  pgsql.Dsn(),
//			PreferSimpleProtocol: true, // disables implicit prepared statement usage
//		}), &gorm.Config{
//			Logger: logger.Default.LogMode(pgsql.LogLevel()),
//		})
//		if err != nil {
//			global.Log.Error("Failed to connect to pgsql:", zap.Error(err))
//		}
//		sqlDB, _ := db.DB()
//		sqlDB.SetMaxIdleConns(pgsql.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(pgsql.MaxOpenConns)
//		return db
//	}
func InitGorm() *gorm.DB {
	ms := global.Config.Mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       ms.Dsn(), // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(ms.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to pgsql:", zap.Error(err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(ms.MaxIdleConns)
	sqlDB.SetMaxOpenConns(ms.MaxOpenConns)
	return db
}
