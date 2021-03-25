package software

type SoftwareLabel = string

// Software holds the software name, and all the installations for that software package
type Software struct {
	Label         SoftwareLabel
	Installations []Installation
}

// Define some constants for use with out application
const (
	SoftwareVectorworks   SoftwareLabel = "Vectorworks"
	SoftwareVision        SoftwareLabel = "Vision"
	SoftwareCloudServices SoftwareLabel = "VCS"
)

// AllSoftwareOptions provides an easy means to add, remove, enable, and disable software options
// to be used by the overall package
var AllSoftwareOptions = []SoftwareLabel{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}
