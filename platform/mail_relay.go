package platform

type MailRelay struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
