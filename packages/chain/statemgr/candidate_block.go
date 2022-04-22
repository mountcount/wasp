// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package statemgr

import (
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/trie"
	"github.com/iotaledger/wasp/packages/state"
)

type candidateBlock struct {
	block               state.Block
	local               bool
	votes               int
	approved            bool
	nextStateCommitment trie.VCommitment
	nextState           state.VirtualStateAccess
}

func newCandidateBlock(block state.Block, nextStateIfProvided state.VirtualStateAccess) *candidateBlock {
	var local bool
	var stateCommitment trie.VCommitment
	if nextStateIfProvided == nil {
		local = false
		stateCommitment = nil
	} else {
		local = true
		stateCommitment = trie.RootCommitment(nextStateIfProvided.TrieNodeStore())
	}
	return &candidateBlock{
		block:               block,
		local:               local,
		votes:               1,
		approved:            false,
		nextStateCommitment: stateCommitment,
		nextState:           nextStateIfProvided,
	}
}

func (cT *candidateBlock) getBlock() state.Block {
	return cT.block
}

func (cT *candidateBlock) addVote() {
	cT.votes++
}

func (cT *candidateBlock) getVotes() int {
	return cT.votes
}

func (cT *candidateBlock) isLocal() bool {
	return cT.local
}

func (cT *candidateBlock) isApproved() bool {
	return cT.approved
}

func (cT *candidateBlock) approveIfRightOutput(output *iscp.AliasOutputWithID) {
	if cT.block.BlockIndex() == output.GetStateIndex() {
		outputID := output.ID()
		finalL1Commitment, err := state.L1CommitmentFromAliasOutput(output.GetAliasOutput())
		if err != nil {
			return
		}
		finalCommitment := finalL1Commitment.StateCommitment
		if cT.isLocal() {
			if trie.EqualCommitments(cT.nextStateCommitment, finalCommitment) {
				cT.approved = true
				cT.block.SetApprovingOutputID(outputID)
			}
		} else {
			if cT.block.ApprovingOutputID().Equals(outputID) {
				cT.approved = true
				cT.nextStateCommitment = finalCommitment
			}
		}
	}
}

func (cT *candidateBlock) getNextStateCommitment() trie.VCommitment {
	return cT.nextStateCommitment
}

func (cT *candidateBlock) getNextState(currentState state.VirtualStateAccess) (state.VirtualStateAccess, error) {
	if cT.isLocal() {
		return cT.nextState.Copy(), nil
	}
	err := currentState.ApplyBlock(cT.block)
	return currentState, err
}

func (cT *candidateBlock) getApprovingOutputID() *iotago.UTXOInput {
	return cT.block.ApprovingOutputID()
}
