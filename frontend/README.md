# Frontend - Burger Fullstack

Este diretório contém a interface web do sistema de pedidos da hamburgueria, permitindo ao usuário montar hambúrgueres, registrar pedidos e acompanhar o status em tempo real.

### Tecnologias Utilizadas
- **Vue.js:** Framework JavaScript para construção de interfaces reativas e componentes reutilizáveis.
- **Vue Router:** Gerenciamento de rotas e navegação entre páginas.
- **Babel:** Transpiler para compatibilidade de código JavaScript moderno.
- **HTML/CSS:** Estrutura e estilização das páginas.

### Estrutura de Pastas
- `public/`: Arquivos estáticos, imagens e HTML principal.
- `src/`: Código-fonte do frontend.
  - `assets/`: Imagens e recursos visuais.
  - `components/`: Componentes Vue reutilizáveis (Banner, BurgerForm, Dashboard, etc).
  - `router/`: Configuração das rotas da aplicação.
  - `views/`: Páginas principais (Home, Pedidos).

### Arquitetura
O frontend é baseado em componentes Vue, promovendo reutilização e organização do código. As rotas são gerenciadas pelo Vue Router, permitindo navegação entre páginas como Home e Pedidos.

### Instalação e Execução
1. Instale as dependências:
	```sh
	npm install
	```
2. Execute o servidor de desenvolvimento:
	```sh
	npm run serve
	```
3. Acesse a aplicação em `http://localhost:8080` (por padrão).

### Funcionalidades
- Montagem de hambúrguer personalizado (pão, carne, opcionais).
- Registro do nome do cliente.
- Visualização dos pedidos e seus status: "solicitado", "em produção", "finalizado", "cancelado".
- Cancelamento de pedidos.

---
Para detalhes sobre componentes e rotas, consulte os arquivos em `src/components/` e `src/router/`.
