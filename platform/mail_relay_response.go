package platform

type MailRelayResponse struct {
	Success bool      `json:"success"`
	Data    MailRelay `json:"data"`
}
