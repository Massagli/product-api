Aplicativos/Ferramentas
- Go lang
- Docker
- PostgreSQL
- Dbeaver
- 


# product-api

# Iniciar Projeto Go
-> no terminal: go mod init [nome do projeto]

# Download Framework (gin)
-> no terminal: go get [path do pacote]
ele cria a pasta "go.mod" que contem as dependencias diretas e indiretas.

# Rodar a aplicação
-> para rodar precisamos entrar na pasta onde se encontra o nosso arquivo principal.
-> no terminal: cd [nome da pasta]
-> no terminal: go run [nome do arquivo .go]

# Derrumar a aplicação
-> no terminal: ctrl + c

# Criar docker
-> criar arquivo "docker-compose.yml"

# Iniciar docker
-> após configurar
-> no terminal: docker compose up -d [nome do container]

# Testar container
-> no terminal: docker container ls


Após configurar o ambiente devemos criar o banco de dados.
Para consumir os dados registrados nas tabelas, precisamos criar um model na API pra comportar a estrutura feita no SQL

# Criar pasta "model"
-> para cada tabela precisamos criar um modelo. 