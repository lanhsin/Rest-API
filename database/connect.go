package database

type Connector interface {
	HGet(key string, field string) (string, error)
	HGetInt64(key string, field string) (int64, error)
	HGetAll(key string) (map[string]string, error)
	HSet(key string, field string, data interface{}) (error)
	Get(key string) (val string, err error)
	Set(key string, val string) error
	Del(key string) error
	GetInt32(key string) (val int32, err error)
	GetInt64(key string) (val int64, err error)
	GetFloat32(key string) (val float32, err error)
	GetFloat64(key string) (val float64, err error)
}
