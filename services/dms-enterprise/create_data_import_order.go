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

// CreateDataImportOrder invokes the dms_enterprise.CreateDataImportOrder API synchronously
func (client *Client) CreateDataImportOrder(request *CreateDataImportOrderRequest) (response *CreateDataImportOrderResponse, err error) {
	response = CreateCreateDataImportOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataImportOrderWithChan invokes the dms_enterprise.CreateDataImportOrder API asynchronously
func (client *Client) CreateDataImportOrderWithChan(request *CreateDataImportOrderRequest) (<-chan *CreateDataImportOrderResponse, <-chan error) {
	responseChan := make(chan *CreateDataImportOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataImportOrder(request)
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

// CreateDataImportOrderWithCallback invokes the dms_enterprise.CreateDataImportOrder API asynchronously
func (client *Client) CreateDataImportOrderWithCallback(request *CreateDataImportOrderRequest, callback func(response *CreateDataImportOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataImportOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateDataImportOrder(request)
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

// CreateDataImportOrderRequest is the request struct for api CreateDataImportOrder
type CreateDataImportOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer `position:"Query" name:"Tid"`
	AttachmentKey   string           `position:"Query" name:"AttachmentKey"`
	Param           string           `position:"Query" name:"Param"`
	Comment         string           `position:"Query" name:"Comment"`
	RelatedUserList string           `position:"Query" name:"RelatedUserList"`
}

// CreateDataImportOrderResponse is the response struct for api CreateDataImportOrder
type CreateDataImportOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	Success           bool    `json:"Success" xml:"Success"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateDataImportOrderRequest creates a request to invoke CreateDataImportOrder API
func CreateCreateDataImportOrderRequest() (request *CreateDataImportOrderRequest) {
	request = &CreateDataImportOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateDataImportOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDataImportOrderResponse creates a response to parse from CreateDataImportOrder response
func CreateCreateDataImportOrderResponse() (response *CreateDataImportOrderResponse) {
	response = &CreateDataImportOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
