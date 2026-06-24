package proto

import (
	"github.com/sila-chain/zrnt/sila/beacon/common"
	. "github.com/sila-chain/zrnt/sila/forkchoice"
)

func NewProtoForkChoice(spec *common.Spec, finalized Checkpoint, justified Checkpoint,
	anchorRoot Root, anchorSlot Slot, anchorParent Root,
	initialBalances []Gwei, sink NodeSink) (Forkchoice, error) {
	return NewForkChoice(spec, finalized, justified, anchorRoot, anchorSlot,
		NewProtoArray(anchorParent, anchorRoot, anchorSlot, justified.Epoch, finalized.Epoch, sink),
		NewProtoVoteStore(spec), initialBalances)
}
