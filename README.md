# desafio_nuveo

# Env file variables
## POSTGRES
PSQL_USER = user
<br/>
PSQL_PWD= password
<br/>
PSQL_DB= database
<br/>
PSQL_HOST= container host
<br/>

## RABBITMQ
RABBITMQ_USER= user
<br/>
RABBITMQ_PWD= password
<br/>
RABBITMQ_HOST= rabbitmq container host
<br/>
RABBITMQ_CUSTOMER_CREATE_QUEUE= create user queue name
<br/>
NOVOS_CLIENTES= create user consumer destiny folder
<br/>

## BUILD
cd desafio_nuveo/
<br/>
cp .env_example .env
<br/>
docker-compose up -d --build
<br/>

## RUN TESTS
docker  exec nuveo_app go test ./... --cover
