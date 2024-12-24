package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Config holds the configuration values.
type Config struct {
	Port          string // The port number that the server listens on.
	Host          string // The host name that the server listens on.
	RedisURL      string // The Redis URL.
	RedisPort     string // The Redis port.
	RedisPassword string // The Redis password.
	CacheDuration string // The cache duration in minutes.
	DebugMode     bool   // Whether to run the server in debug mode.
}

var AppConfig Config

// LoadConfig initializes the AppConfig based on the following priority order:
// 1. Environment variables from the .env file (if debug mode is enabled).
// 2. Command line flags.
// 3. Default values.
// Then, it sets the Gin mode according to the debug mode configuration.
func LoadConfig() {
	AppConfig.DebugMode = strings.ToLower(os.Getenv("DEBUG_MODE")) == "true"
	loadEnvVariables()
	parseFlags()
	setConfigValues()
	setGinMode()
	logConfig()
}

// loadEnvVariables loads environment variables from a .env file if the debug mode is enabled.
// If the debug mode is disabled, no environment variables are loaded.
func loadEnvVariables() {
	if AppConfig.DebugMode {
		// Load environment variables from a .env file if it exists.
		// If the file doesn't exist, godotenv.Load() will return an error.
		// If an error occurs, a log message will be printed.
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using defaults or flags.")
		}
	}
}

// parseFlags parses the command line flags and sets the configuration values accordingly.
func parseFlags() {

	// Flag for the debug mode.
	debugModeFlag := flag.Bool("debug-mode", AppConfig.DebugMode, "Enable or disable debug mode (can also be set in .env as DEBUG_MODE).\n"+
		"Examples: -debug-mode true or --debug-mode true")

	// Flag for the port number.
	// If not provided, the default value is the one set in the .env file or the default value.
	portFlag := flag.String("port", AppConfig.Port, "Port number (can also be set in .env as PORT).\n"+
		"Examples: -port 8080 or --port 8080")

	// Flag for the host address.
	// If not provided, the default value is the one set in the .env file or the default value.
	hostFlag := flag.String("host", AppConfig.Host, "Host address (can also be set in .env as HOST).\n"+
		"Examples: -host http://example.com or --host http://example.com")

	// Flag for the Redis URL.
	// If not provided, the default value is the one set in the .env file or the default value.
	redisURLFlag := flag.String("redis-url", AppConfig.RedisURL, "Redis URL (can also be set in .env as REDIS_URL).\n"+
		"Examples: -redis-url redis://user:password@localhost:6379 or --redis-url redis://user:password@localhost:6379")

	// Flag for the Redis port.
	// If not provided, the default value is the one set in the .env file or the default value.
	redisPortFlag := flag.String("redis-port", AppConfig.RedisPort, "Redis Port (can also be set in .env as REDIS_PORT).\n"+
		"Examples: -redis-port 6379 or --redis-port 6379")

	// Flag for the Redis password.
	// If not provided, the default value is the one set in the .env file or the default value.
	redisPasswordFlag := flag.String("redis-password", AppConfig.RedisPassword, "Redis Password (can also be set in .env as REDIS_PASSWORD).\n"+
		"Examples: -redis-password password or --redis-password password")

	// Flag for the cache duration.
	// If not provided, the default value is the one set in the .env file or the default value.
	cacheDurationFlag := flag.String("cache-duration", AppConfig.CacheDuration, "Cache Duration (can also be set in .env as CACHE_DURATION).\n"+
		"Examples: -cache-duration 60 or --cache-duration 60")

	flag.Parse()

	AppConfig.DebugMode = *debugModeFlag

	setConfigValue(&AppConfig.Port, portFlag, AppConfig.Port, "8080")
	setConfigValue(&AppConfig.Host, hostFlag, AppConfig.Host, "http://127.0.0.1/")
	setConfigValue(&AppConfig.RedisURL, redisURLFlag, AppConfig.RedisURL, "")
	setConfigValue(&AppConfig.RedisPort, redisPortFlag, AppConfig.RedisPort, "6379")
	setConfigValue(&AppConfig.RedisPassword, redisPasswordFlag, AppConfig.RedisPassword, "")
	setConfigValue(&AppConfig.CacheDuration, cacheDurationFlag, AppConfig.CacheDuration, "60")
}

// setConfigValues sets the configuration values based on the parsed flags and environment variables.
// It adds the "http://" prefix to the host if it's not already there, and appends a "/" suffix if it's not already there.
func setConfigValues() {
	// Add "http://" prefix to the host if it's not already there.
	if !strings.HasPrefix(AppConfig.Host, "http://") && !strings.HasPrefix(AppConfig.Host, "https://") {
		AppConfig.Host = "httpzzzz://" + AppConfig.Host
	}
	// Removes "/" suffix fromhost if it's exists.
	AppConfig.Host = strings.TrimSuffix(AppConfig.Host, "/")

	// Add ":" prefix to the port.
	AppConfig.Port = ":" + AppConfig.Port
}

func logConfig() {
	fmt.Printf("\nConfigurations:\n\tDebug Mode: %v\n\tHost: %s\n\tPort: %s\n", AppConfig.DebugMode, AppConfig.Host, AppConfig.Port)
	fmt.Printf("\tRedis URL: %s\n\tRedis Port: %s\n\tCache Duration: %s m\n\n", AppConfig.RedisURL, AppConfig.RedisPort, AppConfig.CacheDuration)
}

// setGinMode sets the Gin mode based on the debug mode configuration.
// If the debug mode is enabled, Gin runs in debug mode.
// Otherwise, Gin runs in release mode.
func setGinMode() {
	if AppConfig.DebugMode {
		fmt.Println("Running in debug mode ...")
		// Gin is already in debug mode by default.
	} else {
		fmt.Println("Running in production mode.")
		// Gin ReleaseMode is used for production.
		// ReleaseMode panics if an error occurs.
		// See https://github.com/gin-gonic/gin#run-modes for more information.
		gin.SetMode(gin.ReleaseMode)
	}
}

// setConfigValue sets the configuration value based on the priority order:
// 1. Command line flag value.
// 2. Environment variable value.
// 3. Default value.
func setConfigValue(cfgVar *string, flag *string, envVar, defaultValue string) {
	// Check if the command line flag is provided. If yes, use it.
	if *flag != "" {
		*cfgVar = *flag
	} else if envVar != "" { // Otherwise, check if the environment variable is set. If yes, use it.
		*cfgVar = envVar
	} else { // If neither is set, fall back to the default value.
		*cfgVar = defaultValue
	}
}
