## Backend - Burger Fullstack

Este diretório contém a API responsável pelo gerenciamento dos pedidos da hamburgueria, incluindo montagem de hambúrgueres, registro de clientes, controle de ingredientes e status dos pedidos.

### Tecnologias Utilizadas
- **Go (Golang):** Linguagem principal do backend, escolhida pela performance e simplicidade.
- **GORM:** ORM para Go, facilita o acesso e manipulação do banco de dados relacional.
- **Docker & Docker Compose:** Orquestração de serviços e ambiente de desenvolvimento.
- **Banco de Dados:** Configurável via `config.toml` e `docker-compose.yml` (ex: PostgreSQL, MySQL).

### Estrutura de Pastas
- `config/`: Configurações da aplicação (ex: leitura de arquivos TOML).
- `db/`: Conexão, migração e queries SQL do banco de dados.
- `domain/`: Modelos de domínio (Burger, Ingredientes, Status, etc).
- `dto/`: Data Transfer Objects para comunicação entre camadas.
- `handler/`: Handlers das rotas HTTP (controllers).
- `repository/`: Repositórios para acesso ao banco de dados.
- `service/`: Regras de negócio e lógica dos serviços.

### Arquitetura
O backend segue uma arquitetura em camadas:
- **Domain:** Modelos e entidades do negócio.
- **Repository:** Abstração do acesso ao banco de dados.
- **Service:** Lógica de negócio e regras do sistema.
- **Handler:** Interface HTTP, recebe e responde requisições.

Essa separação facilita manutenção, testes e escalabilidade.

### Configuração e Execução
1. Configure o banco de dados em `config.toml` ou utilize o `docker-compose.yml` para subir os serviços.
2. Instale as dependências com:
	```sh
	go mod tidy
	```
3. Execute a aplicação:
	```sh
	go run main.go
	```
4. As rotas HTTP estarão disponíveis conforme definido nos handlers.

### Funcionalidades
- Montagem de hambúrguer personalizado (pão, carne, opcionais).
- Registro do nome do cliente.
- Gerenciamento de ingredientes.
- Controle de status do pedido: "solicitado", "em produção", "finalizado", "cancelado".
- Cancelamento de pedidos.

---
Para mais detalhes sobre endpoints e exemplos de uso, consulte a documentação das rotas nos arquivos da pasta `handler/`.
