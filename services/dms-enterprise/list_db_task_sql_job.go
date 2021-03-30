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

// ListDBTaskSQLJob invokes the dms_enterprise.ListDBTaskSQLJob API synchronously
func (client *Client) ListDBTaskSQLJob(request *ListDBTaskSQLJobRequest) (response *ListDBTaskSQLJobResponse, err error) {
	response = CreateListDBTaskSQLJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListDBTaskSQLJobWithChan invokes the dms_enterprise.ListDBTaskSQLJob API asynchronously
func (client *Client) ListDBTaskSQLJobWithChan(request *ListDBTaskSQLJobRequest) (<-chan *ListDBTaskSQLJobResponse, <-chan error) {
	responseChan := make(chan *ListDBTaskSQLJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDBTaskSQLJob(request)
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

// ListDBTaskSQLJobWithCallback invokes the dms_enterprise.ListDBTaskSQLJob API asynchronously
func (client *Client) ListDBTaskSQLJobWithCallback(request *ListDBTaskSQLJobRequest, callback func(response *ListDBTaskSQLJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDBTaskSQLJobResponse
		var err error
		defer close(result)
		response, err = client.ListDBTaskSQLJob(request)
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

// ListDBTaskSQLJobRequest is the request struct for api ListDBTaskSQLJob
type ListDBTaskSQLJobRequest struct {
	*requests.RpcRequest
	DBTaskGroupId requests.Integer `position:"Query" name:"DBTaskGroupId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListDBTaskSQLJobResponse is the response struct for api ListDBTaskSQLJob
type ListDBTaskSQLJobResponse struct {
	*responses.BaseResponse
	TotalCount       int64          `json:"TotalCount" xml:"TotalCount"`
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	ErrorCode        string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage     string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success          bool           `json:"Success" xml:"Success"`
	DBTaskSQLJobList []DBTaskSQLJob `json:"DBTaskSQLJobList" xml:"DBTaskSQLJobList"`
}

// CreateListDBTaskSQLJobRequest creates a request to invoke ListDBTaskSQLJob API
func CreateListDBTaskSQLJobRequest() (request *ListDBTaskSQLJobRequest) {
	request = &ListDBTaskSQLJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDBTaskSQLJob", "", "")
	request.Method = requests.POST
	return
}

// CreateListDBTaskSQLJobResponse creates a response to parse from ListDBTaskSQLJob response
func CreateListDBTaskSQLJobResponse() (response *ListDBTaskSQLJobResponse) {
	response = &ListDBTaskSQLJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
