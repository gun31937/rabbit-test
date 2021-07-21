package env

import (
	"os"
	"strconv"
)

var (
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDB          int
	RedisItemTTL     int
	BlacklistURL     string
	BaseURL          string
	AppPort          string
)

func Init() {
	envInit()
}

func envInit() {

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisItemTTL, _ := strconv.Atoi(os.Getenv("REDIS_ITEM_TTL"))

	AppPort = "8080"
	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresDB = os.Getenv("POSTGRES_DB")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDB = redisDB
	RedisItemTTL = redisItemTTL
	BlacklistURL = os.Getenv("BLACKLIST_URL")
	BaseURL = os.Getenv("BASE_URL")
}
