package datasource

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/migrate"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	DatabaseTypeSQLite   = "sqlite3"
	DatabaseTypeMySQL    = "mysql"
	DatabaseTypePostgres = "postgres"
)

type DatabaseCfg struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
}

func NewClient() (*ent.Client, error) {

	var dfg = &DatabaseCfg{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.dbname"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Type:     viper.GetString("database.type"),
	}

	var client *ent.Client
	var err error
	switch dfg.Type {
	case DatabaseTypeSQLite:
		client, err = ent.Open(dfg.Type, fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1", dfg.DbName))
	case DatabaseTypeMySQL:
		client, err = ent.Open(dfg.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
			dfg.User, dfg.Password, dfg.Host, dfg.Port, dfg.DbName))
	case DatabaseTypePostgres:
		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			dfg.Host, dfg.Port, dfg.User, dfg.DbName, dfg.Password))
	default:
		return nil, fmt.Errorf("unknown database type: %s", dfg.Type)
	}
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %w", dfg.Type, err)
	}
	return client, nil
}

func AutoMigration(client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
	}
}

func DebugMode(err error, client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
	}
}
