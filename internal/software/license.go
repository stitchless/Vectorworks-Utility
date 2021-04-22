package software

// License provides a parsed breakdown of a serial number, including the serial itself
type License struct {
	Serial   string `default:"" json:"serial"`
	Local    string `default:"" json:"local"`
	Platform string `default:"" json:"platform"`
	Type     string `default:"" json:"type"`
}