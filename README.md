<div align="center">
    <img src="https://github.com/Mugen-Builders/.github/assets/153661799/7ed08d4c-89f4-4bde-a635-0b332affbd5d" width="150" height="150">
</div>
<br>
<div align="center">
    <i>An example in Golang using Cartesi CLI, Nonodo, and Rollmelette</i>
</div>
<div align="center">
<b>This example aims to demonstrate the lifecycle of a Cartesi DApp through unit tests using Rollmelette as a framework.</b>
</div>
<br>
<div align="center">
    
  <a href="https://docs.cartesi.io/cartesi-rollups/">![Static Badge](https://img.shields.io/badge/cartesi-1.3.0-5bd1d7)</a>
  <a href="https://docs.cartesi.io/cartesi-rollups/1.3/quickstart/">![Static Badge](https://img.shields.io/badge/cartesi--cli-0.15.0-5bd1d7)</a>
  <a href="https://pkg.go.dev/github.com/calindra/nonodo">![Static Badge](https://img.shields.io/badge/nonodo-1.1.1-blue)</a>
  <a href="https://pkg.go.dev/github.com/gligneul/rollmelette">![Static Badge](https://img.shields.io/badge/rollmelette-0.1.1-yellow)</a>
  <a href="https://book.getfoundry.sh/getting-started/installation">![Static Badge](https://img.shields.io/badge/foundry-0.2.0-red)</a>
</div>

## User Stories:

Here is a list of user stories that the application covers:

| #   | User Story Description                                                                                     |
| --- | ---------------------------------------------------------------------------------------------------------- |
| 1   | As a user, I want to send Ether tokens to my wallet on Layer 2.                                           |
| 2   | As a user, I want to send ERC20 tokens to my wallet on Layer 2.                                           |
| 3   | As a user, I want to transfer Ether tokens between wallets on Layer 2.                                    |
| 4   | As a user, I want to transfer ERC20 tokens between wallets on Layer 2.                                    |
| 5   | As a user, I want to withdraw my deposit in ERC20.                                                        |
| 6   | As a user, I want to withdraw my deposit in Ether.                                                        |
| 7   | As a user, I want to request the balance of Ether in my wallet on Layer 2.                                |
| 8   | As a user, I want to request the balance of ERC20 tokens in my wallet on Layer 2.                         |
| 9   | As a user, I want to verify if the Ether deposit was received correctly on Layer 2.                       |
| 10  | As a user, I want to verify if the ERC20 token deposit was received correctly on Layer 2.                 |
| 11  | As a user, I want to receive a confirmation of Ether transfer between wallets on Layer 2.                 |
| 12  | As a user, I want to receive a confirmation of ERC20 token transfer between wallets on Layer 2.           |

## Setup:

#### The system setup is divided into three parts:
1º - Install all dependencies:
   + Cartesi CLI:
   ```bash
   $ npm i -g @cartesi/cli
   ```
   + Foundry:
   Follow the instruction [here](https://book.getfoundry.sh/getting-started/installation)

2º - Clone this repo using the code below:
```bash
git clone https://github.com/Mugen-Builders/learn-rollmelette.git
```

## Running the tests:
The command below will run all unit tests present in the repository.

```bash
make test
```

> [!NOTE]
> All user stories covered here can also be fulfilled using the CLI. For more information, please refer to the [documentation](https://docs.cartesi.io/cartesi-rollups/1.3/).
