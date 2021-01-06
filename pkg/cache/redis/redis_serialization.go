package redis

//RedisSerialization 是redis序列化接口
type RedisSerialization interface {
	// Marshal 是序列化接口
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal 是反序列化接口
	Unmarshal(data []byte, v interface{}) error
}
