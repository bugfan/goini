package goini

import (
	"log"
	"testing"
)

func TestIni(t *testing.T) {
	// read config file
	log.Println("Start Read config file ./test.config :")
	log.Println("---------------- read config file --------------------")
	if err := LoadConfig("./test.conf"); err != nil {
		log.Fatal(err)
	}
	log.Println("num:", Config.GetInt("num"))
	log.Println("float:", Config.GetFloat("float"))
	log.Println("admin:", Config.GetString("admin"), "len:", len(Config.GetString("admin")))
	log.Println("tiger:", Config.GetString("tiger"))
	log.Println("tiger_int:", Config.GetInt("tiger"))
	log.Println("zhang:", Config.GetString("zhang"))

	log.Println("----------------read config file or system env--------------------")

	// read system env || config file

	// if args is nil it will read system env and return interface:MyEnv
	// save its to goini's Env
	// we can call it like NewEnv("./.env")
	Env = NewEnv()
	log.Println("Env TEST Value:", Env.Getenv("TEST")) //
	log.Println("Env SHELL Value:", Env.Getenv("SHELL"))
	// log.Println("All system env :", Env.GetAll())

	Env = NewEnv("./.env", "./test.conf")
	log.Println("Env TEST Value:", Env.Getenv("TEST"))   //
	log.Println("Env tiger Value:", Env.Getenv("tiger")) //
	log.Println("Env SHELL Value:", Env.Getenv("SHELL")) // can not find because it's read from file not system env
	// log.Println("All system env :", Env.GetAll())

	log.Println("Start Read config file ./test.config :")
	myEnvMap, err := ReadFile(".env") // return a map & error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("myEnvMap:", myEnvMap["ZHAO"])
}
