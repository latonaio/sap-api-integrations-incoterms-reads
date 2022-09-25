package responses

type IncotermsVersion struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			IncotermsVersion       string `json:"IncotermsVersion"`
			ToIncotermsVersionText struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_IncotermsVersionText"`
		} `json:"results"`
	} `json:"d"`
}
