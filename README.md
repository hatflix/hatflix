# hatflix

![hatflix](https://user-images.githubusercontent.com/73971077/163252464-eb09a2fd-c548-4f00-819a-760ea9e56b1c.jpg)

Sistema para centralização de lojas de roupas.

Trabalho desenvolvido para a disciplina de Práticas em Desenvolvimento de Software - DCC - UFMG - 01/2022

Equipe:
  - Líder/Backend: Rafael Massoni Gonçalves Barbosa
  - Backend: Marco Antônio Silveira Souza Alves
  - DBA: Thomas Esteves
  - Frontend: Pedro Henrique Silva Gonçalves de Souza 

Tecnologias:
  - Golang
  - GraphQL
  - JavaScript
  - Tailwindcss
  - REACT
  - MariaDB
  - Docker
  - Github Actions
  - Scrum

----

# Backlog de Produto e Sprint

Histórias de usuário em ordem de prioridade, da mais prioritária para a menos prioritária:

  0) Tarefas técnicas
    - Criar o esquema de diretórios 2
    - Discutir e criar o schema do banco de dados 3
    - Selecionar a versão e Dockerfile do banco de dados 1
    - Discutir e mapear os endpoints da API de acordo com as telas planejadas 5
    - Configurar o gqlgen 5
    - Implementar a lib de database (wrapper lib) 3
    - Implementar o server do backend 3
    - Implementar os middlewares de auth (cookie JWT) 5
    - Modelar e implementar o sistema de cadastro e login, com níveis de permissão 8
    - Modelar e implementar o ElasticSearch como motor de busca da aplicação 5

  1) Eu, **como administrador do Hatflix**, quero poder administrar as lojas do sistema da Hatflix.
    - Projetar e implementar a tela de perfil do admin no frontend 5
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3
    
  2) Eu, **como administrador de uma loja no Hatflix**, quero poder solicitar o cadastro da minha loja no sistema da Hatflix.
    - Projetar e implementar a tela de cadastro no frontend 5
    - Implementar o endpoint na API 3

  3) Eu, **como administrador de uma loja no Hatflix**, quero poder administrar as configurações e o catálogo da minha loja.
    - Projetar e implementar a tela de perfil do admin de loja no frontend 5
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  4) Eu, **como usuário do Hatflix**, quero poder visualizar todas as lojas de roupas disponíveis na Hatflix.
    - Projetar e implementar a tela no frontend 5
    - Implementar o endpoint na API 3
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  5) Eu, **como administrador de uma loja no Hatflix**, quero poder administrar os pedidos dos meus clientes.
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  6) Eu, **como usuário do Hatflix**, quero adicionar itens à um "Carrinho" para poder comprá-los em um mesmo pedido.
    - Projetar e implementar a tela de "Carrinho" no frontend 5
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  7) Eu, **como usuário do Hatflix**, quero administrar minhas informações pessoais através de um sistema de login.
    - Projetar e implementar a tela de perfil do usuário no frontend 5
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  8) Eu, **como usuário do Hatflix**, quero poder administrar meus pedidos.
    - Projetar e implementar a tela de pedidos do usuário no frontend 5
    - Implementar o endpoint na API 3
    - Restringir acesso ao endpoint (checar JWT por nivel de permissão) 1
    - Implementar a lógica de negócio da tela no Service 3
    - Implementar a lógica de negócio no Repository 3

  9) Eu, **como usuário do Hatflix**, quero escolher uma roupa específica à partir de uma categoria de vestimenta.
    - Projetar e implementar a tela no frontend 5
    - Implementar o endpoint na API  3
    - No Service de roupas, adicionar suporte a filtros (categoria)
    - No Repository de roupas, adicionar suporte a filtros (categoria)

  10) Eu, **como usuário do Hatflix**, quero encontrar uma roupa específica de que sei o nome para poder comprá-la.
    - Projetar e implementar a barra de busca dinamica 8
    - Modelar e implementar o serviço de busca, consumindo dados do ElasticSearch 13
      - Ajustar e calibrar os pesos da busca (nome da roupa, nome da marca, categoria, modelo, etc) 
    
