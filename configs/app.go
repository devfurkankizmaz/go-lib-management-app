package configs

import (
	"log"

	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	app.Env = env
	if err != nil {
		log.Fatal(err.Error())
	}
	app.DB = NewDBConnection(env)
	return *app
}
