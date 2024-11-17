package domain

type IPResponse struct {
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Readme   string `json:"readme"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}
