package software

// License provides a parsed breakdown of a serial number, including the serial itself
type License struct {
	Serial     string `default:"" json:"serial"`
	Activation string `default:"" json:"activation"`
	Local      string `default:"" json:"local"`
	Platform   string `default:"" json:"platform"`
	Type       string `default:"" json:"type"`
}

var licenseActivationMap = map[string]string{
	"A": "Retired",
	"B": "Local Dongle",
	"C": "Network Dongle",
	"E": "Serial Activation",
	"G": "Network Server",
	"U": "Updater",
	"V": "Viewer",
}

var licensePlatformMap = map[string]string{
	"W": "Windows",
	"M": "Mac",
	"X": "Cross Platform",
}

var licenseLocalMap = map[string]string{
	"US": "United States",
	"MK": "United Kingdom",
	"NZ": "New Zealand",
	"ZC": "Australia",
	"MR": "Canada (Resolve)",
	"CA": "Canada (Paxar)",
	"CL": "China",
	"BZ": "Brazil",
	"BE": "Belgium",
}

var licenseTypeMap = map[string]string{
	"N": "Not For Resale (Retail)",
	"E": "Educational (Pro Format)",
	"S": "Student (Pro Format)",
	"U": "Student (Student Format)",
	"T": "Teacher (Pro Format)",
	"C": "Teacher (Student Format)",
	"A": "Internal",
}
