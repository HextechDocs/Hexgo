/*
Package configuration implements the configuration logic for the application.

It has 2 provider options: File and Consul

File simply loads the config.json from next to the binary. If it doesn't exist, it'll create one based on the example struct.
Consul loads everything from the dynamic configuration service. If it doesn't exist, it'll attempt to save the default config.

The contents of the configuration is not validated so you're expected to fill out everything you require.
*/
package configuration

type Configuration struct {
	DatabaseHost     string `json:"DatabaseHost"`
	DatabasePort     int    `json:"DatabasePort"`
	DatabaseUser     string `json:"DatabaseUser"`
	DatabasePassword string `json:"DatabasePassword"`
	DatabaseName     string `json:"DatabaseName"`

	ShouldUseRedis         bool   `json:"ShouldUseRedis"`
	RedisUrl               string `json:"RedisUrl"`
	ShouldUseRedisAuth     bool   `json:"ShouldUseRedisAuth"`
	RedisPassword          string `json:"RedisPassword"`
	ShouldUseRedisDatabase bool   `json:"ShouldUseRedisDatabase"`
	RedisDatabase          string `json:"RedisDatabase"`

	ShouldLogToElastic   bool     `json:"ShouldLogToElastic"`
	ElasticHosts         []string `json:"ElasticHosts"`
	ShouldUseElasticAuth bool     `json:"ShouldUseElasticAuth"`
	ElasticUsername      string   `json:"ElasticUsername"`
	ElasticPassword      string   `json:"ElasticPassword"`
	ElasticIndex         string   `json:"ElasticIndex"`
}

// exampleConfig returns a configuration with everything filled in with placeholder values for later customization.
func exampleConfig() *Configuration {
	return &Configuration{
		DatabaseHost:           "dbhost.local",
		DatabasePort:           5432,
		DatabaseUser:           "root",
		DatabasePassword:       "changeme",
		DatabaseName:           "hextechdocs",

		ShouldUseRedis:         true,
		RedisUrl:               "redishost.local:6379",
		ShouldUseRedisAuth:     true,
		RedisPassword:          "changeme",
		ShouldUseRedisDatabase: true,
		RedisDatabase:          "15",

		ShouldLogToElastic:		false,
		ElasticHosts: 			[]string{"elastic.local:9300"},
		ShouldUseElasticAuth: 	false,
		ElasticUsername: 		"admin",
		ElasticPassword: 		"supersecret",
		ElasticIndex:			"hexgo",
	}
}
