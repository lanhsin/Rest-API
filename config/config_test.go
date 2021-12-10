package config

import (
	"testing"
)

const (
	testRedisAddr   = "test.redis.addr"
	testPlatAddr    = "test.plat.addr"
	testServerAddr  = "test.server.addr"
)

func TestLoad(t *testing.T) {
	configFile = "config_test.toml"
	testConfig := Config{}
	testConfig.Load()

	if testConfig.Redis.Addr != testRedisAddr {
		t.Errorf("%v != %v", testConfig.Redis.Addr, testRedisAddr)
	}
	if testConfig.Plat.Addr != testPlatAddr {
		t.Errorf("%v != %v", testConfig.Plat.Addr, testPlatAddr)
	}
	if testConfig.Plat.Addr != testPlatAddr {
		t.Errorf("%v != %v", testConfig.Plat.Addr, testPlatAddr)
	}
}
