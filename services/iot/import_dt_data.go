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

// ImportDTData invokes the iot.ImportDTData API synchronously
func (client *Client) ImportDTData(request *ImportDTDataRequest) (response *ImportDTDataResponse, err error) {
	response = CreateImportDTDataResponse()
	err = client.DoAction(request, response)
	return
}

// ImportDTDataWithChan invokes the iot.ImportDTData API asynchronously
func (client *Client) ImportDTDataWithChan(request *ImportDTDataRequest) (<-chan *ImportDTDataResponse, <-chan error) {
	responseChan := make(chan *ImportDTDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportDTData(request)
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

// ImportDTDataWithCallback invokes the iot.ImportDTData API asynchronously
func (client *Client) ImportDTDataWithCallback(request *ImportDTDataRequest, callback func(response *ImportDTDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportDTDataResponse
		var err error
		defer close(result)
		response, err = client.ImportDTData(request)
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

// ImportDTDataRequest is the request struct for api ImportDTData
type ImportDTDataRequest struct {
	*requests.RpcRequest
	RealTenantId      string               `position:"Body" name:"RealTenantId"`
	RealTripartiteKey string               `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string               `position:"Body" name:"IotInstanceId"`
	ProductKey        string               `position:"Body" name:"ProductKey"`
	ApiProduct        string               `position:"Body" name:"ApiProduct"`
	ApiRevision       string               `position:"Body" name:"ApiRevision"`
	Items             *[]ImportDTDataItems `position:"Body" name:"Items"  type:"Repeated"`
}

// ImportDTDataItems is a repeated param struct in ImportDTDataRequest
type ImportDTDataItems struct {
	DeviceName string `name:"DeviceName"`
	Params     string `name:"Params"`
}

// ImportDTDataResponse is the response struct for api ImportDTData
type ImportDTDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateImportDTDataRequest creates a request to invoke ImportDTData API
func CreateImportDTDataRequest() (request *ImportDTDataRequest) {
	request = &ImportDTDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ImportDTData", "", "")
	request.Method = requests.POST
	return
}

// CreateImportDTDataResponse creates a response to parse from ImportDTData response
func CreateImportDTDataResponse() (response *ImportDTDataResponse) {
	response = &ImportDTDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
