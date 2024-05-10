# Tickethub

Tickethub é uma API desenvolvida em Go que facilita a compra e venda de ingressos para eventos, integrada com o Stripe para processamento de pagamentos.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado em sua máquina:

- Go (versão 1.22 ou superior)

## Instalação

1. Clone o repositório do Tickethub:

    ```bash
    git clone https://github.com/leonardonicola/tickethub.git
    ```

2. Navegue até o diretório do projeto:

    ```bash
    cd tickethub
    ```

3. Instale as dependências do Go:

    ```bash
    go mod tidy
    ```

4. Configure todas as variáveis de ambiente que estão demonstradas no arquivo `.env.example` em um arquivo `.env`

5. Instale a biblioteca Air para ter hot reload em desenvolvimento:

  ```bash
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
  ```

## Executando o servidor

1. Para iniciar o servidor Tickethub em modo dev, execute o seguinte comando na raiz do projeto:

    ```bash
    air -c .air.toml
    ```
    
2. Para iniciar o servidor Tickethub em modo prod, execute o docker compose:

    ```bash
    docker compose up -d --build
    ```

## Uso da API

A API Tickethub expõe endpoints para compra e venda de ingressos, eventos e usuarios que estão documentadas no Swagger, disponiveis em: `/api/v1/swagger/index.html`

## Contribuindo

Sinta-se à vontade para contribuir com o Tickethub. Basta abrir uma issue ou enviar um pull request com suas melhorias.

## Licença

Este projeto é licenciado sob a Licença MIT. Consulte o arquivo `LICENSE` para obter mais detalhes.
