package dataworks_public

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

// ListQualityResultsByEntity invokes the dataworks_public.ListQualityResultsByEntity API synchronously
func (client *Client) ListQualityResultsByEntity(request *ListQualityResultsByEntityRequest) (response *ListQualityResultsByEntityResponse, err error) {
	response = CreateListQualityResultsByEntityResponse()
	err = client.DoAction(request, response)
	return
}

// ListQualityResultsByEntityWithChan invokes the dataworks_public.ListQualityResultsByEntity API asynchronously
func (client *Client) ListQualityResultsByEntityWithChan(request *ListQualityResultsByEntityRequest) (<-chan *ListQualityResultsByEntityResponse, <-chan error) {
	responseChan := make(chan *ListQualityResultsByEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQualityResultsByEntity(request)
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

// ListQualityResultsByEntityWithCallback invokes the dataworks_public.ListQualityResultsByEntity API asynchronously
func (client *Client) ListQualityResultsByEntityWithCallback(request *ListQualityResultsByEntityRequest, callback func(response *ListQualityResultsByEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQualityResultsByEntityResponse
		var err error
		defer close(result)
		response, err = client.ListQualityResultsByEntity(request)
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

// ListQualityResultsByEntityRequest is the request struct for api ListQualityResultsByEntity
type ListQualityResultsByEntityRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Body" name:"ProjectName"`
	EntityId    requests.Integer `position:"Body" name:"EntityId"`
	StartDate   string           `position:"Body" name:"StartDate"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	EndDate     string           `position:"Body" name:"EndDate"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
}

// ListQualityResultsByEntityResponse is the response struct for api ListQualityResultsByEntity
type ListQualityResultsByEntityResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string                           `json:"RequestId" xml:"RequestId"`
	Success        bool                             `json:"Success" xml:"Success"`
	ErrorCode      string                           `json:"ErrorCode" xml:"ErrorCode"`
	Data           DataInListQualityResultsByEntity `json:"Data" xml:"Data"`
}

// CreateListQualityResultsByEntityRequest creates a request to invoke ListQualityResultsByEntity API
func CreateListQualityResultsByEntityRequest() (request *ListQualityResultsByEntityRequest) {
	request = &ListQualityResultsByEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListQualityResultsByEntity", "", "")
	request.Method = requests.POST
	return
}

// CreateListQualityResultsByEntityResponse creates a response to parse from ListQualityResultsByEntity response
func CreateListQualityResultsByEntityResponse() (response *ListQualityResultsByEntityResponse) {
	response = &ListQualityResultsByEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
