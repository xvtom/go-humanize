package humanize

import "testing"

func TestTransferSpeed(t *testing.T) {
	tests := []struct {
		in  uint64
		exp string
	}{
		{42, "42 bps"},
		{1000, "1.0 kbps"},
		{999900, "999.9 kbps"},
		{1000001, "1.0 Mbps"},
		{987654321, "987.7 Mbps"},
		{1000000005, "1.0 Gbps"},
		{123456789000, "123.5 Gbps"},
		{1234000000000, "1.2 Tbps"},
	}

	for _, p := range tests {
		got := TransferSpeed(p.in)
		if p.exp != got {
			t.Errorf("Expected '%v' for %v, got '%v'",
				p.exp, p.in, got)
		}
	}
}
