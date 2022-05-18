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
  - NextJs
  - MariaDB
  - Docker
  - Github Actions
  - Scrum

----

# Backlog de Produto

Histórias de usuário em ordem de prioridade, da mais prioritária para a menos prioritária:

 
  1) Eu, **como administrador do Hatflix**, quero poder administrar as lojas do sistema da Hatflix.

  2) Eu, **como administrador de uma loja no Hatflix**, quero poder solicitar o cadastro da minha loja no sistema da Hatflix.

  3) Eu, **como administrador de uma loja no Hatflix**, quero poder administrar as configurações e o catálogo da minha loja.

  4) Eu, **como usuário do Hatflix**, quero poder visualizar todas as lojas de roupas disponíveis na Hatflix.

  5) Eu, **como administrador de uma loja no Hatflix**, quero poder administrar os pedidos dos meus clientes.

  6) Eu, **como usuário do Hatflix**, quero adicionar itens à um "Carrinho" para poder comprá-los em um mesmo pedido.

  7) Eu, **como usuário do Hatflix**, quero administrar minhas informações pessoais através de um sistema de login.

  8) Eu, **como usuário do Hatflix**, quero poder administrar meus pedidos.

  9) Eu, **como usuário do Hatflix**, quero escolher uma roupa específica à partir de uma categoria de vestimenta.

  10) Eu, **como usuário do Hatflix**, quero encontrar uma roupa específica de que sei o nome para poder comprá-la.


# Descrição do MVP

O MVP para o Hatflix será uma planilha constituída de duas páginas principais:

A primeira página, chamada "Hatflix", irá conter um aglomerado de lojas de roupa com os seus respectivos itens. Dessa forma, é possível procurar por roupas de uma forma unificada em todos os estabelecimentos disponíveis.

A segunda página, nomeada "Carrinho", será a lista de itens que o cliente deseja comprar. Essa lista possui os itens da primeira página e pode conter roupas de diferentes lojas. O carrinho deverá ser preenchido manualmente pelo cliente, assim como as informações necessárias para a entrega.

O backend será simulado manualmente. A lista de lojas e produtos deverá ser construída procurando em sites e atualizando de tempos em tempos. Para simular a compra, iremos consultar a segunda página da planilha, comprar os itens nas respectivas lojas e entregar ao cliente, cobrando o frete e o preço do produto.



