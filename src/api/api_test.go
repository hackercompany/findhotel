package api

import (
	"config"

	"github.com/stretchr/testify/assert"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApi(t *testing.T) {
	assert := assert.New(t)
	config.DoInit()
	respPayload := Response{}

	client := &http.Client{}
	req := http.NewRequest("GET", config.Config.GetString("server.host")+":"+config.Config.GetString("server.port")+"/ipinformation?ip=127.0.0..1")

	resp, err := client.Do(req)
	assert.Nil(err)

	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, respPayload)
	assert.Nil(err)

	assert.Equal(respPayload.Status, false, "Should not pass for malformed payload")
}
