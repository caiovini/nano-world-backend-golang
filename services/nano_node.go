package services

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"io/ioutil"
	"nano-world-backend-golang/payload"
	"net/http"
)

// sudo docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' {container_id_number}
const (
	url = "http://172.19.04:7076"
)

var (
	logger = gin.DefaultWriter
)

// GetFunds Fetch funds from nano node
func GetFunds(address string) (payload.RespBodyBalance, error) {

	var respBody payload.RespBodyBalance

	reqBody, err := json.Marshal(payload.ReqBodyBalance{
		Action:  "account_balance",
		Account: address,
	})

	if err != nil {
		logger.Write([]byte(err.Error()))
		return respBody, err
	}

	resp, err := http.Post(url,
		"text/plain", bytes.NewBuffer(reqBody))

	if err != nil {
		logger.Write([]byte(err.Error()))
		return respBody, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Write([]byte(err.Error()))
		return respBody, err
	}

	if err := json.Unmarshal(body, &respBody); err != nil { // Valid json

		logger.Write([]byte(err.Error()))
		return respBody, err
	}

	return respBody, nil
}

// GetPeers get all peers from node
func GetPeers() ([]byte, error) {

	reqBody, err := json.Marshal(payload.ReqBodyPeers{
		Action: "peers",
	})

	if err != nil {
		logger.Write([]byte(err.Error()))
		return nil, err
	}

	resp, err := http.Post(url,
		"text/plain", bytes.NewBuffer(reqBody))

	if err != nil {
		logger.Write([]byte(err.Error()))
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Write([]byte(err.Error()))
		return nil, err
	}

	return body, err
}
