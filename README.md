# Projeto Integração Gennera-Omie

Este projeto visa facilitar a integração entre as plataformas Gennera e Omie, automatizando o processo de cadastramento de clientes na Omie com base nos dados obtidos do Gennera.

## Funcionalidades

- **Consulta de Dados no Gennera**: Acessa a API do Gennera para obter informações de clientes a partir de uma lista de IDs.
- **Preparação dos Dados**: Processa e prepara os dados dos clientes para atender aos requisitos da API da Omie, incluindo a abreviação de endereços e a conversão dos nomes dos estados para siglas.
- **Cadastro de Clientes na Omie**: Envia os dados preparados para a API da Omie, cadastrando os clientes.
- **Gerenciamento de Erros**: Registra IDs de clientes para os quais o processo de cadastramento falhou, permitindo a reexecução ou análise manual.

## Como Utilizar

### Pré-requisitos

Antes de iniciar, você precisa ter as chaves de API (`app_key` e `app_secret`) tanto da Gennera quanto da Omie. Além disso, é necessário um token de acesso para a API do Gennera.

### Configuração

1. Clone o repositório para sua máquina local.
2. No diretório do projeto, crie um arquivo `.env` com as seguintes variáveis:
    ```plaintext
    GENNERA_TOKEN=seu_token_gennera
    APP_KEY=sua_app_key_omie
    APP_SECRET=seu_app_secret_omie
    ID_PERSON=id1,id2,id3
    ```
3. Certifique-se de que o Go está instalado em sua máquina e que todas as dependências do projeto estão satisfeitas.

### Execução

Para executar o programa, navegue até o diretório do projeto no terminal e execute:

```shell
go run main.go
