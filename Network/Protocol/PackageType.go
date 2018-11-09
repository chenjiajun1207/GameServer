package protocol

type PackageType int

const (
	PKG_HANDSHAKE     PackageType = 1
	PKG_HANDSHAKE_ACK PackageType = 2
	PkgHeartbeat      PackageType = 3
	PKG_DATA          PackageType = 4
	PKG_KICK          PackageType = 5
)
