package cfg

// Conf is the configuration for accessing bitnodes endpoint
type config struct {
	Port    string
	Path    string
	RPC     RPClogin
	CF      CloudFlare
	ApiKeys map[string]string
}
type RPClogin struct {
	Username, Password string
}
type CloudFlare struct {
	CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey string
}

// configurations for jorm
var (
	Path   = "/our/"
	CONFIG *config
)
