// Copyright 2021 The psch Authors
// This file is part of the psch library.
//
// The psch library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The psch library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the psch library. If not, see <http://www.gnu.org/licenses/>.

package extrastate

import (
	"math/big"

	"github.com/petercastron/PSCH/common"
)

// journalEntry is a modification entry in the state change journal that can be
// reverted on demand.
type journalEntry interface {
	// revert undoes the changes introduced by this journal entry.
	revert(*ExtraStateDB)

	// dirtied returns the psch address modified by this journal entry.
	dirtied() *common.Address
}

// journal contains the list of state modifications applied since the last state
// commit. These are tracked to be able to be reverted in case of an execution
// exception or revertal request.
type journal struct {
	entries []journalEntry         // Current changes tracked by the journal
	dirties map[common.Address]int // Dirty accounts and the number of changes
}

// newJournal create a new initialized journal.
func newJournal() *journal {
	return &journal{
		dirties: make(map[common.Address]int),
	}
}

// append inserts a new modification entry to the end of the change journal.
func (j *journal) append(entry journalEntry) {
	j.entries = append(j.entries, entry)
	if addr := entry.dirtied(); addr != nil {
		j.dirties[*addr]++
	}
}

// revert undoes a batch of journalled modifications along with any reverted
// dirty handling too.
func (j *journal) revert(statedb *ExtraStateDB, snapshot int) {
	for i := len(j.entries) - 1; i >= snapshot; i-- {
		// Undo the changes made by the operation
		j.entries[i].revert(statedb)

		// Drop any dirty tracking induced by the change
		if addr := j.entries[i].dirtied(); addr != nil {
			if j.dirties[*addr]--; j.dirties[*addr] == 0 {
				delete(j.dirties, *addr)
			}
		}
	}
	j.entries = j.entries[:snapshot]
}

// dirty explicitly sets an address to dirty, even if the change entries would
// otherwise suggest it as clean. This method is an ugly hack to handle the RIPEMD
// precompile consensus exception.
func (j *journal) dirty(addr common.Address) {
	j.dirties[addr]++
}

// length returns the current number of entries in the journal.
func (j *journal) length() int {
	return len(j.entries)
}

type (
	// Changes to the account trie.
	createObjectChange struct {
		account *common.Address
	}
	resetObjectChange struct {
		prev         *stateObject
		prevdestruct bool
	}
	suicideChange struct {
		account     *common.Address
		prev        bool // whether account had already suicided
		prevbalance *big.Int
	}

	// Changes to individual accounts.
	balanceChange struct {
		account *common.Address
		prev    *big.Int
	}
	rewardChange struct {
		account   *common.Address
		prev      *big.Int
		prevQueue []*big.Int
	}
	releaseidxChange struct {
		account   *common.Address
		prev      uint16
	}
	numberChange struct {
		account   *common.Address
		prev      uint64
	}
	storageChange struct {
		account       *common.Address
		key, prevalue common.Hash
	}
	lockChange struct {
		account *common.Address
	}
	// Changes to other state values.
	touchChange struct {
		account *common.Address
	}
)

func (ch createObjectChange) revert(s *ExtraStateDB) {
	delete(s.stateObjects, *ch.account)
	delete(s.stateObjectsDirty, *ch.account)
}

func (ch createObjectChange) dirtied() *common.Address {
	return ch.account
}

func (ch resetObjectChange) revert(s *ExtraStateDB) {
	s.setStateObject(ch.prev)
	if !ch.prevdestruct && s.snap != nil {
		delete(s.snapDestructs, ch.prev.addrHash)
	}
}

func (ch resetObjectChange) dirtied() *common.Address {
	return nil
}

func (ch suicideChange) revert(s *ExtraStateDB) {
	obj := s.getStateObject(*ch.account)
	if obj != nil {
		obj.suicided = ch.prev
		obj.setBalance(ch.prevbalance)
	}
}

func (ch suicideChange) dirtied() *common.Address {
	return ch.account
}

var ripemd = common.HexToAddress("0000000000000000000000000000000000000003")

func (ch touchChange) revert(s *ExtraStateDB) {
}

func (ch touchChange) dirtied() *common.Address {
	return ch.account
}

func (ch balanceChange) revert(s *ExtraStateDB) {
	s.getStateObject(*ch.account).setBalance(ch.prev)
}

func (ch balanceChange) dirtied() *common.Address {
	return ch.account
}

func (ch rewardChange) revert(s *ExtraStateDB) {
	s.getStateObject(*ch.account).data.LockBalance = ch.prev
	s.getStateObject(*ch.account).data.ReleaseQueue = ch.prevQueue
}

func (ch rewardChange) dirtied() *common.Address {
	return ch.account
}

func (ch releaseidxChange) revert(s *ExtraStateDB) {
	s.getStateObject(*ch.account).data.ReleaseIdx = ch.prev
}

func (ch releaseidxChange) dirtied() *common.Address {
	return ch.account
}

func (ch numberChange) revert(s *ExtraStateDB) {
	s.getStateObject(*ch.account).data.Number = ch.prev
}

func (ch numberChange) dirtied() *common.Address {
	return ch.account
}

func (ch storageChange) dirtied() *common.Address {
	return ch.account
}
