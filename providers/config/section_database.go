package config

import (
	"fmt"
)

// Database database config
type Database struct {
	Driver     string
	MySQL      *MySQL
	SQLite     *SQLite
	Redis      *Redis
	PostgreSQL *PostgreSQL
}

// MySQL database config
type MySQL struct {
	Host         string
	Port         uint
	Database     string
	Username     string
	Password     string
	Prefix       string // 表前缀
	Singular     bool   // 是否开启全局禁用复数，true表示开启
	MaxIdleConns int    // 空闲中的最大连接数
	MaxOpenConns int    // 打开到数据库的最大连接数
	Engine       string // 数据库引擎，默认InnoDB
}

func (m *MySQL) CreateDatabaseSql() string {
	return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", m.Database)
}
func (m *MySQL) EmptyDsn() string {
	dsnTpl := "%s@tcp(%s:%d)/"

	authString := func() string {
		if len(m.Password) > 0 {
			return m.Username + ":" + m.Password
		}
		return m.Username
	}

	return fmt.Sprintf(dsnTpl, authString(), m.Host, m.Port)
}

// DSN connection dsn
func (m *MySQL) DSN() string {
	dsnTpl := "%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	authString := func() string {
		if len(m.Password) > 0 {
			return m.Username + ":" + m.Password
		}
		return m.Username
	}

	return fmt.Sprintf(dsnTpl, authString(), m.Host, m.Port, m.Database)
}

type PostgreSQL struct {
	Username     string
	Password     string
	Database     string
	Host         string
	Port         uint
	SslMode      string
	TimeZone     string
	Prefix       string // 表前缀
	Singular     bool   // 是否开启全局禁用复数，true表示开启
	MaxIdleConns int    // 空闲中的最大连接数
	MaxOpenConns int    // 打开到数据库的最大连接数
}

func (m *PostgreSQL) EmptyDsn() string {
	dsnTpl := "host=%s user=%s password=%s port=%d dbname=postgres sslmode=disable TimeZone=Asia/Shanghai"

	return fmt.Sprintf(dsnTpl, m.Host, m.Username, m.Password, m.Port)
}

// DSN connection dsn
func (m *PostgreSQL) DSN() string {
	dsnTpl := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"
	return fmt.Sprintf(dsnTpl, m.Host, m.Username, m.Password, m.Database, m.Port, m.SslMode, m.TimeZone)
}

type Redis struct {
	Host     string
	Port     uint
	Database uint
	Username string
	Password string
}

// DSN connection dsn
func (m *Redis) DSN() string {
	dsnTpl := "%s:%d"
	return fmt.Sprintf(dsnTpl, m.Host, m.Port)
}

type SQLite struct {
	File string
}

func (m *SQLite) CreateDatabaseSql() string {
	return ""
}
func (m *SQLite) EmptyDsn() string {
	return m.File
}
