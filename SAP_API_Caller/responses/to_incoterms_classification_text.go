package responses

type ToIncotermsClassificationText struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			IncotermsClassification     string `json:"IncotermsClassification"`
			Language                    string `json:"Language"`
			IncotermsClassificationName string `json:"IncotermsClassificationName"`
		} `json:"results"`
	} `json:"d"`
}
