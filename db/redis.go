package db

import (
	// "fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

type Database struct{
	Client *redis.Client
}
var ctx = context.Background()

// func ExampleClient() {
//     rdb := GetRedis().Redis.Client;
// 	// redis:=context.WithValue(ctx, "redis", rdb)

// 	err := rdb.Set(ctx, "key", "value", 0).Err()
//     if err != nil {
//         panic(err)
//     }

//     val, err := rdb.Get(ctx, "key").Result()
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("key", val)

//     val2, err := rdb.Get(ctx, "key2").Result()
//     if err == redis.Nil {
//         fmt.Println("key2 does not exist")
//     } else if err != nil {
//         panic(err)
//     } else {
//         fmt.Println("key2", val2)
//     }
//     // Output: key value
//     // key2 does not exist
// 	// return rdb;

// }

func GetRedis() *Database{
	// check := ctx.Value("redis").(*redis.Client)
	// if check != nil {
	// 	return check
	// }
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	

	// context.WithValue(ctx, "redis", rdb)
	return &Database{
		Client: client,
	 }
}