run:
	docker-compose up --build -d

test: 
	go test ta/pkg/users ta/web/users

goose:
	goose -dir migrate postgres `host=postgres port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable` up
	