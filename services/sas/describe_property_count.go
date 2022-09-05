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

// DescribePropertyCount invokes the sas.DescribePropertyCount API synchronously
func (client *Client) DescribePropertyCount(request *DescribePropertyCountRequest) (response *DescribePropertyCountResponse, err error) {
	response = CreateDescribePropertyCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertyCountWithChan invokes the sas.DescribePropertyCount API asynchronously
func (client *Client) DescribePropertyCountWithChan(request *DescribePropertyCountRequest) (<-chan *DescribePropertyCountResponse, <-chan error) {
	responseChan := make(chan *DescribePropertyCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertyCount(request)
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

// DescribePropertyCountWithCallback invokes the sas.DescribePropertyCount API asynchronously
func (client *Client) DescribePropertyCountWithCallback(request *DescribePropertyCountRequest, callback func(response *DescribePropertyCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertyCountResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertyCount(request)
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

// DescribePropertyCountRequest is the request struct for api DescribePropertyCount
type DescribePropertyCountRequest struct {
	*requests.RpcRequest
	Type     string `position:"Query" name:"Type"`
	SourceIp string `position:"Query" name:"SourceIp"`
	UuidList string `position:"Query" name:"UuidList"`
}

// DescribePropertyCountResponse is the response struct for api DescribePropertyCount
type DescribePropertyCountResponse struct {
	*responses.BaseResponse
	Process   int    `json:"Process" xml:"Process"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	User      int    `json:"User" xml:"User"`
	Software  int    `json:"Software" xml:"Software"`
	Cron      int    `json:"Cron" xml:"Cron"`
	Port      int    `json:"Port" xml:"Port"`
	Sca       int    `json:"Sca" xml:"Sca"`
	Web       int    `json:"Web" xml:"Web"`
	Database  int    `json:"Database" xml:"Database"`
	Lkm       int    `json:"Lkm" xml:"Lkm"`
	Autorun   int    `json:"Autorun" xml:"Autorun"`
	Webserver int    `json:"Webserver" xml:"Webserver"`
}

// CreateDescribePropertyCountRequest creates a request to invoke DescribePropertyCount API
func CreateDescribePropertyCountRequest() (request *DescribePropertyCountRequest) {
	request = &DescribePropertyCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertyCount", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePropertyCountResponse creates a response to parse from DescribePropertyCount response
func CreateDescribePropertyCountResponse() (response *DescribePropertyCountResponse) {
	response = &DescribePropertyCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
