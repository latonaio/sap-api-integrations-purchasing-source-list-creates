package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-purchasing-source-list-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-purchasing-source-list-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncPostList(
	list *requests.PurchasingSourceList,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "List":
			func() {
				c.List(purchasingSourceList)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) List(purchasingSourceList *requests.PurchasingSourceList) {
	outputDataList, err := c.callPurchasingSourceListSrvAPIRequirementList("A_PurchasingSourceList", purchasingSourceList)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataList)
}

func (c *SAPAPICaller) callPurchasingSourceListSrvAPIRequirementList(api string, purchasingSourceList *requests.PurchasingSourceList) (*sap_api_output_formatter.List, error) {
	body, err := json.Marshal(purchasingSourceList)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PURCHASING_SOURCE_LIST_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
