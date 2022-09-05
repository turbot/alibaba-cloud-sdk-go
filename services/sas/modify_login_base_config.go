package sas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyLoginBaseConfig invokes the sas.ModifyLoginBaseConfig API synchronously
func (client *Client) ModifyLoginBaseConfig(request *ModifyLoginBaseConfigRequest) (response *ModifyLoginBaseConfigResponse, err error) {
	response = CreateModifyLoginBaseConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLoginBaseConfigWithChan invokes the sas.ModifyLoginBaseConfig API asynchronously
func (client *Client) ModifyLoginBaseConfigWithChan(request *ModifyLoginBaseConfigRequest) (<-chan *ModifyLoginBaseConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyLoginBaseConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoginBaseConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyLoginBaseConfigWithCallback invokes the sas.ModifyLoginBaseConfig API asynchronously
func (client *Client) ModifyLoginBaseConfigWithCallback(request *ModifyLoginBaseConfigRequest, callback func(response *ModifyLoginBaseConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoginBaseConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoginBaseConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyLoginBaseConfigRequest is the request struct for api ModifyLoginBaseConfig
type ModifyLoginBaseConfigRequest struct {
	*requests.RpcRequest
	Type     string `position:"Query" name:"Type"`
	Target   string `position:"Query" name:"Target"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Config   string `position:"Query" name:"Config"`
}

// ModifyLoginBaseConfigResponse is the response struct for api ModifyLoginBaseConfig
type ModifyLoginBaseConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLoginBaseConfigRequest creates a request to invoke ModifyLoginBaseConfig API
func CreateModifyLoginBaseConfigRequest() (request *ModifyLoginBaseConfigRequest) {
	request = &ModifyLoginBaseConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyLoginBaseConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyLoginBaseConfigResponse creates a response to parse from ModifyLoginBaseConfig response
func CreateModifyLoginBaseConfigResponse() (response *ModifyLoginBaseConfigResponse) {
	response = &ModifyLoginBaseConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
