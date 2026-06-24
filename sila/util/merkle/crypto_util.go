package merkle

import (
	"github.com/protolambda/ztyp/tree"
	"github.com/sila-chain/zrnt/sila/util/hashing"
)

// VerifyMerkleBranch verifies that the given leaf is
// on the merkle branch at the given depth, at the index at that depth.
func VerifyMerkleBranch(leaf tree.Root, branch []tree.Root, depth uint64, index uint64, root tree.Root) bool {
	value := leaf
	for i := uint64(0); i < depth; i++ {
		if (index>>i)&1 == 1 {
			value = hashing.Hash(append(branch[i][:], value[:]...))
		} else {
			value = hashing.Hash(append(value[:], branch[i][:]...))
		}
	}
	return value == root
}
