# Burger Fullstack

Sistema completo para hamburgueria, permitindo ao cliente montar hambúrgueres personalizados, registrar pedidos e acompanhar o status em tempo real. O projeto integra backend em Go e frontend em Vue.js, proporcionando uma experiência moderna tanto para o usuário quanto para o operador da loja.

## Resumo do Projeto
O Burger Fullstack é uma solução web onde o usuário pode:
- Montar seu hambúrguer escolhendo pão, carne e opcionais.
- Registrar o nome do cliente.
- Acompanhar o status do pedido: "solicitado", "em produção", "finalizado" ou "cancelado".
- Cancelar pedidos.

O sistema foi desenvolvido com foco em organização, escalabilidade e facilidade de manutenção, utilizando boas práticas de arquitetura tanto no backend quanto no frontend.

## Tecnologias Utilizadas

### Backend
- **Go (Golang):** Linguagem principal do backend.
- **GORM:** ORM para manipulação do banco de dados.
- **Docker & Docker Compose:** Orquestração de serviços e ambiente de desenvolvimento.
- **Banco de Dados:** Configurável (ex: PostgreSQL, MySQL).

### Frontend
- **Vue.js:** Framework JavaScript para construção da interface.
- **Vue Router:** Gerenciamento de rotas.
- **Babel:** Transpiler para compatibilidade de código.
- **HTML/CSS:** Estrutura e estilização das páginas.

## Estrutura do Projeto
- `backend/`: API e lógica de negócio.
- `frontend/`: Interface web e experiência do usuário.

## Como Executar
1. Siga as instruções nos READMEs das pastas `backend/` e `frontend/` para instalar dependências e executar cada parte do sistema.
2. Certifique-se de que o backend está rodando antes de iniciar o frontend.
3. Acesse a aplicação via navegador para utilizar todas as funcionalidades.

---
Para detalhes técnicos e instruções específicas, consulte os READMEs das pastas `backend/` e `frontend/`.
