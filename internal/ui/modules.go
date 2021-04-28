package ui

// Controls all features found within the application.


type FeatureName string

const (
	featureTraceApplication FeatureName = "Trace Application"
	featureShowSerial       FeatureName = "Edit Serial"
	featureDemoWindow       FeatureName = "Demo Window"
)

var AllActiveFeatures = []FeatureName {
	featureTraceApplication,
	featureShowSerial,
	featureDemoWindow,
}

var currentFeature FeatureName

var runonce bool