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

// ListDDLPublishRecords invokes the dms_enterprise.ListDDLPublishRecords API synchronously
func (client *Client) ListDDLPublishRecords(request *ListDDLPublishRecordsRequest) (response *ListDDLPublishRecordsResponse, err error) {
	response = CreateListDDLPublishRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDDLPublishRecordsWithChan invokes the dms_enterprise.ListDDLPublishRecords API asynchronously
func (client *Client) ListDDLPublishRecordsWithChan(request *ListDDLPublishRecordsRequest) (<-chan *ListDDLPublishRecordsResponse, <-chan error) {
	responseChan := make(chan *ListDDLPublishRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDDLPublishRecords(request)
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

// ListDDLPublishRecordsWithCallback invokes the dms_enterprise.ListDDLPublishRecords API asynchronously
func (client *Client) ListDDLPublishRecordsWithCallback(request *ListDDLPublishRecordsRequest, callback func(response *ListDDLPublishRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDDLPublishRecordsResponse
		var err error
		defer close(result)
		response, err = client.ListDDLPublishRecords(request)
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

// ListDDLPublishRecordsRequest is the request struct for api ListDDLPublishRecords
type ListDDLPublishRecordsRequest struct {
	*requests.RpcRequest
	OrderId requests.Integer `position:"Query" name:"OrderId"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
}

// ListDDLPublishRecordsResponse is the response struct for api ListDDLPublishRecords
type ListDDLPublishRecordsResponse struct {
	*responses.BaseResponse
	RequestId            string             `json:"RequestId" xml:"RequestId"`
	ErrorCode            string             `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage         string             `json:"ErrorMessage" xml:"ErrorMessage"`
	Success              bool               `json:"Success" xml:"Success"`
	DDLPublishRecordList []DDLPublishRecord `json:"DDLPublishRecordList" xml:"DDLPublishRecordList"`
}

// CreateListDDLPublishRecordsRequest creates a request to invoke ListDDLPublishRecords API
func CreateListDDLPublishRecordsRequest() (request *ListDDLPublishRecordsRequest) {
	request = &ListDDLPublishRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDDLPublishRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateListDDLPublishRecordsResponse creates a response to parse from ListDDLPublishRecords response
func CreateListDDLPublishRecordsResponse() (response *ListDDLPublishRecordsResponse) {
	response = &ListDDLPublishRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
