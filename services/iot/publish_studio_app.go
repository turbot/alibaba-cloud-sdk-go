package iot

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

// PublishStudioApp invokes the iot.PublishStudioApp API synchronously
func (client *Client) PublishStudioApp(request *PublishStudioAppRequest) (response *PublishStudioAppResponse, err error) {
	response = CreatePublishStudioAppResponse()
	err = client.DoAction(request, response)
	return
}

// PublishStudioAppWithChan invokes the iot.PublishStudioApp API asynchronously
func (client *Client) PublishStudioAppWithChan(request *PublishStudioAppRequest) (<-chan *PublishStudioAppResponse, <-chan error) {
	responseChan := make(chan *PublishStudioAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishStudioApp(request)
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

// PublishStudioAppWithCallback invokes the iot.PublishStudioApp API asynchronously
func (client *Client) PublishStudioAppWithCallback(request *PublishStudioAppRequest, callback func(response *PublishStudioAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishStudioAppResponse
		var err error
		defer close(result)
		response, err = client.PublishStudioApp(request)
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

// PublishStudioAppRequest is the request struct for api PublishStudioApp
type PublishStudioAppRequest struct {
	*requests.RpcRequest
	Description   string `position:"Body" name:"Description"`
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	ProjectId     string `position:"Body" name:"ProjectId"`
	AppId         string `position:"Body" name:"AppId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// PublishStudioAppResponse is the response struct for api PublishStudioApp
type PublishStudioAppResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         bool   `json:"Data" xml:"Data"`
}

// CreatePublishStudioAppRequest creates a request to invoke PublishStudioApp API
func CreatePublishStudioAppRequest() (request *PublishStudioAppRequest) {
	request = &PublishStudioAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "PublishStudioApp", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublishStudioAppResponse creates a response to parse from PublishStudioApp response
func CreatePublishStudioAppResponse() (response *PublishStudioAppResponse) {
	response = &PublishStudioAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
