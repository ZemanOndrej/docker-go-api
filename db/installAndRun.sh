docker build -t go-postgre-db .
docker run --rm -p 5432:5432 go-postgre-db