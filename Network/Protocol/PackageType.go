package protocol

type PackageType int

const (
	PkgHeartbeat PackageType = 0
	PkgData      PackageType = 1
)
