package window

import "github.com/jpeizer/Vectorworks-Utility/internal/software"

func RunModule(feature software.FeatureName) {
	switch feature {
	case software.FeatureTraceApplication:
		ShowTraceApplication = !ShowTraceApplication
		RenderTraceApplication()
	}
}
