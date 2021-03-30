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

// GetDataCronClearTaskDetailList invokes the dms_enterprise.GetDataCronClearTaskDetailList API synchronously
func (client *Client) GetDataCronClearTaskDetailList(request *GetDataCronClearTaskDetailListRequest) (response *GetDataCronClearTaskDetailListResponse, err error) {
	response = CreateGetDataCronClearTaskDetailListResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataCronClearTaskDetailListWithChan invokes the dms_enterprise.GetDataCronClearTaskDetailList API asynchronously
func (client *Client) GetDataCronClearTaskDetailListWithChan(request *GetDataCronClearTaskDetailListRequest) (<-chan *GetDataCronClearTaskDetailListResponse, <-chan error) {
	responseChan := make(chan *GetDataCronClearTaskDetailListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataCronClearTaskDetailList(request)
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

// GetDataCronClearTaskDetailListWithCallback invokes the dms_enterprise.GetDataCronClearTaskDetailList API asynchronously
func (client *Client) GetDataCronClearTaskDetailListWithCallback(request *GetDataCronClearTaskDetailListRequest, callback func(response *GetDataCronClearTaskDetailListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataCronClearTaskDetailListResponse
		var err error
		defer close(result)
		response, err = client.GetDataCronClearTaskDetailList(request)
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

// GetDataCronClearTaskDetailListRequest is the request struct for api GetDataCronClearTaskDetailList
type GetDataCronClearTaskDetailListRequest struct {
	*requests.RpcRequest
	OrderId    requests.Integer `position:"Query" name:"OrderId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// GetDataCronClearTaskDetailListResponse is the response struct for api GetDataCronClearTaskDetailList
type GetDataCronClearTaskDetailListResponse struct {
	*responses.BaseResponse
	TotalCount                  int64                     `json:"TotalCount" xml:"TotalCount"`
	RequestId                   string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode                   string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage                string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                     bool                      `json:"Success" xml:"Success"`
	DataCronClearTaskDetailList []DataCronClearTaskDetail `json:"DataCronClearTaskDetailList" xml:"DataCronClearTaskDetailList"`
}

// CreateGetDataCronClearTaskDetailListRequest creates a request to invoke GetDataCronClearTaskDetailList API
func CreateGetDataCronClearTaskDetailListRequest() (request *GetDataCronClearTaskDetailListRequest) {
	request = &GetDataCronClearTaskDetailListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataCronClearTaskDetailList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDataCronClearTaskDetailListResponse creates a response to parse from GetDataCronClearTaskDetailList response
func CreateGetDataCronClearTaskDetailListResponse() (response *GetDataCronClearTaskDetailListResponse) {
	response = &GetDataCronClearTaskDetailListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
