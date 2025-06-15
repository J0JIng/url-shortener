package database

import (
	"fmt"
	"log"
	"time"
	"learn-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
    var (
        db  *gorm.DB
        err error
        maxRetry = 5
        dsn = "host=db user=postgres password=mypassword dbname=mydb port=5432 sslmode=disable"
    )

    for i := 0; i < maxRetry; i++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            sqlDB, errPing := db.DB()
            if errPing == nil {
                errPing = sqlDB.Ping()
            }
            if errPing == nil {
                // Successful connection and ping
				db.AutoMigrate(&models.URL{})
                return db
            } else {
                err = errPing
            }
        }
        log.Printf("Attempt %d/%d: failed to connect to DB: %v", i+1, maxRetry, err)
        time.Sleep(10 * time.Second)  // Add delay between retries
    }

    panic(fmt.Sprintf("failed to connect database after %d attempts: %v", maxRetry, err))
}
