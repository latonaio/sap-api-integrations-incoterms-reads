package responses

type IncotermsClassification struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			IncotermsClassification       string `json:"IncotermsClassification"`
			LocationIsMandatory           bool   `json:"LocationIsMandatory"`
			ToIncotermsClassificationText struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_IncotermsClassificationText"`
		} `json:"results"`
	} `json:"d"`
}
