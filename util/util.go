package util

func PortsMatch(p1, p2 string) bool {
	if p1 == p2 {
		return true
	}
	return false
}

func FriendlyAddress(address string) string {
	var friendlyAddress string
	if len(address) == 0 {
		friendlyAddress = "0.0.0.0"
	} else {
		friendlyAddress = address
	}
	return friendlyAddress
}
