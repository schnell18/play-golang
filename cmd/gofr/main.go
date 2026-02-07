package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

type user struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	IsEmployed bool   `json:"isEmployed"`
}

func (u *user) TableName() string {
	return "users"
}

// // GetAll : User can overwrite the specific handlers by implementing them like this
// func (u *user) GetAll(c *gofr.Context) (any, error) {
// 	return "user GetAll called", nil
// }

func main() {
	// Create a new application
	a := gofr.New()

	// Add migrations to run
	// a.Migrate(migrations.All())
	// Run the cron job every 10 seconds(*/10)
	a.AddCronJob("*/10 * * * * *", "", func(ctx *gofr.Context) {
		ctx.Logger.Infof("current time is %v", time.Now())
	})

	// AddRESTHandlers creates CRUD handles for the given entity
	err := a.AddRESTHandlers(&user{})
	if err != nil {
		return
	}

	// Run the application
	a.Run()
}
