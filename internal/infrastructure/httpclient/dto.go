package httpclient

type responseBody struct {
	UlCount int `json:"ul_count"`
	Ul      []struct {
		AciID string `json:"aci_id"`
	} `json:"ul"`
}
