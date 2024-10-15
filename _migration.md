plugin yang di pake "golang-migrate" install dulu di pc
untuk membuat file table migration 
migrate create -ext sql -dir database/migration/ -seq init_mg

migrate create -ext sql -dir database/migrations create_users_table

running cli migration
migrate -path database/migration/ -database "mysql://root:@tcp(localhost:3306)/poskita?sslmode=disable" -verbose up

migrate -path db/migrations -database "mysql://user:password@tcp(localhost:3306)/dbname" up
