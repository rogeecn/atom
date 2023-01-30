package config

import (
	"fmt"
)

// Database database config
type Database struct {
	Driver     string
	MySQL      MySQL
	Redis      Redis
	PostgreSQL PostgreSQL
}

// MySQL database config
type MySQL struct {
	Host     string
	Port     uint
	Database string
	Username string
	Password string
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
	User     string
	Password string
	Database string
	Host     string
	Port     uint
	SslMode  string
	TimeZone string
}

// DSN connection dsn
func (m *PostgreSQL) DSN() string {
	dsnTpl := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"
	return fmt.Sprintf(dsnTpl, m.Host, m.User, m.Password, m.Database, m.Port, m.SslMode, m.TimeZone)
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
