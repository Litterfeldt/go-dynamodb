package dynamodb

// This file overrides godynamos hardcoded necessiy for a config file
// I much prefer env variables and don't need half of all the fields of
// the configfile. Workaround.

import (
	"github.com/smugmug/godynamo/conf"
	keepalive "github.com/smugmug/godynamo/keepalive"
	"net"
	"net/url"
	"os"
)

type Config struct {
	AWSSecret    string
	AWSKeyId     string
	DynamoDBHost string
	DynamoDBZone string
}

var config Config

func GetConfig() Config {
	return config
}

func GetConfigFromEnv() (config Config) {
	config.AWSSecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	config.AWSKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	config.DynamoDBHost = os.Getenv("DYNAMODB_HOST")
	config.DynamoDBZone = os.Getenv("DYNAMODB_ZONE")
	return
}

func ConfigureDbFromConfig(c *Config) {
	configureDB(
		c.AWSKeyId,
		c.AWSSecret,
		c.DynamoDBHost,
		c.DynamoDBZone,
	)
}

func configureDB(aws_key_id, aws_secret, dynamo_host, dynamo_zone string) {
	conf.Vals.ConfLock.Lock()
	defer conf.Vals.ConfLock.Unlock()

	// make sure the dynamo endpoint is available
	addrs, addrs_err := net.LookupIP(dynamo_host)
	if addrs_err != nil {
		panic("cannot look up hostname: " + dynamo_host)
	}
	dynamo_ip := (addrs[0]).String()

	// assign the values to our globally-available conf.Vals struct instance
	conf.Vals.Auth.AccessKey = aws_key_id
	conf.Vals.Auth.Secret = aws_secret
	conf.Vals.UseSysLog = true
	conf.Vals.Network.DynamoDB.Host = dynamo_host
	conf.Vals.Network.DynamoDB.IP = dynamo_ip
	conf.Vals.Network.DynamoDB.Zone = dynamo_zone
	scheme := "http"
	port := "80"
	conf.Vals.Network.DynamoDB.Port = port
	conf.Vals.Network.DynamoDB.Scheme = scheme
	conf.Vals.Network.DynamoDB.URL = scheme + "://" + conf.Vals.Network.DynamoDB.Host +
		":" + port
	_, url_err := url.Parse(conf.Vals.Network.DynamoDB.URL)
	if url_err != nil {
		panic("confload.init: read err: conf.Vals.Network.DynamoDB.URL malformed")
	}

	// If set to true, programs that are written with godynamo may
	// opt to launch the keepalive goroutine to keep conns open.
	conf.Vals.Network.DynamoDB.KeepAlive = true

	conf.Vals.Initialized = true

	// Keep dynamodb connection alive if keepalive == true
	if conf.Vals.Network.DynamoDB.KeepAlive {
		go keepalive.KeepAlive([]string{conf.Vals.Network.DynamoDB.URL})
	}
}
