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

// ModifyAssetGroup invokes the sas.ModifyAssetGroup API synchronously
func (client *Client) ModifyAssetGroup(request *ModifyAssetGroupRequest) (response *ModifyAssetGroupResponse, err error) {
	response = CreateModifyAssetGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAssetGroupWithChan invokes the sas.ModifyAssetGroup API asynchronously
func (client *Client) ModifyAssetGroupWithChan(request *ModifyAssetGroupRequest) (<-chan *ModifyAssetGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyAssetGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAssetGroup(request)
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

// ModifyAssetGroupWithCallback invokes the sas.ModifyAssetGroup API asynchronously
func (client *Client) ModifyAssetGroupWithCallback(request *ModifyAssetGroupRequest, callback func(response *ModifyAssetGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAssetGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyAssetGroup(request)
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

// ModifyAssetGroupRequest is the request struct for api ModifyAssetGroup
type ModifyAssetGroupRequest struct {
	*requests.RpcRequest
	GroupId  requests.Integer `position:"Query" name:"GroupId"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Uuids    string           `position:"Query" name:"Uuids"`
}

// ModifyAssetGroupResponse is the response struct for api ModifyAssetGroup
type ModifyAssetGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAssetGroupRequest creates a request to invoke ModifyAssetGroup API
func CreateModifyAssetGroupRequest() (request *ModifyAssetGroupRequest) {
	request = &ModifyAssetGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyAssetGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyAssetGroupResponse creates a response to parse from ModifyAssetGroup response
func CreateModifyAssetGroupResponse() (response *ModifyAssetGroupResponse) {
	response = &ModifyAssetGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
