package main

import (
	"log"

	mise "github.com/kkhnifes/kratos-layout/.mise/internal/mise"
)

func main() {
	cur, _ := mise.Get("env.APP_NAME")
	v, changed, err := mise.Prompt("Enter app name", cur)
	if err != nil {
		log.Fatalf("read input: %v", err)
	}
	if !changed {
		log.Printf("APP_NAME kept: %s", v)
		return
	}
	if v == "" {
		log.Fatal("app name cannot be empty")
	}
	if err := mise.Set("env.APP_NAME", v); err != nil {
		log.Fatalf("set APP_NAME: %v", err)
	}
	log.Printf("APP_NAME set to: %s", v)
}
