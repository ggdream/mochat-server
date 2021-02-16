package main


//go:generate docker run -d -p 6379:6379 --name mo-redis --rm redis:6.0.10
//go:generate docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=test --name mo-mysql --rm mysql:5.7.33
