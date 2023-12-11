package config

type Config struct {
	Database Database `json:"database"`
	Facility Facility `json:"facility"`
	OAuth    OAuth    `json:"oauth"`
	Server   Server   `json:"server"`
	Session  Session  `json:"session"`
}

type Database struct {
	Driver       string `json:"driver"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	AutoMigrate  bool   `json:"auto_migrate"`
	CACert       string `json:"ca_cert"`
}

type Facility struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type FacilityADH struct {
	APIBase  string `json:"api_base"`
	Rostered bool   `json:"rostered"`
}

type OAuth struct {
	Provider     string         `json:"provider"`
	BaseURL      string         `json:"base_url"`
	ClientID     string         `json:"client_id"`
	ClientSecret string         `json:"client_secret"`
	MyBaseURL    string         `json:"my_base_url"`
	Endpoints    OAuthEndpoints `json:"endpoints"`
}

type OAuthEndpoints struct {
	Authorization string `json:"authorize"`
	Token         string `json:"token"`
	Userinfo      string `json:"userinfo"`
}

type Server struct {
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

type Session struct {
	Algorithm     string `json:"algorithm"`
	Secret        string `json:"secret"`
	AccessExpire  int    `json:"access_expire"`
	RefreshExpire int    `json:"refresh_expire"`
}

type Weather struct {
	Airports  []string `json:"airports"`
	Frequency int      `json:"frequency"`
}