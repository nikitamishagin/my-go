package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)
	redisEndPoint := setCmd.String("endpoint", "localhost:6379", "Redis server address")
	redisPassword := setCmd.String("password", "", "Redis password")
	redisDB := setCmd.Int("db", 0, "Redis data base number")
	key := setCmd.String("key", "", "Your key")
	value := setCmd.String("value", "", "Value of your key")

	if len(os.Args) < 2 {
		fmt.Println("error: there is no subcommand")
		os.Exit(1)
	}

	setCmd.Parse(os.Args[2:])

	if *key == "" {
		fmt.Println("error: there is no key")
		os.Exit(1)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     *redisEndPoint,
		Password: *redisPassword,
		DB:       *redisDB,
	})

	switch os.Args[1] {
	case "set":
		ctx := context.Background()

		err := client.Set(ctx, *key, *value, 0).Err()
		if err != nil {
			panic(err)
		}

		fmt.Println("Your data has been successfully written to Redis")
	case "get":
		ctx := context.Background()

		result, err := client.Get(ctx, *key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value: %s\n", result)
	default:
		fmt.Println("error: unknow subcommand")
		os.Exit(1)
	}
}
