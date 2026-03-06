package tests

import (
	"testing"
)

func TestClusterSingleNode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping cluster tests in short mode")
	}

	// setup a cluster with a single node
	env := NewChainEnv(t, ClusterOptions{NumNodes: 1, Committee: []int{0}, DeactivateOnCleanup: true})

	// fails in CI
	t.Run("permissionless access node", env.testPermissionlessAccessNode)

	// fails in CI
	t.Run("spam evm", env.testSpamEVM)

	t.Run("spam onledger", env.testSpamOnledger)

	t.Run("accounts dump", env.testDumpAccounts)
}

func TestClusterMultiNodeCommittee(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping cluster tests in short modezxc")
	}

	// setup a cluster with 4 nodes
	env := NewChainEnv(t, ClusterOptions{NumNodes: 4, Committee: []int{0, 1, 2, 3}, DeactivateOnCleanup: true})

	t.Run("deploy basic", env.testDeployChain)

	t.Run("OffLedger Deposit,Withdraw,Transfer", env.testOffLedgerDepositWithdrawTransfer)

	t.Run("OnLedger Deposit", env.testOnLedgerDeposit)

	t.Run("EVM jsonrpc", env.testEVMJsonRPCCluster)

	t.Run("offledger nonce1", env.testOffLedgerNonce)

	t.Run("webapi ISC estimategas onledger", env.testEstimateGasOnLedger)
	t.Run("webapi ISC estimategas offledger", env.testEstimateGasOffLedger)
}

