# API REST CAMPEONATO

## Preparando o ambiente
Primeiro passo é criar o arquivo _**.env**_, existe um arquio de exemplo na pasta _**configs/.env.example**_, copie o conteudo e cole no arquivo .env

O arquivo .env devem conter as seguintes variáveis:
```sh
PUBLIC_HOST=
WEB_SERVER_PORT=

DB_DRIVER=
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_DATABASE=
DB_EXTERNAL_PORT=

JWT_SECRET=
JWT_EXPIRATION_IN_SECONDS=

FOOTBALL_API_URL=
FOOTBALL_API_TOKEN=
```
> **Obs1:** As variável _**FOOTBALL_API_TOKEN**_  deve ser preenchida com o mesmo valor que está no arquivo PDF enviado.

> **Obs2:** A variável _**DB_HOST**_  deve ser preenchido com **db** pois é o nome do container do docker. 
> Caso um banco de dados local seja utilizado, mudar as outras variáveis do arquivo.


### Iniciando a aplicação com Docker
Para iniciar a aplicação, execute o comando abaixo:
```sh
  docker-compose --env-file ./configs/.env up --build`
```

> **Obs:** Caso não gere o banco de dados ao subir o container, por favor, pegue o script que se encontra na pasta **db/initdb.sql** da raiz do projeto e execute no banco de dados:

**IMPORTANTE:** Caso ele nao crie a rede interna, execute o comando abaixo:
```sh
  docker network create --subnet=192.168.200.0/24 local-network
```

### DOCUMENTAÇÃO
Para acessar a documentação da API, acesse o link abaixo:
```http
  http://localhost:8080/docs/index.html
```
A documentação foi feita utilizando o **Swagger**. O primeiro passo é logar com o usuário e senha padrão:
```
  email: user@test.com
  password: 123456
```
Após logar ele deverá gerar um token, com esse token você poderá acessar as rotas da API. Pegue o token gerado e
clique no botão **Authorize** no canto superior direito da tela, cole o token no campo **Value** e clique em **Authorize**.

Retorno esperado:
```json
{
  "code": 200,
  "expire": "2025-03-23T03:47:40Z",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI3MDE2NjAsIm9yaWdfaWF0IjoxNzQyMDk2ODYwfQ.8kBYIrCUCNF045HwDMK_nCLIEMzPb5j4peAehwAzhEY"
}
```


> Obs: Não esqueca de adicionar a palavara **Bearer** antes do token. Ex: Bearer eyJhbGciOiJIUzI1NiIsI...

**Caso não queira utilizar o swagger do projeto, fique a vontade de utilizar outra ferramenta como postman ou insomia.**

### Testes
Para rodar os testes usando **Makefile**,caso tenha, execute o comando abaixo:
```sh
  @go test -v -run . ./...
```
 ou pode simplesmente rodar o comando abaixo:
```sh
  go test -v -run . ./...
```

