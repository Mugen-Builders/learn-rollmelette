package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/Mugen-Builders/learn-rollmelette/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
)

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

type RouterSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *RouterSuite) SetupTest() {
	app := NewApp()
	s.tester = rollmelette.NewTester(app)
}

func (s *RouterSuite) TestItEchoAdvance() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	input := []byte(`{"path":"echo","payload":{"message":"Hello, Rollmelette!"}}`)
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(`{"message":"Hello, Rollmelette!"}`, string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItDepositEther() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"path":"depositEther"}`)
	advanceResult := s.tester.DepositEther(sender, big.NewInt(1000000000000000000), payload)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving deposit from %v with value %v", sender, big.NewInt(1000000000000000000).String()), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItTransferEther() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	depositPayload := []byte(`{"path":"depositEther"}`)
	depositAdvanceResult := s.tester.DepositEther(sender, big.NewInt(100), depositPayload)
	s.Len(depositAdvanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving deposit from %v with value %v", sender, big.NewInt(100).String()), string(depositAdvanceResult.Notices[0].Payload))

	recipient := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	tranferPayload := []byte(fmt.Sprintf(`{"to":"%s","value":10}`, recipient.Hex()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "transferEther",
		Payload: tranferPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("transfering from %v to %v value %v", sender, recipient, "10"), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItWithdrawEther() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	depositPayload := []byte(`{"path":"depositEther"}`)
	depositAdvanceResult := s.tester.DepositEther(sender, big.NewInt(100), depositPayload)
	s.Len(depositAdvanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving deposit from %v with value %v", sender, big.NewInt(100).String()), string(depositAdvanceResult.Notices[0].Payload))

	payload := []byte(`{"value":100}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawEther",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("withdrawing from %v value %v", sender, "100"), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItDepositERC20() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tokenAddress := common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2")
	payload := []byte(`{"path":"depositERC20"}`)
	advanceResult := s.tester.DepositERC20(tokenAddress, sender, big.NewInt(10000), payload)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving ERC20 deposit from %v with token %v and value %v", sender, tokenAddress, "10000"), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItTransferERC20() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tokenAddress := common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2")
	depositPayload := []byte(`{"path":"depositERC20"}`)
	depositAdvanceResult := s.tester.DepositERC20(tokenAddress, sender, big.NewInt(10000), depositPayload)
	s.Len(depositAdvanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving ERC20 deposit from %v with token %v and value %v", sender, tokenAddress, "10000"), string(depositAdvanceResult.Notices[0].Payload))

	recipient := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	payload := []byte(fmt.Sprintf(`{"token":"%s","to":"%s","value":100}`, tokenAddress.Hex(), recipient.Hex()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "transferERC20",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("transferring ERC20 from %v to %v token %v value %v", sender, recipient, tokenAddress, "100"), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItWithdrawERC20() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tokenAddress := common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2")
	depositPayload := []byte(`{"path":"depositERC20"}`)
	depositAdvanceResult := s.tester.DepositERC20(tokenAddress, sender, big.NewInt(10000), depositPayload)
	s.Len(depositAdvanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving ERC20 deposit from %v with token %v and value %v", sender, tokenAddress, "10000"), string(depositAdvanceResult.Notices[0].Payload))

	payload := []byte(fmt.Sprintf(`{"token":"%s","value":100}`, tokenAddress.Hex()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawERC20",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("withdrawing ERC20 from %v token %v value %v", sender, tokenAddress, "100"), string(advanceResult.Notices[0].Payload))
}

func (s *RouterSuite) TestItInspectEtherBalance() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"path":"depositEther"}`)
	advanceResult := s.tester.DepositEther(sender, big.NewInt(1000000000000000000), payload)
	s.Len(advanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving deposit from %v with value %v", sender, big.NewInt(1000000000000000000).String()), string(advanceResult.Notices[0].Payload))

	inspectResult := s.tester.Inspect([]byte("/balance/0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"))
	s.Equal("1000000000000000000", string(inspectResult.Reports[0].Payload))
}

func (s *RouterSuite) TestItInspectERC20Balance() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	tokenAddress := common.HexToAddress("0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2")
	depositPayload := []byte(`{"path":"depositERC20"}`)
	depositAdvanceResult := s.tester.DepositERC20(tokenAddress, sender, big.NewInt(10000), depositPayload)
	s.Len(depositAdvanceResult.Notices, 1)
	s.Equal(fmt.Sprintf("receiving ERC20 deposit from %v with token %v and value %v", sender, tokenAddress, "10000"), string(depositAdvanceResult.Notices[0].Payload))
	path := []byte("/balance/0x92C6bcA388E99d6B304f1Af3c3Cd749Ff0b591e2/0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	
	inspectResult := s.tester.Inspect(path)
	s.Equal("10000", string(inspectResult.Reports[0].Payload))
}
