package middleware

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func PostgresMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		connStr := ""
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)
		c.Set("postgresClient", db)
	}
}
