package main

import (
		"./proposal"
		"github.com/BKXSupplyChain/Energy/tree/master/utils"
)


func main() {
	prop := &pb.Proposal{}

	// TODO write to prop
	out, err := proto.Marshal(prop)
	CheckNotFatal(err);

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
        log.Fatalln("Failed to propose:", err)
	}
}
