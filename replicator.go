package traffic_replicator

import (
	"fmt"
	"net"
)

type TrafficReplicator struct {
	EnableUDP         bool     // Enable UDP traffic replication
	EnableTCP         bool     // Enable TCP traffic replication
	Ports             []int    // Ports to listen on
	ReplicateTo       []string // Addresses to replicate traffic to
	ShowPacketAsASCII bool     // Show packets as ASCII characters
	ShowPacketAsHex   bool     // Show packets as hex characters
	exitSignal        chan bool
}

// NewTrafficReplicator creates a new TrafficReplicator object
//
// Parameters:
//
//	enableUDP - Enable UDP traffic replication
//	enableTCP - Enable TCP traffic replication
//	ports - Ports to listen on
//	replicateTo - Addresses to replicate traffic to
//	showPackets - Show packets as they are received and sent
//	showPacketAsASCII - Show packets as ASCII characters
//	showPacketAsHex - Show packets as hex characters
//
// Returns:
//
//	A new TrafficReplicator object
func NewTrafficReplicator(enableUDP bool, enableTCP bool, ports []int, replicateTo []string, showPacketAsASCII bool, showPacketAsHex bool) *TrafficReplicator {
	return &TrafficReplicator{
		EnableUDP:         enableUDP,
		EnableTCP:         enableTCP,
		Ports:             ports,
		ReplicateTo:       replicateTo,
		ShowPacketAsASCII: showPacketAsASCII,
		ShowPacketAsHex:   showPacketAsHex,
		exitSignal:        make(chan bool),
	}
}

// Run starts the traffic replicator
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
func (tr *TrafficReplicator) Run() {
	if tr.EnableUDP {
		go tr.runUDP()
	}

	if tr.EnableTCP {
		go tr.runTCP()
	}

	<-tr.exitSignal
}

// Start starts the traffic replicator
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
func (tr *TrafficReplicator) Start() {
	go tr.Run()
}

// Stop stops the traffic replicator
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
func (tr *TrafficReplicator) Stop() {
	tr.exitSignal <- true
}

// runUDP starts the UDP traffic replicator
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
func (tr *TrafficReplicator) runUDP() {
	for _, port := range tr.Ports {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			fmt.Println("ResolveUDPAddr failed:", err)
			continue
		}

		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println("ListenUDP failed:", err)
			continue
		}

		go tr.handleUDPConnection(conn, port)
	}
}

// runTCP starts the TCP traffic replicator
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
func (tr *TrafficReplicator) runTCP() {
	for _, port := range tr.Ports {
		go func(port int) {
			addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				fmt.Println("ResolveTCPAddr failed:", err)
				return
			}
			fmt.Println("Listening on TCP port", port)
			listener, err := net.ListenTCP("tcp", addr)
			if err != nil {
				fmt.Println("ListenTCP failed:", err)
				return
			}
			for {
				select {
				case <-tr.exitSignal:
					return
				default:
					conn, err := listener.AcceptTCP()
					if err != nil {
						fmt.Println("AcceptTCP failed:", err)
						return
					}
					fmt.Println("Accepted TCP connection")
					go tr.handleTCPConnection(conn, port)
				}
			}
		}(port)

	}
}
