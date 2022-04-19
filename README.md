# desafio_nuveo

# Env file variables
## POSTGRES
PSQL_USER = user
PSQL_PWD= password
PSQL_DB= database
PSQL_HOST= container host

## RABBITMQ
RABBITMQ_USER= user
RABBITMQ_PWD= password
RABBITMQ_HOST= rabbitmq container host
RABBITMQ_CUSTOMER_CREATE_QUEUE= create user queue name
NOVOS_CLIENTES= create user consumer destiny folder

## BUILD
cd desafio_nuveo/
cp .env_example .env
docker-compose up -d --build

## RUN TESTS
docker  exec nuveo_app go test ./... --cover