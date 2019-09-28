build:
	go build -o todolist

run-mongodb:
	docker run -id -p "0.0.0.0:27017:27017" -v ./docker/data/mongodb:/data/db mongo

run-redis:
	docker run -id -p "0.0.0.0:6379:6379" -v ./docker/data/redis:/data redis
