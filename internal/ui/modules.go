package ui

// featureName provides the user readable string for a supported software package
type featureName string

// hard coded feature names that are possible for all implemented software packages
const (
	featureTraceApplication featureName = "Trace Application"
	featureSoftware         featureName = "Software"
	featureDemoWindow       featureName = "Demo Window"
	featureSettings         featureName = "Settings"
)

// AllActiveFeatures is a list of all the currently supported features the application supports
var AllActiveFeatures = []featureName{
	featureTraceApplication,
	featureSoftware,
	featureDemoWindow,
	featureSettings,
}

// currentFeature is used to control the flow of actively rendered features
var currentFeature featureName
