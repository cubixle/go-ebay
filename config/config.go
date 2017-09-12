package config

const sandboxURL = "https://api.sandbox.ebay.com"
const productionURL = "https://api.ebay.com"

type Config struct {
	BaseUrl string

	DevId, AppId, CertId string
	RuName, AuthToken    string
	SiteId               int
}

func (e *Config) Sandbox() {
	e.BaseUrl = sandboxURL
}

func (e *Config) Production() {
	e.BaseUrl = productionURL
}
