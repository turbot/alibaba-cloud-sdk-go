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

// ExportRecord invokes the sas.ExportRecord API synchronously
func (client *Client) ExportRecord(request *ExportRecordRequest) (response *ExportRecordResponse, err error) {
	response = CreateExportRecordResponse()
	err = client.DoAction(request, response)
	return
}

// ExportRecordWithChan invokes the sas.ExportRecord API asynchronously
func (client *Client) ExportRecordWithChan(request *ExportRecordRequest) (<-chan *ExportRecordResponse, <-chan error) {
	responseChan := make(chan *ExportRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportRecord(request)
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

// ExportRecordWithCallback invokes the sas.ExportRecord API asynchronously
func (client *Client) ExportRecordWithCallback(request *ExportRecordRequest, callback func(response *ExportRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportRecordResponse
		var err error
		defer close(result)
		response, err = client.ExportRecord(request)
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

// ExportRecordRequest is the request struct for api ExportRecord
type ExportRecordRequest struct {
	*requests.RpcRequest
	ExportType string `position:"Query" name:"ExportType"`
	Params     string `position:"Query" name:"Params"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// ExportRecordResponse is the response struct for api ExportRecord
type ExportRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	FileName  string `json:"FileName" xml:"FileName"`
	Id        int64  `json:"Id" xml:"Id"`
}

// CreateExportRecordRequest creates a request to invoke ExportRecord API
func CreateExportRecordRequest() (request *ExportRecordRequest) {
	request = &ExportRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ExportRecord", "", "")
	request.Method = requests.POST
	return
}

// CreateExportRecordResponse creates a response to parse from ExportRecord response
func CreateExportRecordResponse() (response *ExportRecordResponse) {
	response = &ExportRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
