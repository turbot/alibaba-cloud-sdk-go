package das

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

// GetGeneralAbnormalEventsCount invokes the das.GetGeneralAbnormalEventsCount API synchronously
func (client *Client) GetGeneralAbnormalEventsCount(request *GetGeneralAbnormalEventsCountRequest) (response *GetGeneralAbnormalEventsCountResponse, err error) {
	response = CreateGetGeneralAbnormalEventsCountResponse()
	err = client.DoAction(request, response)
	return
}

// GetGeneralAbnormalEventsCountWithChan invokes the das.GetGeneralAbnormalEventsCount API asynchronously
func (client *Client) GetGeneralAbnormalEventsCountWithChan(request *GetGeneralAbnormalEventsCountRequest) (<-chan *GetGeneralAbnormalEventsCountResponse, <-chan error) {
	responseChan := make(chan *GetGeneralAbnormalEventsCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGeneralAbnormalEventsCount(request)
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

// GetGeneralAbnormalEventsCountWithCallback invokes the das.GetGeneralAbnormalEventsCount API asynchronously
func (client *Client) GetGeneralAbnormalEventsCountWithCallback(request *GetGeneralAbnormalEventsCountRequest, callback func(response *GetGeneralAbnormalEventsCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGeneralAbnormalEventsCountResponse
		var err error
		defer close(result)
		response, err = client.GetGeneralAbnormalEventsCount(request)
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

// GetGeneralAbnormalEventsCountRequest is the request struct for api GetGeneralAbnormalEventsCount
type GetGeneralAbnormalEventsCountRequest struct {
	*requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
}

// GetGeneralAbnormalEventsCountResponse is the response struct for api GetGeneralAbnormalEventsCount
type GetGeneralAbnormalEventsCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Synchro   string `json:"Synchro" xml:"Synchro"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateGetGeneralAbnormalEventsCountRequest creates a request to invoke GetGeneralAbnormalEventsCount API
func CreateGetGeneralAbnormalEventsCountRequest() (request *GetGeneralAbnormalEventsCountRequest) {
	request = &GetGeneralAbnormalEventsCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetGeneralAbnormalEventsCount", "", "")
	request.Method = requests.POST
	return
}

// CreateGetGeneralAbnormalEventsCountResponse creates a response to parse from GetGeneralAbnormalEventsCount response
func CreateGetGeneralAbnormalEventsCountResponse() (response *GetGeneralAbnormalEventsCountResponse) {
	response = &GetGeneralAbnormalEventsCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
