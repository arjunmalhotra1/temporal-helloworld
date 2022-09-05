# temporal-helloworld

## Bring the Temporal cluster up
git clone https://github.com/temporalio/docker-compose.git
cd docker-compose
docker-compose up

## Run the worker
go run worker/main.go

## Run the starter
go run starter/main.go
