package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/Mugen-Builders/learn-rollmelette/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

func NewApp() *router.Router {

	//////////////////////// Router //////////////////////////
	dapp := router.NewRouter()

	//////////////////////// Advance //////////////////////////
	dapp.HandleAdvance("echo", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		slog.Info("received payload", "payload", string(payload))
		env.Notice(payload)
		return nil
	})

	dapp.HandleAdvance("depositEther", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		etherDeposit, ok := deposit.(*rollmelette.EtherDeposit)
		if etherDeposit == nil || !ok {
			return fmt.Errorf("unsupported deposit type for bid creation: %T", deposit)
		}
		env.Notice([]byte(fmt.Sprintf("receiving deposit from %v with value %v", etherDeposit.Sender.String(), etherDeposit.Value)))
		return nil
	})

	dapp.HandleAdvance("transferEther", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		var input struct {
			To    string `json:"to"`
			Value int64  `json:"value"`
		}
		if err := json.Unmarshal(payload, &input); err != nil {
			return fmt.Errorf("failed to unmarshal payload: %w", err)
		}
		if env.EtherBalanceOf(metadata.MsgSender).Cmp(big.NewInt(input.Value)) < 0 {
			return fmt.Errorf("insufficient balance: %v < %v", env.EtherBalanceOf(metadata.MsgSender), input.Value)
		}
		env.EtherTransfer(metadata.MsgSender, common.HexToAddress(input.To), big.NewInt(input.Value))
		env.Notice([]byte(fmt.Sprintf("transfering from %v to %v value %v", metadata.MsgSender, input.To, input.Value)))
		return nil
	})

	dapp.HandleAdvance("withdrawEther", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		var input struct {
			Value int64  `json:"value"`
		}
		if err := json.Unmarshal(payload, &input); err != nil {
			return fmt.Errorf("failed to unmarshal payload: %w", err)
		}
		if env.EtherBalanceOf(metadata.MsgSender).Cmp(big.NewInt(input.Value)) < 0 {
			return fmt.Errorf("insufficient balance: %v < %v", env.EtherBalanceOf(metadata.MsgSender), input.Value)
		}
		env.EtherWithdraw(metadata.MsgSender, big.NewInt(input.Value))
		env.Notice([]byte(fmt.Sprintf("withdrawing from %v value %v", metadata.MsgSender, input.Value)))
		return nil
	})

	dapp.HandleAdvance("depositERC20", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		erc20Deposit, ok := deposit.(*rollmelette.ERC20Deposit)
		if erc20Deposit == nil || !ok {
			return fmt.Errorf("unsupported deposit type for ERC20 deposit: %T", deposit)
		}

		if erc20Deposit.Token != common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2") {
			return fmt.Errorf("unsupported token address: %v", erc20Deposit.Token)
		}
		
		env.Notice([]byte(fmt.Sprintf("receiving ERC20 deposit from %v with token %v and value %v", erc20Deposit.Sender, erc20Deposit.Token, erc20Deposit.Amount)))
		slog.Info("received ERC20 deposit", "deposit", deposit)
		return nil
	})

	dapp.HandleAdvance("transferERC20", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		var input struct {
			Token string `json:"token"`
			To    string `json:"to"`
			Value int64  `json:"value"`
		}
		if err := json.Unmarshal(payload, &input); err != nil {
			return fmt.Errorf("failed to unmarshal payload: %w", err)
		}
		to := common.HexToAddress(input.To)
		tokenAddress := common.HexToAddress(input.Token)
		value := big.NewInt(input.Value)

		if tokenAddress != common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2") {
			return fmt.Errorf("unsupported token address: %v", input.Token)
		}

		if err := env.ERC20Transfer(tokenAddress, metadata.MsgSender, to, value); err != nil {
			return fmt.Errorf("failed to transfer ERC20 tokens: %w", err)
		}
		env.Notice([]byte(fmt.Sprintf("transferring ERC20 from %v to %v token %v value %v", metadata.MsgSender, input.To, input.Token, input.Value)))
		return nil
	})

	dapp.HandleAdvance("withdrawERC20", func(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
		var input struct {
			Token string `json:"token"`
			Value int64  `json:"value"`
		}
		if err := json.Unmarshal(payload, &input); err != nil {
			return fmt.Errorf("failed to unmarshal payload: %w", err)
		}
		tokenAddress := common.HexToAddress(input.Token)
		value := big.NewInt(input.Value)

		if env.ERC20BalanceOf(tokenAddress, metadata.MsgSender).Cmp(value) < 0 {
			return fmt.Errorf("insufficient balance: %v < %v", env.ERC20BalanceOf(tokenAddress, metadata.MsgSender), value)
		}

		if tokenAddress != common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2") {
			return fmt.Errorf("unsupported token address: %v", input.Token)
		}

		if _, err := env.ERC20Withdraw(tokenAddress, metadata.MsgSender, value); err != nil {
			return fmt.Errorf("failed to withdraw ERC20 tokens: %w", err)
		}
		env.Notice([]byte(fmt.Sprintf("withdrawing ERC20 from %v token %v value %v", metadata.MsgSender, input.Token, input.Value)))
		return nil
	})

	//////////////////////// Inspect //////////////////////////
	dapp.HandleInspect("/echo/{data}", func(env rollmelette.EnvInspector, ctx context.Context) error {
		data := router.PathValue(ctx, "data")
		env.Report([]byte(data))
		return nil
	})

	dapp.HandleInspect("/balance/{token}/{address}", func(env rollmelette.EnvInspector, ctx context.Context) error {
		balance := env.ERC20BalanceOf(common.HexToAddress(router.PathValue(ctx, "token")), common.HexToAddress(router.PathValue(ctx, "address")))
		if common.HexToAddress(router.PathValue(ctx, "token")) != common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2") {
			return fmt.Errorf("failed to get balance for given address: %v", router.PathValue(ctx, "address"))
		}
		env.Report([]byte(balance.String()))
		return nil
	})

	dapp.HandleInspect("/balance/{address}", func(env rollmelette.EnvInspector, ctx context.Context) error {
		balance := env.EtherBalanceOf(common.HexToAddress(router.PathValue(ctx, "address")))
		env.Report([]byte(balance.String()))
		return nil
	})
	return dapp
}
