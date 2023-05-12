package tun_tap

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/net-byte/water"
	"golang.org/x/exp/slog"
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = "TAPTEST"

	iface, err := water.New(config)
	if err != nil {
		slog.Error("create iface error", err)
		return
	}
	b := make([]byte, 65535)

	for {
		n, err := iface.Read(b)
		if err != nil {
			slog.Error("copy error", err)
			return
		}
		slog.Info("length of bytes", n)
		packet := gopacket.NewPacket(b[:n], layers.LayerTypeEthernet, gopacket.Default)
		//if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		//	fmt.Println("This is a TCP packet!")
		//	// Get actual TCP data from this layer
		//	tcp, _ := tcpLayer.(*layers.TCP)
		//	fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
		//}
		if ICMPLayer := packet.Layer(layers.LayerTypeICMPv4); ICMPLayer != nil {
			fmt.Println("This is a ICMP packet!")
			// Get actual TCP data from this layer
			ICMP, _ := ICMPLayer.(*layers.ICMPv4)
			fmt.Println(ICMP)
		}
	}
}
