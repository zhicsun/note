package pkg

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"note/pkg/dal/query"
	"testing"
	"time"
)

func TestOrmGen(t *testing.T) {
	ormGenerate()
}

func TestOrmFind(t *testing.T) {
	ormFind(context.Background())
}

func ormFind(ctx context.Context) {
	conn, err := getOrmDbConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	q := query.Use(conn)
	fmt.Println(q.EeoScorm.WithContext(ctx).First())
}

func getOrmDbConn() (*gorm.DB, error) {
	user := "root"
	password := ""
	hostPort := "127.0.0.1:3306"
	database := "eo_oslms_scorm"
	charset := "utf8mb4"
	format := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(format, user, password, hostPort, database, charset)
	config := mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
	}

	l := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)
	gormConfig := &gorm.Config{Logger: l}
	conn, err := gorm.Open(mysql.New(config), gormConfig)
	if err != nil {
		return conn, err
	}

	if err := conn.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return conn, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return conn, nil
}

func ormGenerate() {
	conn, err := getOrmDbConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./dal/query",
		ModelPkgPath:      "./dal/model",
		Mode:              0,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		WithUnitTest:      false,
		FieldSignable:     true,
	})

	g.UseDB(conn)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
