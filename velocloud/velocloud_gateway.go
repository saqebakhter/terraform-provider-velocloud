package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GWSite struct {
	ContantName  string `json:"contactName"`
	ContantEmail string `json:"ContantEmail"`
}

type Gateway_insert_gateway struct {
	IPAddress        string `json:"ipAddress"`
	ServiceState     string `json:"serviceState"`
	PrivateIpAddress string `json:"privateIpAddress"`
	EndpointPkiMode  string `json:"endpointPkiMode"`
	IsLoadBalanced   int    `json:"isLoadBalanced"`
	GatewayPoolID    int    `json:"gatewayPoolId"`
	NetworkID        int    `json:"networkId"`
	Name             string `json:"name"`
	Site             GWSite `json:"site"`
}

type Gateway_insert_gateway_result struct {
	ID            int `json:"id"`
	ActivationKey int `json:"activationKey"`
}

type Gateway_delete_gateway struct {
	GatewayID int `json:"id"`
}

type Gateway_delete_gateway_result struct {
	ID   int `json:"id"`
	Rows int `json:"rows"`
}

// GetEnterprise ...
func GetGateway(c *Client, enterprisename string) (int, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/getNetworkGateways", c.HostURL), nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return 0, err
	}

	// Unmarschal
	var list []map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return 0, err
	}

	for _, v := range list {
		if v["name"] == enterprisename {

			return int(v["id"].(float64)), nil
		}
	}

	return 0, errors.New("cant find enterprise")

}

// InsertGateway ...
func InsertGateway(c *Client, body Gateway_insert_gateway) (Gateway_insert_gateway_result, error) {

	resp := Gateway_insert_gateway_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gateway/gatewayProvision", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}

// DeleteEnterprise ...
func DeleteGateway(c *Client, body Gateway_delete_gateway) (Gateway_delete_gateway_result, error) {

	resp := Gateway_delete_gateway_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gateway/deleteGateway", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}
