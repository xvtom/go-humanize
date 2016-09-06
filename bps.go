package humanize

import "fmt"

// TransferSpeed formats a network transfer speed to a string
// with best suited unit.
//
// For example:
//    TransferSpeed(12340) -> 12.34 kbps
func TransferSpeed(bps uint64) string {
	switch {
	case bps < 1000:
		return fmt.Sprintf("%d bps", bps)
	case bps < 1000000:
		kbps := float64(bps) / 1000
		return fmt.Sprintf("%.1f kbps", kbps)
	case bps < 1000000000:
		mbps := float64(bps) / 1000000
		return fmt.Sprintf("%.1f Mbps", mbps)
	case bps < 1000000000000:
		gbps := float64(bps) / 1000000000
		return fmt.Sprintf("%.1f Gbps", gbps)
	default:
		tbps := float64(bps) / 1000000000000
		return fmt.Sprintf("%.1f Tbps", tbps)
	}
}
