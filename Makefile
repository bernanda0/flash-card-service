USERNAME = bernanda
DB_NAME = sr-db
IMAGE_NAME = my-postgres
CONTAINER_NAME = spaced-repetition-db

# make custom image of pg based on dockerfile
pg-build:
	docker build -t $(IMAGE_NAME) .
# make a container and run it
pg-run:
	docker run -d --name $(CONTAINER_NAME) -p 5432:5432 $(IMAGE_NAME)
pg-createuser:
	docker exec -it $(CONTAINER_NAME) createuser -U postgres $(USERNAME)
pg-dropuser:
	# make pg-dropdb && \
	docker exec -it $(CONTAINER_NAME) dropuser -U postgres $(USERNAME)
pg-createdb:
	docker exec -it $(CONTAINER_NAME) createdb -U postgres -O $(USERNAME) $(DB_NAME)
pg-dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb -U postgres $(DB_NAME)
pg-psql:
	docker exec -it $(CONTAINER_NAME) psql -U $(USERNAME) -d $(DB_NAME) 

migrate-init:
	migrate create -ext sql -dir db/migration -seq db_scheme

# add the password first inside psql, e.g. ALTER USER yourusername WITH PASSWORD yourpassword
migrate-up:
	migrate -path db/migration -database "postgresql://$(USERNAME):bernanda@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://$(USERNAME):bernanda@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

