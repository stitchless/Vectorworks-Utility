package software

type SoftwareName = string
var HomeDirectory = GetHomeDirectory()

// Software holds the software name, and all the installations for that software package
type Software struct {
	Name          SoftwareName
	Installations []Installation
}

// Define some constants for use with our application
const (
	SoftwareVectorworks   SoftwareName = "Vectorworks"
	SoftwareVision        SoftwareName = "Vision"
	SoftwareCloudServices SoftwareName = "VCS"
)

// AllActiveSoftwareNames provides an easy means to add, remove, enable, and disable software options
// to be used by the overall package
var AllActiveSoftwareNames = []SoftwareName{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}