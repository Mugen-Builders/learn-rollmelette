<div align="center">
    <img src="https://github.com/Mugen-Builders/.github/assets/153661799/7ed08d4c-89f4-4bde-a635-0b332affbd5d" width="150" height="150">
</div>
<br>
<div align="center">
    <i>Um exemplo em Golang usando Cartesi Cli, Nonodo e Rollmelette como framework</i>
</div>
<div align="center">
Este exemplo tem como objetivo demonstrar o ciclo de vida de um DApp Cartesi por meio de testes unitários usando o Rollmelette como framework.
</div>
<br>
<div align="center">
    
  <a href="">[![Static Badge](https://img.shields.io/badge/cartesi-1.3.0-5bd1d7)](https://docs.cartesi.io/cartesi-rollups/)</a>
  <a href="">[![Static Badge](https://img.shields.io/badge/cartesi--cli-0.15.0-5bd1d7)](https://docs.cartesi.io/cartesi-rollups/1.3/quickstart/)</a>
  <a href="">[![Static Badge](https://img.shields.io/badge/nonodo-1.1.1-blue)](https://pkg.go.dev/github.com/calindra/nonodo)</a>
  <a href="">[![Static Badge](https://img.shields.io/badge/rollmelette-0.1.1-yellow)](https://pkg.go.dev/github.com/gligneul/rollmelette)</a>
  <a href="">[![Static Badge](https://img.shields.io/badge/foundry-0.2.0-red)](https://book.getfoundry.sh/getting-started/installation)</a>
</div>

## User Stories:

Here is a list of user stories that the application covers:

| #   | Descrição da História de Usuário                                                                         |
| --- | -------------------------------------------------------------------------------------------------------- |
| 1   | Como usuário, quero enviar tokens Ether para minha carteira na Layer 2.                                  |
| 2   | Como usuário, quero enviar tokens ERC20 para minha carteira na Layer 2.                                  |
| 3   | Como usuário, quero transferir tokens Ether entre carteiras na Layer 2.                                  |
| 4   | Como usuário, quero transferir tokens ERC20 entre carteiras na Layer 2.                                  |
| 5   | Como usuário, quero retirar meu depósito em ERC20.                                                       |
| 6   | Como usuário, quero retirar meu depósito em Ether.                                                       |
| 7   | Como usuário, quero solicitar o saldo de Ether na minha carteira na Layer 2.                             |
| 8   | Como usuário, quero solicitar o saldo de tokens ERC20 na minha carteira na Layer 2.                      |
| 9   | Como usuário, quero verificar se o depósito de Ether foi recebido corretamente na Layer 2.               |
| 10  | Como usuário, quero verificar se o depósito de tokens ERC20 foi recebido corretamente na Layer 2.        |
| 11  | Como usuário, quero receber uma confirmação de transferência de Ether entre carteiras na Layer 2.        |
| 12  | Como usuário, quero receber uma confirmação de transferência de tokens ERC20 entre carteiras na Layer 2. |

## Configuração:

#### A configuração do sistema é dividida em três partes:
1º - Instale todas as dependências:
   + Cartesi Cli:
   ```bash
   $ npm i -g @cartesi/cli
   ```
   + Foundry:
   Siga as instruções da própria [documentação da dependência](https://book.getfoundry.sh/)

2º - Clone este repositório usando o código abaixo:
```Bash
git clone --recursive git@github.com:Mugen-Builders/learn-deroll.git
```

## Rodando os testes:
O comando abaixo vai rodar todos os testes unitário presentes no repositório.

```bash
make test
```

> [!NOTE]
> Todas as user stories aqui realizadas também podem ser cumpridas por meio da cli. Para isso, siga a referência presente na [documentação](https://docs.cartesi.io/cartesi-rollups/1.3/).