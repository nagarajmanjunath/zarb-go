package consensus

import (
	"github.com/zarbchain/zarb-go/consensus/hrs"
	"github.com/zarbchain/zarb-go/vote"
)

type ConsensusReader interface {
	PickRandomVote(round int) *vote.Vote
	RoundVotes(round int) []*vote.Vote
	RoundProposal(round int) *vote.Proposal
	HRS() hrs.HRS
	Fingerprint() string
}

type Consensus interface {
	ConsensusReader

	MoveToNewHeight()
	Stop()
	AddVote(v *vote.Vote)
	SetProposal(proposal *vote.Proposal)
}
