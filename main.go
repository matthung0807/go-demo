package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	c := http.DefaultClient
	baseUrl := "https://api-staging.megaport.com"
	username := "john@abc.com"
	password := "Abcd@1234"
	accessToken, err := login(c, baseUrl, username, password)
	if err != nil {
		panic(err)
	}

	// mcr
	metro := "Hong Kong"
	location, err := getLocationId(c, baseUrl, metro, accessToken)
	if err != nil {
		panic(err)
	}
	validateMcrOrderReq := ValidateMcrOrderRequest{
		LocationId:  location.Id,
		ProductName: "mcr-1",
		Term:        1,
		ProductType: "MCR2",
		PortSpeed:   1000,
	}
	_, err = validateMcrOrder(c, baseUrl, accessToken, validateMcrOrderReq)
	if err != nil {
		panic(err)
	}

	buyMcrReq := BuyMcrRequest(validateMcrOrderReq)
	buyMcrData, err := buyMcr(c, baseUrl, accessToken, buyMcrReq)

	// aws vxc
	partnerPort, err := getPartnerPort(c, baseUrl, accessToken)
	if err != nil {
		panic(err)
	}

	validateAwsVxcOrderReq := ValidateVxcOrderRequest{
		ProductUid: buyMcrData.TechnicalServiceUid,
		AssociatedVxcs: []AssociatedVxc{
			{
				ProductName: "aws-vxc-1",
				RateLimit:   "50",
				AEnd: AEnd{
					Vlan: 101,
				},
				BEnd: BEnd{
					ProductUid: partnerPort.ProductUid,
					PartnerConfig: PartnerConfig{
						ConnectType:  "AWSHC",
						OwnerAccount: "123456789012",
						Name:         "aws-vxc-1",
						Prefixes:     "10.0.0.1/30",
					},
				},
			},
		},
	}

	_, err = validateVxcOrder(c, baseUrl, accessToken, validateAwsVxcOrderReq)
	buyAwsVxcReq := BuyVxcRequest(validateAwsVxcOrderReq)
	_, err = buyVxc(c, baseUrl, accessToken, buyAwsVxcReq)

	// gcp vxc
	googlePorts, err := getGooglePort(c, baseUrl, accessToken, "2288c4f0-1d89-4b26-97c2-8f7d7f71985a/asia-east1/1")
	googlePort := MegaPort{}
	for _, port := range googlePorts {
		if port.Country == "Hong Kong" {
			googlePort = port
		}
	}

	validateGcpVxcOrderReq := ValidateVxcOrderRequest{
		ProductUid: buyMcrData.TechnicalServiceUid,
		AssociatedVxcs: []AssociatedVxc{
			{
				ProductName: "gcp-vxc-1",
				RateLimit:   "50",
				AEnd: AEnd{
					Vlan: 90,
				},
				BEnd: BEnd{
					ProductUid: googlePort.ProductUid,
					PartnerConfig: PartnerConfig{
						ConnectType: "GOOGLE",
						PairingKey:  "2288c4f0-1d89-4b26-97c2-8f7d7f71985a/asia-east1/1",
					},
				},
			},
		},
	}

	_, err = validateVxcOrder(c, baseUrl, accessToken, validateGcpVxcOrderReq)
	buyGcpVxcReq := BuyVxcRequest(validateGcpVxcOrderReq)
	_, err = buyVxc(c, baseUrl, accessToken, buyGcpVxcReq)

}

type LoginResponse struct {
	Message string
	Data    LoginData
}

type LoginData struct {
	OAuthToken OAuthToken
}

type OAuthToken struct {
	AccessToken string
}

func login(c *http.Client, baseUrl, username, password string) (accessToken string, err error) {
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)

	loginUrl := fmt.Sprintf("%s/v2/login", baseUrl)
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	loginResp := LoginResponse{}
	json.Unmarshal(data, &loginResp)

	accessToken = loginResp.Data.OAuthToken.AccessToken
	return accessToken, nil
}

type LocationResponse struct {
	Data []LocationData
}

type LocationData struct {
	Id    int
	Name  string
	Metro string
}

func getLocationId(c *http.Client, baseUrl, metro, accessToken string) (location *LocationData, err error) {
	locationsUrl := fmt.Sprintf("%s/v2/locations", baseUrl)
	req, err := http.NewRequest("GET", locationsUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	q := req.URL.Query()
	q.Add("locationStatus", "Active")
	q.Add("metro", metro)
	req.URL.RawQuery = q.Encode()
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	locationResp := LocationResponse{}
	json.Unmarshal(data, &locationResp)

	for _, locationData := range locationResp.Data {
		if locationData.Id == 54 {
			return &locationData, nil
		}
	}
	return nil, nil
}

type ValidateMcrOrderRequest struct {
	LocationId  int    `json:"locationId"`
	ProductName string `json:"productName"`
	Term        int    `json:"term"`
	ProductType string `json:"productType"`
	PortSpeed   int    `json:"portSpeed"`
}

type ValidateMcrOrderResponse struct {
	Data []ValidateMcrOrderData
}

type ValidateMcrOrderData struct {
	ServiceName string
	ProductUid  string
	Price       ValidateMcrOrderPrice
}

type ValidateMcrOrderPrice struct {
	MonthlyRate float64
	MbpsRate    float64
	Currency    string
}

func validateMcrOrder(c *http.Client, baseUrl, accessToken string, validateMcrOrderReq ValidateMcrOrderRequest) (validateMcrOrderData *ValidateMcrOrderData, err error) {
	validateMcrOrderUrl := fmt.Sprintf("%s/v2/networkdesign/validate", baseUrl)

	validateMcrOrderReqs := make([]ValidateMcrOrderRequest, 0)
	validateMcrOrderReqs = append(validateMcrOrderReqs, validateMcrOrderReq)
	b, err := json.Marshal(validateMcrOrderReqs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", validateMcrOrderUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	validateMcrOrderResp := &ValidateMcrOrderResponse{}
	err = json.Unmarshal(data, &validateMcrOrderResp)
	if err != nil {
		return nil, err
	}

	if len(validateMcrOrderResp.Data) > 0 {
		validateMcrOrderData = &validateMcrOrderResp.Data[0]
	}

	return validateMcrOrderData, nil
}

type BuyMcrRequest struct {
	LocationId  int    `json:"locationId"`
	ProductName string `json:"productName"`
	Term        int    `json:"term"`
	ProductType string `json:"productType"`
	PortSpeed   int    `json:"portSpeed"`
}

type BuyMcrResponse struct {
	Data []BuyMcrData
}

type BuyMcrData struct {
	ServiceName         string
	Name                string
	TechnicalServiceUid string
	ProductType         string
	CreateDate          int
	RateType            string
	RateLimit           int
}

func buyMcr(c *http.Client, baseUrl, accessToken string, buyMcrRequest BuyMcrRequest) (buyMcrData *BuyMcrData, err error) {
	buyMcrUrl := fmt.Sprintf("%s/v2/networkdesign/buy", baseUrl)

	buyMcrReqs := make([]BuyMcrRequest, 0)
	buyMcrReqs = append(buyMcrReqs, buyMcrRequest)
	b, err := json.Marshal(buyMcrReqs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", buyMcrUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	buyMcrResp := &BuyMcrResponse{}
	err = json.Unmarshal(data, &buyMcrResp)
	if err != nil {
		return nil, err
	}

	if len(buyMcrResp.Data) > 0 {
		buyMcrData = &buyMcrResp.Data[0]
	}
	return buyMcrData, nil
}

type PartnerPortResponse struct {
	Data []PartnerPortData
}

type PartnerPortData struct {
	CspName       string
	ProductUid    string
	Title         string
	LocationId    int
	DiversityZone string
	ConnectType   string
}

func getPartnerPort(c *http.Client, baseUrl string, accessToken string) (partnerPortData *PartnerPortData, err error) {
	partnerPortUrl := fmt.Sprintf("%s/v2/dropdowns/partner/megaports", baseUrl)
	req, err := http.NewRequest("GET", partnerPortUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	q := req.URL.Query()
	q.Add("vxcPermitted", "true")
	q.Add("connectType", "AWSHC")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	partnerPortResp := PartnerPortResponse{}
	json.Unmarshal(data, &partnerPortResp)

	for _, partnerPortData := range partnerPortResp.Data {
		strs := strings.Split(partnerPortData.Title, " ")
		region := strings.TrimSuffix(strings.TrimPrefix(strs[len(strs)-1], "("), ")")
		if region == "ap-east-1" && partnerPortData.DiversityZone == "red" {
			return &partnerPortData, nil
		}
	}
	return
}

type ValidateVxcOrderRequest struct {
	ProductUid     string          `json:"productUid"`
	AssociatedVxcs []AssociatedVxc `json:"associatedVxcs"`
}

type AssociatedVxc struct {
	ProductName string `json:"productName"`
	RateLimit   string `json:"rateLimit"`
	AEnd        AEnd   `json:"aEnd"`
	BEnd        BEnd   `json:"bEnd"`
}

type AEnd struct {
	Vlan int `json:"vlan"`
}

type BEnd struct {
	ProductUid    string        `json:"productUid"`
	PartnerConfig PartnerConfig `json:"partnerConfig"`
}

type PartnerConfig struct {
	ConnectType  string `json:"connectType"`
	OwnerAccount string `json:"ownerAccount"` // for aws
	Name         string `json:"name"`         // for aws
	Prefixes     string `json:"prefixes"`     // for aws
	PairingKey   string `json:"pairingKey"`   // for gcp
}

type ValidateVxcOrderResponse struct {
	Data []ValidateVxcOrderData
}

type ValidateVxcOrderData struct {
	ServiceName string
	ProductUid  string
	Price       ValidateVxcOrderPrice
}

type ValidateVxcOrderPrice struct {
	MonthlyRate float64
	MbpsRate    float64
	Currency    string
	AEndRate    float64
	BEndRate    float64
}

func validateVxcOrder(c *http.Client, baseUrl, accessToken string, validateAwsVxcOrderReq ValidateVxcOrderRequest) (validateAwsVxcOrderData *ValidateVxcOrderData, err error) {
	validateAwsVxcOrderUrl := fmt.Sprintf("%s/v2/networkdesign/validate", baseUrl)

	validateAwsVxcOrderReqs := make([]ValidateVxcOrderRequest, 0)
	validateAwsVxcOrderReqs = append(validateAwsVxcOrderReqs, validateAwsVxcOrderReq)
	b, err := json.Marshal(validateAwsVxcOrderReqs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", validateAwsVxcOrderUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	validateAwsVxcOrderResp := &ValidateVxcOrderResponse{}
	err = json.Unmarshal(data, &validateAwsVxcOrderResp)
	if err != nil {
		return nil, err
	}

	if len(validateAwsVxcOrderResp.Data) > 0 {
		validateAwsVxcOrderData = &validateAwsVxcOrderResp.Data[0]
	}

	return validateAwsVxcOrderData, nil
}

type BuyVxcRequest struct {
	ProductUid     string          `json:"productUid"`
	AssociatedVxcs []AssociatedVxc `json:"associatedVxcs"`
}

type BuyVxcResponse struct {
	Data []BuyVxcData
}

type BuyVxcData struct {
	PayerMegaPortName       string
	PayerCompanyName        string
	PayerVlanId             int
	Speed                   int
	RateType                string
	VxcJTechnicalServiceUid string
	ConnectType             string
}

func buyVxc(c *http.Client, baseUrl, accessToken string, buyAwsVxcRequest BuyVxcRequest) (buyAwsVxcData *BuyVxcData, err error) {
	buyAwsVxcUrl := fmt.Sprintf("%s/v2/networkdesign/buy", baseUrl)

	buyAwsVxcReqs := make([]BuyVxcRequest, 0)
	buyAwsVxcReqs = append(buyAwsVxcReqs, buyAwsVxcRequest)
	b, err := json.MarshalIndent(buyAwsVxcReqs, "", "    ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", buyAwsVxcUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	buyAwsVxcResp := &BuyVxcResponse{}
	err = json.Unmarshal(data, &buyAwsVxcResp)
	if err != nil {
		return nil, err
	}

	if len(buyAwsVxcResp.Data) > 0 {
		buyAwsVxcData = &buyAwsVxcResp.Data[0]
	}
	return buyAwsVxcData, nil
}

type GooglePortResponse struct {
	Data GooglePortData
}

type GooglePortData struct {
	Bandwidths []int
	MegaPorts  []MegaPort
}

type MegaPort struct {
	ProductUid string
	Name       string
	PortSpeed  int
	Country    string
}

func getGooglePort(c *http.Client, baseUrl, accessToken, pairingKey string) (megaports []MegaPort, err error) {
	googlePortUrl := fmt.Sprintf("%s/v2/secure/google/%s", baseUrl, pairingKey)
	req, err := http.NewRequest("GET", googlePortUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	googlePortResp := &GooglePortResponse{}
	json.Unmarshal(data, googlePortResp)

	return googlePortResp.Data.MegaPorts, nil
}
