# Desafios SmartMEI - Backend

## Requisitos

Desenvolver um serviço GraphQL que disponibilize a seguinte funcionalidade:

- Fazer um crawler para buscar o valor atual de uma Transferência do Plano Profissional no site da SmartMei (https://www.smartmei.com.br)

- Chamar uma API aberta que converta esse preço para USD (dolar americano) e EUR (Euro)

- Retornar as seguintes informações:
  
  - Data da consulta
  
  - Descrição da tarifa (como está no site da smartmei)

	- Valor em R$

  - Valor em USD

  - Valor em EUR

- Como parâmetro de entrada, receber o endereço do site da SmartMei (obrigatório)

- Uma sugestão de API de cotação de moedas é https://exchangeratesapi.io (o candidato pode usar essa ou outra se preferir)

- O candidato pode usar qualquer lib em NodeJS ou Golang que implemente o GraphQL

- O projeto final precisa ter o Playground habilitado

- Fazer ao menos um cenário de teste unitário (o framework é livre)

- O projeto deve ser feito em NodeJS (js ou typescript) ou Golang. A escolha é do candidato

## Entrega

- Pode ser um ZIP enviado por email ou disponibilizado para download

- Se o candidato preferir, pode salvar no Github e enviar o link do projeto

- Caso o candidato queira assumir alguma premissa não informada, basta detalhar e enviar junto com o projeto final