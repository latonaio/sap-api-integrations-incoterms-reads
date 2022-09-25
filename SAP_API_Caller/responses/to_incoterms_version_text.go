package responses

type ToIncotermsVersionText struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			IncotermsVersion     string `json:"IncotermsVersion"`
			Language             string `json:"Language"`
			IncotermsVersionName string `json:"IncotermsVersionName"`
		} `json:"results"`
	} `json:"d"`
}
