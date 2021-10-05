sudo chown -R 1000:1000 ./mysql/data/

docker-compose down --rmi all --volumes --remove-orphans

docker-compose build --no-cache

docker-compose up -d

docker-compose exec app go run main.go