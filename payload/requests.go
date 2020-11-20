package payload

// ReqBodyBalance is a request to the wallet balance
type ReqBodyBalance struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

// ReqBodyPeers is a request to the peers
type ReqBodyPeers struct {
	Action string `json:"action"`
}
