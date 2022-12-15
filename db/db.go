package db

import (
	"embed"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	UserName string
	Password string
	Host     string
	Dbname   string
	SSLMode  string
}

var (
	DB              *sqlx.DB
	embedMigrations embed.FS
)

func init() {
	DB = newDB()
}

func newDB() *sqlx.DB {
	c := initConfig()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.UserName, c.Password, c.Host, c.Port, c.Dbname, c.SSLMode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error to connect database: %v", err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "."); err != nil {
		panic(err)
	}

	return db
}

func initConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	config := &Config{
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Dbname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	return config
}
