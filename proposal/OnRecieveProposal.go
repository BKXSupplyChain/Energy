package main

import (
		"./proposal"
		"github.com/BKXSupplyChain/Energy/tree/master/utils"
)


func main() {
// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	CheckNotFatal(err);

	prop := &pb.Proposal{}
	if err := proto.Unmarshal(in, book); err != nil {
        log.Fatalln("Failed to parse address book:", err)
	}
}
