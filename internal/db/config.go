package db

type DatabaseConfig interface {
	GetConnString() string
	GetMigrationsPath() string
}
