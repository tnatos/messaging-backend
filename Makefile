.PHONY: create-db migrate-force migrate-up migrate-down

current_directory = $(shell pwd)
migration_directory = $(current_directory)/db/migration
port = 5432
n = 1
database_name = messaging_backend_db
container_name = messaging_backend_postgres
username = postgres
password = password

create-db:
	docker exec -it messaging_backend_postgres createdb --username=$(username) --owner=$(username) messaging_backend_db

# Run migrate-forcd, when dirty read error occurs
migrate-force:
	migrate -path $(migration_directory) -database postgresql://$(username):$(password)@localhost:$(port)/$(database_name)?sslmode=disable force $(n)

migrate-up:
	migrate -path $(migration_directory) -database postgresql://$(username):$(password)@localhost:$(port)/$(database_name)?sslmode=disable -verbose up $(n)

migrate-down:
	migrate -path $(migration_directory) -database postgresql://$(username):$(password)@localhost:$(port)/$(database_name)?sslmode=disable -verbose down $(n)
