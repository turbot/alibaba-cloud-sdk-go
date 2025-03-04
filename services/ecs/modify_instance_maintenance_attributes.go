package ecs

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

// ModifyInstanceMaintenanceAttributes invokes the ecs.ModifyInstanceMaintenanceAttributes API synchronously
func (client *Client) ModifyInstanceMaintenanceAttributes(request *ModifyInstanceMaintenanceAttributesRequest) (response *ModifyInstanceMaintenanceAttributesResponse, err error) {
	response = CreateModifyInstanceMaintenanceAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceMaintenanceAttributesWithChan invokes the ecs.ModifyInstanceMaintenanceAttributes API asynchronously
func (client *Client) ModifyInstanceMaintenanceAttributesWithChan(request *ModifyInstanceMaintenanceAttributesRequest) (<-chan *ModifyInstanceMaintenanceAttributesResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceMaintenanceAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceMaintenanceAttributes(request)
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

// ModifyInstanceMaintenanceAttributesWithCallback invokes the ecs.ModifyInstanceMaintenanceAttributes API asynchronously
func (client *Client) ModifyInstanceMaintenanceAttributesWithCallback(request *ModifyInstanceMaintenanceAttributesRequest, callback func(response *ModifyInstanceMaintenanceAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceMaintenanceAttributesResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceMaintenanceAttributes(request)
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

// ModifyInstanceMaintenanceAttributesRequest is the request struct for api ModifyInstanceMaintenanceAttributes
type ModifyInstanceMaintenanceAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                                        `position:"Query" name:"ResourceOwnerId"`
	MaintenanceWindow    *[]ModifyInstanceMaintenanceAttributesMaintenanceWindow `position:"Query" name:"MaintenanceWindow"  type:"Repeated"`
	ActionOnMaintenance  string                                                  `position:"Query" name:"ActionOnMaintenance"`
	ResourceOwnerAccount string                                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                                        `position:"Query" name:"OwnerId"`
	NotifyOnMaintenance  requests.Boolean                                        `position:"Query" name:"NotifyOnMaintenance"`
	InstanceId           *[]string                                               `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// ModifyInstanceMaintenanceAttributesMaintenanceWindow is a repeated param struct in ModifyInstanceMaintenanceAttributesRequest
type ModifyInstanceMaintenanceAttributesMaintenanceWindow struct {
	EndTime   string `name:"EndTime"`
	StartTime string `name:"StartTime"`
}

// ModifyInstanceMaintenanceAttributesResponse is the response struct for api ModifyInstanceMaintenanceAttributes
type ModifyInstanceMaintenanceAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceMaintenanceAttributesRequest creates a request to invoke ModifyInstanceMaintenanceAttributes API
func CreateModifyInstanceMaintenanceAttributesRequest() (request *ModifyInstanceMaintenanceAttributesRequest) {
	request = &ModifyInstanceMaintenanceAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceMaintenanceAttributes", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceMaintenanceAttributesResponse creates a response to parse from ModifyInstanceMaintenanceAttributes response
func CreateModifyInstanceMaintenanceAttributesResponse() (response *ModifyInstanceMaintenanceAttributesResponse) {
	response = &ModifyInstanceMaintenanceAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
