package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-incoterms-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetIncoterms(incotermsClassification, incotermsVersion string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "IncotermsClassification":
			func() {
				c.IncotermsClassification(incotermsClassification)
				wg.Done()
			}()
		case "IncotermsVersion":
			func() {
				c.IncotermsVersion(incotermsVersion)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) IncotermsClassification(incotermsClassification string) {
	incotermsClassificationData, err := c.callIncotermsSrvAPIRequirementIncotermsClassification("A_IncotermsClassification", incotermsClassification)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(incotermsClassificationData)

	incotermsClassificationTextData, err := c.callToIncotermsClassificationText(incotermsClassificationData[0].ToIncotermsClassificationText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(incotermsClassificationTextData)
}

func (c *SAPAPICaller) callIncotermsSrvAPIRequirementIncotermsClassification(api, incotermsClassification string) ([]sap_api_output_formatter.IncotermsClassification, error) {
	url := strings.Join([]string{c.baseURL, "API_SD_INCOTERMS_SRV", api}, "/")
	param := c.getQueryWithIncotermsClassification(map[string]string{}, incotermsClassification)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIncotermsClassification(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToIncotermsClassificationText(url string) ([]sap_api_output_formatter.ToIncotermsClassificationText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToIncotermsClassificationText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) IncotermsVersion(incotermsVersion string) {
	incotermsVersionData, err := c.callIncotermsSrvAPIRequirementIncotermsVersion("A_IncotermsVersion", incotermsVersion)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(incotermsVersionData)

	incotermsClassificationTextData, err := c.callToIncotermsClassificationText(incotermsVersionData[0].ToIncotermsVersionText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(incotermsClassificationTextData)
}

func (c *SAPAPICaller) callIncotermsSrvAPIRequirementIncotermsVersion(api, incotermsVersion string) ([]sap_api_output_formatter.IncotermsVersion, error) {
	url := strings.Join([]string{c.baseURL, "API_SD_INCOTERMS_SRV", api}, "/")
	param := c.getQueryWithIncotermsVersion(map[string]string{}, incotermsVersion)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIncotermsVersion(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToIncotermsVersionText(url string) ([]sap_api_output_formatter.ToIncotermsVersionText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToIncotermsVersionText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithIncotermsClassification(params map[string]string, incotermsClassification string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("IncotermsClassification eq '%s'", incotermsClassification)
	return params
}

func (c *SAPAPICaller) getQueryWithIncotermsVersion(params map[string]string, incotermsVersion string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("IncotermsVersion eq '%s'", incotermsVersion)
	return params
}
