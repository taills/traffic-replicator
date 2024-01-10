package main

import (
	"flag"
	"fmt"
	traffic_replicator "github.com/taills/traffic-replicator"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	// Parse command line arguments
	var targets string
	var ports string
	var udp bool
	var tcp bool
	var showPacketAsASCII bool
	var showPacketAsHex bool

	flag.StringVar(&targets, "targets", "", "The target IP addresses, separated by comma")
	flag.StringVar(&ports, "ports", "", "The ports, separated by comma")
	flag.BoolVar(&udp, "udp", false, "Enable UDP")
	flag.BoolVar(&tcp, "tcp", false, "Enable TCP")
	flag.BoolVar(&showPacketAsASCII, "ascii", false, "Show packets as ASCII characters")
	flag.BoolVar(&showPacketAsHex, "hex", false, "Show packets as hex characters")

	flag.Parse()

	if targets == "" || ports == "" {
		fmt.Errorf("targets and ports must be specified")
		os.Exit(-1)
	}
	if !udp && !tcp {
		fmt.Errorf("either UDP or TCP must be enabled")
		os.Exit(-2)
	}

	// Create a new TrafficReplicator object

	_ports := strings.Split(ports, ",")
	portsInt := make([]int, 0)
	for _, v := range _ports {
		// if v has - in it, it's a range
		if strings.Contains(v, "-") {
			_range := strings.Split(v, "-")
			if len(_range) != 2 {
				panic("Invalid port range")
			}
			start, _ := strconv.Atoi(_range[0])
			end, _ := strconv.Atoi(_range[1])
			if start > end {
				panic("Invalid port range")
			}
			for j := start; j <= end; j++ {
				portsInt = append(portsInt, j)
			}
			continue
		} else {
			p, _ := strconv.Atoi(v)
			portsInt = append(portsInt, p)
		}
	}

	tr := traffic_replicator.NewTrafficReplicator(udp, tcp, portsInt, strings.Split(targets, ","), showPacketAsASCII, showPacketAsHex)

	var exitSignal = make(chan os.Signal)
	// register signal handler
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	tr.Start()
	<-exitSignal
	tr.Stop()
	os.Exit(0)
}
