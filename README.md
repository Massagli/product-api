# 🧩 product-api

API desenvolvida em **Go (Golang)** utilizando o framework **Gin**, **Docker** para containerização e **PostgreSQL** como banco de dados.  
O objetivo deste projeto é demonstrar uma estrutura básica de API REST com camadas bem definidas (Model, Repository, Controller e Usecase).

---

## 🚀 Tecnologias Utilizadas

- **Go lang** — Linguagem principal do projeto.
- **Gin** — Framework web para criação de APIs REST de forma rápida e eficiente.
- **Docker** — Utilizado para containerizar a aplicação e o banco de dados.
- **PostgreSQL** — Banco de dados relacional.
- **DBeaver** — Ferramenta para visualizar e manipular o banco de dados.

---

## 🛠️ Como Iniciar o Projeto

### 1. Iniciar o Projeto Go
No terminal, inicialize o módulo do Go:
```bash
go mod init [nome-do-projeto]
```

### 2. Instalar o Framework Gin
No terminal, baixe a dependência do Gin:
```bash
go get github.com/gin-gonic/gin
```
Isso cria o arquivo go.mod, responsável por gerenciar as dependências do projeto.

---

## ▶️ Rodar a Aplicação

### 1. Acesse a pasta do arquivo principal:
No terminal:
```bash
cd [nome-da-pasta]
```

### 2. Execute a aplicação:
No terminal:
```bash
go run [nome-do-arquivo].go
```

### 3. Para parar a execução:
No terminal:
```bash
Ctrl + C
```

---

## 🐳 Docker

### 1. Criar o Arquivo docker-compose.yml
Esse arquivo define os serviços (containers) necessários, como:
- API (Golang)
- Banco de Dados PostgreSQL

### 2. Subir os Containers
Após configurar o docker-compose.yml, execute:
```bash
docker compose up -d [nome-do-container]
``` 

### 3. Verificar os Containers Ativos
Após configurar o docker-compose.yml, execute:
```bash
docker container ls
``` 
---

## 🗄️ Banco de Dados

Após configurar o ambiente e o Docker, crie o banco de dados no PostgreSQL.
O DBeaver pode ser usado para visualizar e gerenciar as tabelas.

Para consumir os dados do banco dentro da API, precisamos criar models que representem a estrutura de cada tabela.

---

## 🧠 Explicação das Camadas

### 🧩 Model
Representa a estrutura dos dados do banco.
Cada tabela no banco deve ter um model equivalente na API.

### 💾 Repository
Responsável pela comunicação direta com o banco de dados.
Aqui ficam as funções de CRUD (Create, Read, Update, Delete).

### 🌐 Controller
Responsável por receber as requisições HTTP e retornar as respostas.
Ele chama os usecases para processar os dados e responde ao cliente.

### ⚙️ Usecase
Camada intermediária entre o Controller e o Repository.
É onde fica a regra de negócio da aplicação.

---

## 🔄 Relação entre as Camadas

                 ┌────────────────────────────┐
                 │          CLIENTE           │
                 │ (Postman, Frontend, etc.)  │
                 └──────────────┬─────────────┘
                                │ HTTP Request
                                ▼
                   ┌────────────────────────┐
                   │       CONTROLLER       │
                   │ Recebe e envia dados   │
                   └────────────┬───────────┘
                                │ Chama regra de negócio
                                ▼
                    ┌──────────────────────┐
                    │       USECASE        │
                    │ Regras de negócio    │
                    └──────────┬───────────┘
                               │ Solicita dados
                               ▼
                   ┌────────────────────────┐
                   │      REPOSITORY        │
                   │ Acesso ao banco        │
                   └────────────┬───────────┘
                                │
                                ▼
                    ┌──────────────────────┐
                    │       DATABASE       │
                    │ (PostgreSQL)         │
                    └──────────────────────┘
                    
                               ▲
                               │
                        ┌─────────────┐
                        │    MODEL    │
                        │ Estrutura   │
                        │ dos Dados   │
                        └─────────────┘


### 🍽️ Analogia:
Pensa no seu sistema como um restaurante:

- 🧺 Banco de dados → é o estoque do restaurante, onde ficam guardados todos os ingredientes (as informações do sistema).

- 🧑‍🍳 Repository → é o ajudante de cozinha. Ele vai até o estoque (banco de dados), busca os ingredientes (dados) e leva para o chef.

- 👨‍🍳 Usecase → é o chef de cozinha. Ele decide como o prato vai ser preparado, aplica as receitas e garante que tudo siga as regras do restaurante (as regras de negócio).

- 🧾 Controller → é o garçom que fala com o cliente. Ele recebe o pedido (requisição HTTP), leva até o chef (usecase) e depois entrega o prato pronto (resposta).

### Resumo:
-  Em resumo então o repository pega os dados do banco, passa para o usecase para os dados serem tratados e joga para o controller para exibir após receber a requisição.

### 💡 Resumo da Comunicação:
- O Cliente envia uma requisição → Controller.
- O Controller chama o Usecase, que aplica as regras de negócio.
- O Usecase usa o Repository para buscar ou salvar dados.
- O Repository conversa com o Banco de Dados.
- O Model representa a estrutura que trafega entre todas as camadas.

---

## 📚 Observação
- Esse padrão é inspirado na Clean Architecture, garantindo:
- Separação clara de responsabilidades;
- Facilidade de manutenção;
- Testes mais simples;
- Escalabilidade a longo prazo.


## 🐳 Docker

### 1. Criar o Arquivo DockerFile
....

### 2. Gerar Imagem da API
Após configurar o Dockerfile, execute:
```bash
docker build - t [nome-da-imagem] .
``` 

### 3. Verificar as Imagens Ativas
Após configurar o docker-compose.yml, execute:
```bash
docker container ls
``` 

### 4. Verificar as Imagens Ativas
No docker compose precisamos criar o service tanto do banco quanto da API

### 5. Subir os Containers
Após configurar o docker-compose.yml, execute:
```bash
docker compose up -d 
``` 

---


