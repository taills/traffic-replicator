package traffic_replicator

import (
	"encoding/hex"
	"fmt"
)

// ShowPacket shows the packet
//
// Parameters:
//
//	data - the packet data
//
//	asASCII - whether to show the packet as ASCII
//
//	asHex - whether to show the packet as hex
//
// Returns:
//
//	None
func ShowPacket(data []byte, asASCII bool, asHex bool) {
	if asASCII {
		fmt.Println("ASCII:")
		fmt.Println(string(data))
	}
	if asHex {
		fmt.Println("Hex:")
		fmt.Println(hex.Dump(data))
	}
}
