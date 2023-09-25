package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	redisEndPoint = "localhost:6379"
	redisPassword = "123456"
	redisDB       = 0
)

func main() {
	// Set subcommand flags
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)
	setKey := setCmd.String("key", "", "Your key")
	setValue := setCmd.String("value", "", "Value of your key")

	// Json subcommand flags
	jsonCmd := flag.NewFlagSet("json", flag.ExitOnError)
	jsonKey := jsonCmd.String("key", "", "Your key")
	jsonFilename := jsonCmd.String("data", "", "Path to json data file")

	// Get subcommand flags
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getKey := getCmd.String("key", "", "Your key")

	if len(os.Args) < 2 {
		fmt.Println("error: there is no subcommand")
		os.Exit(1)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisEndPoint,
		Password: redisPassword,
		DB:       redisDB,
	})

	switch os.Args[1] {
	case "set":
		setCmd.Parse(os.Args[2:])
		if *setKey == "" {
			fmt.Println("error: there is no key")
			os.Exit(1)
		}

		ctx := context.Background()

		err := client.Set(ctx, *setKey, *setValue, 0).Err()
		if err != nil {
			panic(err)
		}

		fmt.Println("Your data has been successfully written to Redis")
	case "json":
		jsonCmd.Parse(os.Args[2:])
		if *jsonKey == "" {
			fmt.Println("error: there is no key")
			os.Exit(1)
		}

		content, err := os.ReadFile(*jsonFilename)
		if err != nil {
			fmt.Println("error: file opening failed")
		}

		ctx := context.Background()

		err = client.Set(ctx, *jsonKey, content, 0).Err()
		if err != nil {
			panic(err)
		}

		fmt.Println("Your data has been successfully written to Redis")
	case "get":
		getCmd.Parse(os.Args[2:])
		if *getKey == "" {
			fmt.Println("error: there is no key")
			os.Exit(1)
		}

		ctx := context.Background()

		value, err := client.Get(ctx, *getKey).Result()
		if err != nil {
			panic(err)
		}

		fmt.Printf("\"%s\": \"%s\"\n", *getKey, value)
	default:
		fmt.Println("error: unknow subcommand")
		os.Exit(1)
	}
}
