// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"bytes"
	"fmt"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/ids"
)

var (
	inboundSmallerValuePrefix = []byte{0}
	inboundSmallerIndexPrefix = []byte{1}
	inboundLargerValuePrefix  = []byte{2}
	inboundLargerIndexPrefix  = []byte{3}

	// inbound and outbound have their smaller and larger values swapped
	inbound = prefixes{
		smallerValuePrefix: inboundSmallerValuePrefix,
		smallerIndexPrefix: inboundSmallerIndexPrefix,
		largerValuePrefix:  inboundLargerValuePrefix,
		largerIndexPrefix:  inboundLargerIndexPrefix,
	}
	outbound = prefixes{
		smallerValuePrefix: inboundLargerValuePrefix,
		smallerIndexPrefix: inboundLargerIndexPrefix,
		largerValuePrefix:  inboundSmallerValuePrefix,
		largerIndexPrefix:  inboundSmallerIndexPrefix,
	}
)

type prefixes struct {
	smallerValuePrefix []byte
	smallerIndexPrefix []byte
	largerValuePrefix  []byte
	largerIndexPrefix  []byte
}

func (p *prefixes) getValueDB(myChainID, peerChainID ids.ID, db database.Database) database.Database {
	if bytes.Compare(myChainID[:], peerChainID[:]) == -1 {
		fmt.Println("smaller value")
		return prefixdb.New(p.smallerValuePrefix, db)
	}
	fmt.Println("larger value")
	return prefixdb.New(p.largerValuePrefix, db)
}

func (p *prefixes) getValueAndIndexDB(myChainID, peerChainID ids.ID, db database.Database) (database.Database, database.Database) {
	var valueDB, indexDB database.Database
	if bytes.Compare(myChainID[:], peerChainID[:]) == -1 {
		fmt.Println("smaller state")
		valueDB = prefixdb.New(p.smallerValuePrefix, db)
		indexDB = prefixdb.New(p.smallerIndexPrefix, db)
	} else {
		fmt.Println("larger state")
		valueDB = prefixdb.New(p.largerValuePrefix, db)
		indexDB = prefixdb.New(p.largerIndexPrefix, db)
	}
	return valueDB, indexDB
}