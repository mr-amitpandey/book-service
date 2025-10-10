1. first install pgadmin 4 and create a blank postgressdb named 'book' under that and run the book_postgress_db_backup.sql to create the tables and required stored functions. 

swagger url: http://localhost:7000/v1/book-service/swagger/index.html

commands normally used:

cd book-service  ( make currect working directory)

swagger-cli bundle .\docs\main.yaml -o .\docs\swagger.yaml -t yaml

go build ./...

go run .\cmd\main\main.go


