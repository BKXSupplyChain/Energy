package db

import (
	"fmt"
	"github.com/BKXSupplyChain/Energy/types"
	"github.com/globalsign/mgo/bson"
	"testing"
)

func checkSerialisation(doc interface{}, t *testing.T) {
	a, _ := bson.Marshal(doc)
	r := doc
	if doc != bson.Unmarshal(a, &r) {
		t.Error(fmt.Sprintf("(%v, %T)\n", doc, doc))
	}
}

func TestAdd(t *testing.T) {
	checkSerialisation(types.Proposal{"123", "312", 3.2, [4]byte{'a', 'b', 'c', 'd'}, 32132}, t)
}
