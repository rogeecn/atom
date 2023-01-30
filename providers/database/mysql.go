package mysql

import (
	"atom/container"
	"atom/providers/config"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	if err := container.Container.Provide(NewDatabase); err != nil {
		log.Fatal(err)
	}
}

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	if err := createDatabase(config.Database.MySQL.EmptyDsn(), "mysql", config.Database.MySQL.CreateDatabaseSql()); err != nil {
		return nil, err
	}

	mysqlConfig := mysql.Config{
		DSN:                       config.Database.MySQL.DSN(), // DSN data source name
		DefaultStringSize:         191,                         // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                       // 根据版本自动配置
	}

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.MySQL.Prefix,
			SingularTable: config.Database.MySQL.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	// TODO: config logger
	// _default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
	// 	SlowThreshold: 200 * time.Millisecond,
	// 	LogLevel:      logger.Warn,
	// 	Colorful:      true,
	// })
	// config.Logger = _default.LogMode(logger.Warn)

	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		return nil, err
	}

	// config instance
	db.InstanceSet("gorm:table_options", "ENGINE="+config.Database.MySQL.Engine)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.Database.MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MySQL.MaxOpenConns)

	return db, err
}

// createDatabase 创建数据库（ EnsureDB() 中调用 ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
