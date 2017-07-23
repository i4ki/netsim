package ipv4

import (
	"bytes"
	"net"
	"testing"

	"golang.org/x/net/ipv4"
)

type (
	parseTestcase struct {
		desc        string
		dump        []byte
		expectedPkt ipv4.Header
		expectedErr string
	}
)

const TCPNum = 0x6

func compare(pkt1, pkt2 ipv4.Header) bool {
	dump1, err1 := pkt1.Marshal()
	dump2, err2 := pkt2.Marshal()
	if err1 != err2 {
		return false
	}
	return bytes.Equal(dump1, dump2)
}

func testPacket(t *testing.T, tc parseTestcase) {
	pkt := ipv4.Header{}
	err := pkt.Parse(tc.dump)
	if err != nil {
		if tc.expectedErr != err.Error() {
			t.Fatalf("Err differs: '%s' != '%s'", err.Error(),
				tc.expectedErr)
		}
		return
	}
	if !compare(pkt, tc.expectedPkt) {
		t.Fatalf("Pkt differs: \nPkt1: %s\nPkt2: %s", pkt.String(),
			tc.expectedPkt.String())
	}
}

func TestIPv4Parse(t *testing.T) {
	for _, tc := range []parseTestcase{
		{
			desc: "ethernet frame",
			dump: []byte{
				0x0, 0x15, 0x6d, 0xc4, 0x27, 0x4b, 0x54, 0x04,
				0xa6, 0x3c, 0xed, 0x2b, 0x08, 0x00,
			},
			expectedErr: "header too short",
		},
		{
			desc: "simple pkt",
			dump: []byte{
				0x45, 0x0, 0x0, 0x28, 0x3e, 0x8a, 0x40, 0x0,
				0x80, 0x06, 0x9b, 0xd2, 0xc0, 0xa8, 0x02, 0x65,
				0xc7, 0x3b, 0x96, 0x2a,
			},
			expectedPkt: ipv4.Header{
				Version:  0x4,
				Len:      20,
				TotalLen: 40,
				ID:       0x3e8a,
				Flags:    0x2,
				FragOff:  0x0,
				TTL:      128,
				Protocol: TCPNum,
				Checksum: 0x9bd2,
				Src:      net.ParseIP("192.168.2.101"),
				Dst:      net.ParseIP("199.59.150.42"),
			},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			testPacket(t, tc)
		})
	}
}
