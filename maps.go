package main

import "fmt"

func MapDemo() {
	config := map[string]string{
		"env": "prod",
		"db":  "neon",
	}

	fmt.Println("Config:", config)

	if env, ok := config["env"]; ok { // a map returns an key and a boolean and that's why go knows that ok is an boolean
		fmt.Println("Environment:", env)
	} else {
		fmt.Println("Environment not configured")
	}
}

