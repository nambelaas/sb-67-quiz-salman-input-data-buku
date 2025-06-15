package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
)

var (
	DBConn *sql.DB
	err    error
)

func Init() {
	env := viper.GetString("App.Env")
	dbData := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString(fmt.Sprintf("Database.%s.DBHOST", env)),
		viper.GetInt(fmt.Sprintf("Database.%s.DBPORT", env)),
		viper.GetString(fmt.Sprintf("Database.%s.DBUSER", env)),
		viper.GetString(fmt.Sprintf("Database.%s.DBPASSWORD", env)),
		viper.GetString(fmt.Sprintf("Database.%s.DBNAME", env)),
	)

	DBConn, err = sql.Open("postgres", dbData)

	err = DBConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connect database")

	RunMigration()
}

func RunMigration() {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}

	n, err := migrate.Exec(DBConn, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
