package database

import (
	"os"
	"testing"
)

func TestRedisDb(t *testing.T) {

	if os.Getenv("TEST_PLATFORM") == "mockup" {
		t.Logf("Skip the test (TEST_PLATFORM=%v)", os.Getenv("TEST_PLATFORM"))
		return
	}

	var redis RedisControl
	redis.Client = redis.NewClient(&redis.Options{
		Addr:   "localhost:6379",
		Password: "",
		DB:     0,
	})
	_, err := redis.Client.Ping().Result()
	if err != nil {
		t.Error("Connect to DB error!")
		t.FailNow()
	}
	defer redis.Close()

	t.Logf("Set key:test1 val:val1")
	err = redis.Set("test1", "val1")
	if err != nil {
		t.Errorf("Set DB error!: err=%v\n", err)
		t.FailNow()
	}

	val, err := redis.Get("test1")
	if err != nil {
		t.Errorf("Get DB error!: err=%v\n", err)
		t.FailNow()
	}
	if val != "val1" {
		t.Error("val != val1 error!")
		t.FailNow()
	}
	t.Logf("Get key:test1 val:%v", val)

	t.Logf("Del key:test1")
	err = redis.Del("test1")
	if err != nil {
		t.Errorf("Set DB error!: err=%v\n", err)
		t.FailNow()
	}

	val, err = redis.Get("test1")
	if err == nil {
		t.Errorf("Fail to delete key: err=%v\n", err)
		t.FailNow()
	}
	if err != nil {
		if err.Error() != "redis: nil" {
			t.Errorf("Get DB error!: err=%v\n", err)
			t.FailNow()
		}
		t.Log("The key has been deleted successfully")
	}
}
