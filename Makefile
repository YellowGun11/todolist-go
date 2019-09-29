build:
	go build -o todolist

run-mongodb:
	docker run -id -v "/tmp/docker/data/mongodb:/data/db" -p "0.0.0.0:27017:27017" mongo

run-redis:
	docker run -id -p "0.0.0.0:6379:6379" -v "/tmp/docker/data/redis:/data" redis
