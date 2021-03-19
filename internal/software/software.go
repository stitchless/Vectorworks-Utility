package software

type Software = string
const (
	SoftwareVectorworks   Software = "Vectorworks"
	SoftwareVision        Software = "Vision"
	SoftwareCloudServices Software = "VCS"
)

// All the softwares that the UI presents to the user.
var AllSoftwares = []Software{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}