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

// FeatureName illustrates all feature types
type FeatureName = string

// SoftwareName illustrates all Software Names
type SoftwareName = string

// All possible application features
const (
	FeatureTraceApplication FeatureName = "Trace Application"
	FeatureEditSerial       FeatureName = "Edit Serial"
	FeatureDemoWindow       FeatureName = "Demo Window"
)

// All possible SoftwareName (s)
const (
	SoftwareVectorworks   SoftwareName = "Vectorworks"
	SoftwareVision        SoftwareName = "Vision"
	SoftwareCloudServices SoftwareName = "VCS"
)

// AllActiveFeatures is used to turn on and off application features
// slice of FeatureName
var AllActiveFeatures = []FeatureName{
	FeatureTraceApplication,
	//FeatureEditSerial,
	FeatureDemoWindow,
}

// AllActiveSoftwareNames is used to turn on and off software to track
// slice of SoftwareName
var AllActiveSoftwareNames = []SoftwareName{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}
