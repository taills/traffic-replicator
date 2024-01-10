package traffic_replicator

import (
	"fmt"
	"net"
)

// handleTCPConnection handles TCP connections
//
// Parameters:
//
// conn - the TCP connection
//
// Returns:
//
//	None
func (tr *TrafficReplicator) handleTCPConnection(conn *net.TCPConn, port int) {
	var connMap = make(map[string]*net.TCPConn)
	for _, targetIP := range tr.ReplicateTo {
		destAddr := fmt.Sprintf("%s:%d", targetIP, port)
		tcpAddr, err := net.ResolveTCPAddr("tcp", destAddr)
		if err != nil {
			fmt.Println("ResolveTCPAddr failed:", err)
			continue
		}

		destConn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println("DialTCP failed:", err)
			continue
		}
		connMap[targetIP] = destConn
	}

	buffer := make([]byte, 1024*1024)
	go func() {
		for {
			// read from source
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Errorf("read failed: %v", err)
				return
			}
			fmt.Printf("Received TCP packet from %s, length: %v\n", conn.RemoteAddr().String(), n)
			ShowPacket(buffer[:n], tr.ShowPacketAsASCII, tr.ShowPacketAsHex)
			// write to all destinations
			for targetIP, destConn := range connMap {
				_, err = destConn.Write(buffer[:n])
				if err != nil {
					fmt.Errorf("write failed: %v", err)
					return
				}
				fmt.Printf("Sent TCP packet to %s\n", targetIP)
			}
		}
	}()
	// read from first destination and write to source
	for {
		n, err := connMap[tr.ReplicateTo[0]].Read(buffer)
		if err != nil {
			fmt.Errorf("read failed: %v", err)
			return
		}
		fmt.Printf("Received TCP packet from %s, length: %v\n", tr.ReplicateTo[0], n)
		ShowPacket(buffer[:n], tr.ShowPacketAsASCII, tr.ShowPacketAsHex)
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Errorf("write failed: %v", err)
			return
		}
		fmt.Printf("Sent TCP packet to %s,length: %v\n", conn.RemoteAddr().String(), n)

	}
}
