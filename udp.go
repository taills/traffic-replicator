package traffic_replicator

import (
	"fmt"
	"net"
)

// handleUDPConnection handles UDP connections
//
// Parameters:
//
//	conn - the UDP connection
//
// Returns:
//
//	None
func (tr *TrafficReplicator) handleUDPConnection(conn *net.UDPConn, port int) {
	buffer := make([]byte, 1024*1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Errorf("ReadFromUDP from %v failed: %v", conn, err)
			continue
		}

		fmt.Printf("Received UDP packet from %s，length %v\n", addr.String(), n)

		for _, targetIP := range tr.ReplicateTo {
			destAddr := fmt.Sprintf("%s:%d", targetIP, port)
			remoteAddr, err := net.ResolveUDPAddr("udp", destAddr)
			if err != nil {
				fmt.Errorf("ResolveUDPAddr failed: %v", err)
				continue
			}

			_, err = conn.WriteToUDP(buffer[:n], remoteAddr)
			if err != nil {
				fmt.Errorf("WriteToUDP %v failed: %v", remoteAddr, err)
				continue
			}

			ShowPacket(buffer[:n], tr.ShowPacketAsASCII, tr.ShowPacketAsHex)

			fmt.Printf("Sent UDP packet to %s\n", destAddr)
		}
	}
}
