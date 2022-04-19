# desafio_nuveo

## Endpoints
### GET /cliente
### GET /client/{uuid}
### DELETE /client/{uuid}
### POST /cliente
```
{
	"endereco": "Rua 123",
	"nome": "test"
}
```
### PUT /client/{uuid}
```
{
	"endereco": "Nova Rua 123",
	"nome": "Novo Nome"
}
```

## Env file variables
### POSTGRES
PSQL_USER = user
<br/>
PSQL_PWD= password
<br/>
PSQL_DB= database
<br/>
PSQL_HOST= container host
<br/>

### RABBITMQ
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
