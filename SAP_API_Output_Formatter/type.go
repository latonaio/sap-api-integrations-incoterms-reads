package sap_api_output_formatter

type Incoterms struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	IncotermsCode string `json:"incoterms_code"`
	Deleted       string `json:"deleted"`
}

type IncotermsClassification struct {
	IncotermsClassification       string `json:"IncotermsClassification"`
	LocationIsMandatory           bool   `json:"LocationIsMandatory"`
	ToIncotermsClassificationText string `json:"to_IncotermsClassificationText"`
}

type ToIncotermsClassificationText struct {
	IncotermsClassification     string `json:"IncotermsClassification"`
	Language                    string `json:"Language"`
	IncotermsClassificationName string `json:"IncotermsClassificationName"`
}

type IncotermsVersion struct {
	IncotermsVersion       string `json:"IncotermsVersion"`
	ToIncotermsVersionText string `json:"to_IncotermsVersionText"`
}

type ToIncotermsVersionText struct {
	IncotermsVersion     string `json:"IncotermsVersion"`
	Language             string `json:"Language"`
	IncotermsVersionName string `json:"IncotermsVersionName"`
}
