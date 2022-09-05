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

// CreateHoneypotNode invokes the sas.CreateHoneypotNode API synchronously
func (client *Client) CreateHoneypotNode(request *CreateHoneypotNodeRequest) (response *CreateHoneypotNodeResponse, err error) {
	response = CreateCreateHoneypotNodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHoneypotNodeWithChan invokes the sas.CreateHoneypotNode API asynchronously
func (client *Client) CreateHoneypotNodeWithChan(request *CreateHoneypotNodeRequest) (<-chan *CreateHoneypotNodeResponse, <-chan error) {
	responseChan := make(chan *CreateHoneypotNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHoneypotNode(request)
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

// CreateHoneypotNodeWithCallback invokes the sas.CreateHoneypotNode API asynchronously
func (client *Client) CreateHoneypotNodeWithCallback(request *CreateHoneypotNodeRequest, callback func(response *CreateHoneypotNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHoneypotNodeResponse
		var err error
		defer close(result)
		response, err = client.CreateHoneypotNode(request)
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

// CreateHoneypotNodeRequest is the request struct for api CreateHoneypotNode
type CreateHoneypotNodeRequest struct {
	*requests.RpcRequest
	AvailableProbeNum           requests.Integer `position:"Query" name:"AvailableProbeNum"`
	AllowHoneypotAccessInternet requests.Boolean `position:"Query" name:"AllowHoneypotAccessInternet"`
	NodeName                    string           `position:"Query" name:"NodeName"`
	SecurityGroupProbeIpList    *[]string        `position:"Query" name:"SecurityGroupProbeIpList"  type:"Repeated"`
	NodeVersion                 string           `position:"Query" name:"NodeVersion"`
}

// CreateHoneypotNodeResponse is the response struct for api CreateHoneypotNode
type CreateHoneypotNodeResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateHoneypotNodeRequest creates a request to invoke CreateHoneypotNode API
func CreateCreateHoneypotNodeRequest() (request *CreateHoneypotNodeRequest) {
	request = &CreateHoneypotNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "CreateHoneypotNode", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateHoneypotNodeResponse creates a response to parse from CreateHoneypotNode response
func CreateCreateHoneypotNodeResponse() (response *CreateHoneypotNodeResponse) {
	response = &CreateHoneypotNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
