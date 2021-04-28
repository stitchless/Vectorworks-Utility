package software

// HomeDirectory : Home Directory based on the running operating system.
var HomeDirectory = GetHomeDirectory()

// Software is all information that makes up a tracked piece of software
// Name provides the SoftwareName
// Installations is a slice of Installation
type Software struct {
	Name          SoftwareName
	Installations []Installation
}

// SoftwareName illustrates all Software Names
type SoftwareName = string

// All possible SoftwareName (s)
const (
	SoftwareVectorworks   SoftwareName = "Vectorworks"
	SoftwareVision        SoftwareName = "Vision"
	SoftwareCloudServices SoftwareName = "VCS"
)

// AllActiveSoftwareNames is used to turn on and off software to track
// slice of SoftwareName
var AllActiveSoftwareNames = []SoftwareName{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}
