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

// CreateDownloadDataJob invokes the iot.CreateDownloadDataJob API synchronously
func (client *Client) CreateDownloadDataJob(request *CreateDownloadDataJobRequest) (response *CreateDownloadDataJobResponse, err error) {
	response = CreateCreateDownloadDataJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDownloadDataJobWithChan invokes the iot.CreateDownloadDataJob API asynchronously
func (client *Client) CreateDownloadDataJobWithChan(request *CreateDownloadDataJobRequest) (<-chan *CreateDownloadDataJobResponse, <-chan error) {
	responseChan := make(chan *CreateDownloadDataJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDownloadDataJob(request)
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

// CreateDownloadDataJobWithCallback invokes the iot.CreateDownloadDataJob API asynchronously
func (client *Client) CreateDownloadDataJobWithCallback(request *CreateDownloadDataJobRequest, callback func(response *CreateDownloadDataJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDownloadDataJobResponse
		var err error
		defer close(result)
		response, err = client.CreateDownloadDataJob(request)
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

// CreateDownloadDataJobRequest is the request struct for api CreateDownloadDataJob
type CreateDownloadDataJobRequest struct {
	*requests.RpcRequest
	StartTime        requests.Integer `position:"Query" name:"StartTime"`
	FileConfig       string           `position:"Query" name:"FileConfig"`
	IotInstanceId    string           `position:"Body" name:"IotInstanceId"`
	Context          string           `position:"Body" name:"Context"`
	TableName        string           `position:"Query" name:"TableName"`
	EndTime          requests.Integer `position:"Query" name:"EndTime"`
	DownloadDataType string           `position:"Query" name:"DownloadDataType"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
}

// CreateDownloadDataJobResponse is the response struct for api CreateDownloadDataJob
type CreateDownloadDataJobResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      bool                        `json:"Success" xml:"Success"`
	Code         string                      `json:"Code" xml:"Code"`
	ErrorMessage string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInCreateDownloadDataJob `json:"Data" xml:"Data"`
}

// CreateCreateDownloadDataJobRequest creates a request to invoke CreateDownloadDataJob API
func CreateCreateDownloadDataJobRequest() (request *CreateDownloadDataJobRequest) {
	request = &CreateDownloadDataJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateDownloadDataJob", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDownloadDataJobResponse creates a response to parse from CreateDownloadDataJob response
func CreateCreateDownloadDataJobResponse() (response *CreateDownloadDataJobResponse) {
	response = &CreateDownloadDataJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
