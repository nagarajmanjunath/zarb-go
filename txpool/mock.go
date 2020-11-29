package txpool

import (
	"github.com/zarbchain/zarb-go/crypto"
	"github.com/zarbchain/zarb-go/tx"
)

// For testing purpose

type MockTxPool struct {
	txs map[crypto.Hash]tx.Tx
}

func NewMockTxPool() *MockTxPool {
	return &MockTxPool{
		txs: make(map[crypto.Hash]tx.Tx),
	}
}

func (m *MockTxPool) PendingTx(hash crypto.Hash) *tx.Tx {
	tx, ok := m.txs[hash]
	if !ok {
		return nil
	}
	return &tx
}

func (m *MockTxPool) HasTx(hash crypto.Hash) bool {
	_, has := m.txs[hash]
	return has
}

func (m *MockTxPool) Size() int {
	return len(m.txs)
}

func (m *MockTxPool) Fingerprint() string {
	return ""
}

func (m *MockTxPool) AppendTxs(txs []tx.Tx) {
	for _, t := range txs {
		m.txs[t.Hash()] = t
	}
}

func (m *MockTxPool) AppendTx(tx tx.Tx) {
	m.txs[tx.Hash()] = tx

}
func (m *MockTxPool) AppendTxAndBroadcast(trx tx.Tx) {
	m.txs[trx.Hash()] = trx
}

func (m *MockTxPool) RemoveTx(hash crypto.Hash) *tx.Tx {
	tx := m.txs[hash]
	//	delete(m.txs, hash)
	return &tx
}