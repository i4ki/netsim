package ipv4

type (
	IDH struct {
		Version uint8  // Internet version - 4-bit
		IHL     uint8  // Internet Header Length - 4-bit
		TOS     uint8  // Type of Service - 8-bit
		Len     uint16 // Total length - 16-bit
		Id      uint16 // Identification - 16-bit

		Flags    uint8  // 3-bit
		FragOff  uint16 // fragment offset - 13-bit
		TTL      uint8  // Time to Live - 8-bit
		Protocol uint8  // data protocol (tcp, udp, icmp, etc) - 8-bit
		Checksum uint16 // Header checksum - 16-bit
		SrcAddr  uint32 // Source address - 32-bit
		DstAddr  uint32 // Destination address - 32-bit

		OptType uint8
		OptLen  uint8
		Options []byte
	}
)
