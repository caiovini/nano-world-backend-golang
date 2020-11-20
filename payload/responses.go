package payload

// RespBodyBalance gets response from nano node
type RespBodyBalance struct {
	Balance string `json:"balance"`
	Pending string `json:"pending"`
}
