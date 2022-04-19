# desafio_nuveo

## BUILD
cd desafio_nuveo/
docker-compose up -d --build

## RUN TESTS
docker  exec nuveo_app go test ./... --cover