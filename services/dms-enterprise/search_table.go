package dms_enterprise

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

// SearchTable invokes the dms_enterprise.SearchTable API synchronously
func (client *Client) SearchTable(request *SearchTableRequest) (response *SearchTableResponse, err error) {
	response = CreateSearchTableResponse()
	err = client.DoAction(request, response)
	return
}

// SearchTableWithChan invokes the dms_enterprise.SearchTable API asynchronously
func (client *Client) SearchTableWithChan(request *SearchTableRequest) (<-chan *SearchTableResponse, <-chan error) {
	responseChan := make(chan *SearchTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchTable(request)
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

// SearchTableWithCallback invokes the dms_enterprise.SearchTable API asynchronously
func (client *Client) SearchTableWithCallback(request *SearchTableRequest, callback func(response *SearchTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchTableResponse
		var err error
		defer close(result)
		response, err = client.SearchTable(request)
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

// SearchTableRequest is the request struct for api SearchTable
type SearchTableRequest struct {
	*requests.RpcRequest
	ReturnGuid   requests.Boolean `position:"Query" name:"ReturnGuid"`
	SearchKey    string           `position:"Query" name:"SearchKey"`
	SearchRange  string           `position:"Query" name:"SearchRange"`
	Tid          requests.Integer `position:"Query" name:"Tid"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	SearchTarget string           `position:"Query" name:"SearchTarget"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EnvType      string           `position:"Query" name:"EnvType"`
	DbType       string           `position:"Query" name:"DbType"`
}

// SearchTableResponse is the response struct for api SearchTable
type SearchTableResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	ErrorMessage    string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode       string          `json:"ErrorCode" xml:"ErrorCode"`
	TotalCount      int64           `json:"TotalCount" xml:"TotalCount"`
	SearchTableList SearchTableList `json:"SearchTableList" xml:"SearchTableList"`
}

// CreateSearchTableRequest creates a request to invoke SearchTable API
func CreateSearchTableRequest() (request *SearchTableRequest) {
	request = &SearchTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SearchTable", "dmsenterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchTableResponse creates a response to parse from SearchTable response
func CreateSearchTableResponse() (response *SearchTableResponse) {
	response = &SearchTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
