connect:
		docker exec -it messaging_backend_postgres psql -U root

createdb:
		docker exec -it messaging_backend_postgres createdb --username=root --owner=root messaging_backend_db

migrateup:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/messaging_backend_db?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/messaging_backend_db?sslmode=disable" -verbose down

.PHONY: migrateup migratedown
