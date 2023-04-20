package sqlite

type Config struct {
	File string
}

func (m *Config) CreateDatabaseSql() string {
	return ""
}

func (m *Config) EmptyDsn() string {
	return m.File
}
