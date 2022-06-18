test:
	docker compose -f dbutil/docker-compose.yml up -d
	#go test -count=1 -v -test.short dbutil/dbutil_test.go  -run TestInitDB
	go test -count=1 -v -covermode=count -coverprofile=coverage.out ./...
	docker compose -f dbutil/docker-compose.yml stop

testall:
	#go get github.com/smartystreets/goconvey
	#goconvey

bash:
	docker compose -f dbutil/docker-compose.yml exec testdb bash

redis:
	docker-compose -f redisutil/docker-compose.yml up -d
	go test -count=1 -v -test.short redisutil/redis_test.go -run TestNewRedisClient