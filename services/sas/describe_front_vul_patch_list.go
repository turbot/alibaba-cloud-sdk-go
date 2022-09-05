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

// DescribeFrontVulPatchList invokes the sas.DescribeFrontVulPatchList API synchronously
func (client *Client) DescribeFrontVulPatchList(request *DescribeFrontVulPatchListRequest) (response *DescribeFrontVulPatchListResponse, err error) {
	response = CreateDescribeFrontVulPatchListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFrontVulPatchListWithChan invokes the sas.DescribeFrontVulPatchList API asynchronously
func (client *Client) DescribeFrontVulPatchListWithChan(request *DescribeFrontVulPatchListRequest) (<-chan *DescribeFrontVulPatchListResponse, <-chan error) {
	responseChan := make(chan *DescribeFrontVulPatchListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFrontVulPatchList(request)
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

// DescribeFrontVulPatchListWithCallback invokes the sas.DescribeFrontVulPatchList API asynchronously
func (client *Client) DescribeFrontVulPatchListWithCallback(request *DescribeFrontVulPatchListRequest, callback func(response *DescribeFrontVulPatchListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFrontVulPatchListResponse
		var err error
		defer close(result)
		response, err = client.DescribeFrontVulPatchList(request)
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

// DescribeFrontVulPatchListRequest is the request struct for api DescribeFrontVulPatchList
type DescribeFrontVulPatchListRequest struct {
	*requests.RpcRequest
	Type        string `position:"Query" name:"Type"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	OperateType string `position:"Query" name:"OperateType"`
	Info        string `position:"Query" name:"Info"`
}

// DescribeFrontVulPatchListResponse is the response struct for api DescribeFrontVulPatchList
type DescribeFrontVulPatchListResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	FrontPatchList []FrontPatch `json:"FrontPatchList" xml:"FrontPatchList"`
}

// CreateDescribeFrontVulPatchListRequest creates a request to invoke DescribeFrontVulPatchList API
func CreateDescribeFrontVulPatchListRequest() (request *DescribeFrontVulPatchListRequest) {
	request = &DescribeFrontVulPatchListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeFrontVulPatchList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeFrontVulPatchListResponse creates a response to parse from DescribeFrontVulPatchList response
func CreateDescribeFrontVulPatchListResponse() (response *DescribeFrontVulPatchListResponse) {
	response = &DescribeFrontVulPatchListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
