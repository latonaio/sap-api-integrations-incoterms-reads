package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-incoterms-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToIncotermsClassification(raw []byte, l *logger.Logger) ([]IncotermsClassification, error) {
	pm := &responses.IncotermsClassification{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to IncotermsClassification. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	incotermsClassification := make([]IncotermsClassification, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		incotermsClassification = append(incotermsClassification, IncotermsClassification{
			IncotermsClassification:       data.IncotermsClassification,
			LocationIsMandatory:           data.LocationIsMandatory,
			ToIncotermsClassificationText: data.ToIncotermsClassificationText.Deferred.URI,
		})
	}

	return incotermsClassification, nil
}

func ConvertToToIncotermsClassificationText(raw []byte, l *logger.Logger) ([]ToIncotermsClassificationText, error) {
	pm := &responses.ToIncotermsClassificationText{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToIncotermsClassificationText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toIncotermsClassificationText := make([]ToIncotermsClassificationText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toIncotermsClassificationText = append(toIncotermsClassificationText, ToIncotermsClassificationText{
			IncotermsClassification:     data.IncotermsClassification,
			Language:                    data.Language,
			IncotermsClassificationName: data.IncotermsClassificationName,
		})
	}

	return toIncotermsClassificationText, nil
}

func ConvertToIncotermsVersion(raw []byte, l *logger.Logger) ([]IncotermsVersion, error) {
	pm := &responses.IncotermsVersion{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to IncotermsVersion. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	incotermsVersion := make([]IncotermsVersion, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		incotermsVersion = append(incotermsVersion, IncotermsVersion{
			IncotermsVersion:       data.IncotermsVersion,
			ToIncotermsVersionText: data.ToIncotermsVersionText.Deferred.URI,
		})
	}

	return incotermsVersion, nil
}

func ConvertToToIncotermsVersionText(raw []byte, l *logger.Logger) ([]ToIncotermsVersionText, error) {
	pm := &responses.ToIncotermsVersionText{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToIncotermsVersionText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toIncotermsVersionText := make([]ToIncotermsVersionText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toIncotermsVersionText = append(toIncotermsVersionText, ToIncotermsVersionText{
			IncotermsVersion:     data.IncotermsVersion,
			Language:             data.Language,
			IncotermsVersionName: data.IncotermsVersionName,
		})
	}

	return toIncotermsVersionText, nil
}
