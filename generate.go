package main


//go:generate docker run -d -p 6379:6379 --name mo-redis --rm redis:6.0.10
//go:generate docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=test --name mo-mysql --rm mysql:5.7.33
//go:generate docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 --name mo-mongo --rm mongo:4.4.4
