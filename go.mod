module sunny/test-golang

go 1.15

replace sunny/test-golang/models => ./models

replace sunny/test-golang/db => ./db

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-redis/redis/v8 v8.4.4
	github.com/joho/godotenv v1.3.0
)
