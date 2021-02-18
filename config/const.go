package config


const (
	// MySQLURI 连接mysql的uri
	MySQLURI	= "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"

	// RedisURI 连接redis的uri
	RedisURI	= "localhost:6379"

	// RpcURI	连接Dart GRPC 服务
	RpcURI		= ":9999"
	)
