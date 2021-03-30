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

// RegisterInstance invokes the dms_enterprise.RegisterInstance API synchronously
func (client *Client) RegisterInstance(request *RegisterInstanceRequest) (response *RegisterInstanceResponse, err error) {
	response = CreateRegisterInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterInstanceWithChan invokes the dms_enterprise.RegisterInstance API asynchronously
func (client *Client) RegisterInstanceWithChan(request *RegisterInstanceRequest) (<-chan *RegisterInstanceResponse, <-chan error) {
	responseChan := make(chan *RegisterInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterInstance(request)
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

// RegisterInstanceWithCallback invokes the dms_enterprise.RegisterInstance API asynchronously
func (client *Client) RegisterInstanceWithCallback(request *RegisterInstanceRequest, callback func(response *RegisterInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterInstanceResponse
		var err error
		defer close(result)
		response, err = client.RegisterInstance(request)
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

// RegisterInstanceRequest is the request struct for api RegisterInstance
type RegisterInstanceRequest struct {
	*requests.RpcRequest
	EcsRegion        string           `position:"Query" name:"EcsRegion"`
	DdlOnline        requests.Integer `position:"Query" name:"DdlOnline"`
	UseDsql          requests.Integer `position:"Query" name:"UseDsql"`
	NetworkType      string           `position:"Query" name:"NetworkType"`
	Tid              requests.Integer `position:"Query" name:"Tid"`
	Sid              string           `position:"Query" name:"Sid"`
	DataLinkName     string           `position:"Query" name:"DataLinkName"`
	InstanceSource   string           `position:"Query" name:"InstanceSource"`
	EnvType          string           `position:"Query" name:"EnvType"`
	Host             string           `position:"Query" name:"Host"`
	InstanceType     string           `position:"Query" name:"InstanceType"`
	QueryTimeout     requests.Integer `position:"Query" name:"QueryTimeout"`
	EcsInstanceId    string           `position:"Query" name:"EcsInstanceId"`
	ExportTimeout    requests.Integer `position:"Query" name:"ExportTimeout"`
	DatabasePassword string           `position:"Query" name:"DatabasePassword"`
	InstanceAlias    string           `position:"Query" name:"InstanceAlias"`
	DatabaseUser     string           `position:"Query" name:"DatabaseUser"`
	Port             requests.Integer `position:"Query" name:"Port"`
	VpcId            string           `position:"Query" name:"VpcId"`
	DbaUid           requests.Integer `position:"Query" name:"DbaUid"`
	SkipTest         requests.Boolean `position:"Query" name:"SkipTest"`
	SafeRule         string           `position:"Query" name:"SafeRule"`
}

// RegisterInstanceResponse is the response struct for api RegisterInstance
type RegisterInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateRegisterInstanceRequest creates a request to invoke RegisterInstance API
func CreateRegisterInstanceRequest() (request *RegisterInstanceRequest) {
	request = &RegisterInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "RegisterInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateRegisterInstanceResponse creates a response to parse from RegisterInstance response
func CreateRegisterInstanceResponse() (response *RegisterInstanceResponse) {
	response = &RegisterInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
