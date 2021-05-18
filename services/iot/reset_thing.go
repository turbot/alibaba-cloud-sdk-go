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

// ResetThing invokes the iot.ResetThing API synchronously
func (client *Client) ResetThing(request *ResetThingRequest) (response *ResetThingResponse, err error) {
	response = CreateResetThingResponse()
	err = client.DoAction(request, response)
	return
}

// ResetThingWithChan invokes the iot.ResetThing API asynchronously
func (client *Client) ResetThingWithChan(request *ResetThingRequest) (<-chan *ResetThingResponse, <-chan error) {
	responseChan := make(chan *ResetThingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetThing(request)
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

// ResetThingWithCallback invokes the iot.ResetThing API asynchronously
func (client *Client) ResetThingWithCallback(request *ResetThingRequest, callback func(response *ResetThingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetThingResponse
		var err error
		defer close(result)
		response, err = client.ResetThing(request)
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

// ResetThingRequest is the request struct for api ResetThing
type ResetThingRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// ResetThingResponse is the response struct for api ResetThing
type ResetThingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateResetThingRequest creates a request to invoke ResetThing API
func CreateResetThingRequest() (request *ResetThingRequest) {
	request = &ResetThingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ResetThing", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetThingResponse creates a response to parse from ResetThing response
func CreateResetThingResponse() (response *ResetThingResponse) {
	response = &ResetThingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
