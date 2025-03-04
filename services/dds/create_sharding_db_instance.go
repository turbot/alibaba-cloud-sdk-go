package dds

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

// CreateShardingDBInstance invokes the dds.CreateShardingDBInstance API synchronously
func (client *Client) CreateShardingDBInstance(request *CreateShardingDBInstanceRequest) (response *CreateShardingDBInstanceResponse, err error) {
	response = CreateCreateShardingDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateShardingDBInstanceWithChan invokes the dds.CreateShardingDBInstance API asynchronously
func (client *Client) CreateShardingDBInstanceWithChan(request *CreateShardingDBInstanceRequest) (<-chan *CreateShardingDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateShardingDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateShardingDBInstance(request)
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

// CreateShardingDBInstanceWithCallback invokes the dds.CreateShardingDBInstance API asynchronously
func (client *Client) CreateShardingDBInstanceWithCallback(request *CreateShardingDBInstanceRequest, callback func(response *CreateShardingDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateShardingDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateShardingDBInstance(request)
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

// CreateShardingDBInstanceRequest is the request struct for api CreateShardingDBInstance
type CreateShardingDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	SecondaryZoneId       string                                  `position:"Query" name:"SecondaryZoneId"`
	EngineVersion         string                                  `position:"Query" name:"EngineVersion"`
	NetworkType           string                                  `position:"Query" name:"NetworkType"`
	ReplicaSet            *[]CreateShardingDBInstanceReplicaSet   `position:"Query" name:"ReplicaSet"  type:"Repeated"`
	StorageType           string                                  `position:"Query" name:"StorageType"`
	ResourceGroupId       string                                  `position:"Query" name:"ResourceGroupId"`
	SecurityToken         string                                  `position:"Query" name:"SecurityToken"`
	DBInstanceDescription string                                  `position:"Query" name:"DBInstanceDescription"`
	Period                requests.Integer                        `position:"Query" name:"Period"`
	ConfigServer          *[]CreateShardingDBInstanceConfigServer `position:"Query" name:"ConfigServer"  type:"Repeated"`
	OwnerId               requests.Integer                        `position:"Query" name:"OwnerId"`
	SecurityIPList        string                                  `position:"Query" name:"SecurityIPList"`
	VSwitchId             string                                  `position:"Query" name:"VSwitchId"`
	Mongos                *[]CreateShardingDBInstanceMongos       `position:"Query" name:"Mongos"  type:"Repeated"`
	AutoRenew             string                                  `position:"Query" name:"AutoRenew"`
	ZoneId                string                                  `position:"Query" name:"ZoneId"`
	ClientToken           string                                  `position:"Query" name:"ClientToken"`
	StorageEngine         string                                  `position:"Query" name:"StorageEngine"`
	Engine                string                                  `position:"Query" name:"Engine"`
	HiddenZoneId          string                                  `position:"Query" name:"HiddenZoneId"`
	RestoreTime           string                                  `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string                                  `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId       string                                  `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount          string                                  `position:"Query" name:"OwnerAccount"`
	AccountPassword       string                                  `position:"Query" name:"AccountPassword"`
	VpcId                 string                                  `position:"Query" name:"VpcId"`
	ProtocolType          string                                  `position:"Query" name:"ProtocolType"`
	ChargeType            string                                  `position:"Query" name:"ChargeType"`
}

// CreateShardingDBInstanceReplicaSet is a repeated param struct in CreateShardingDBInstanceRequest
type CreateShardingDBInstanceReplicaSet struct {
	ReadonlyReplicas string `name:"ReadonlyReplicas"`
	Storage          string `name:"Storage"`
	Class            string `name:"Class"`
}

// CreateShardingDBInstanceConfigServer is a repeated param struct in CreateShardingDBInstanceRequest
type CreateShardingDBInstanceConfigServer struct {
	Storage string `name:"Storage"`
	Class   string `name:"Class"`
}

// CreateShardingDBInstanceMongos is a repeated param struct in CreateShardingDBInstanceRequest
type CreateShardingDBInstanceMongos struct {
	Class string `name:"Class"`
}

// CreateShardingDBInstanceResponse is the response struct for api CreateShardingDBInstance
type CreateShardingDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateShardingDBInstanceRequest creates a request to invoke CreateShardingDBInstance API
func CreateCreateShardingDBInstanceRequest() (request *CreateShardingDBInstanceRequest) {
	request = &CreateShardingDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CreateShardingDBInstance", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateShardingDBInstanceResponse creates a response to parse from CreateShardingDBInstance response
func CreateCreateShardingDBInstanceResponse() (response *CreateShardingDBInstanceResponse) {
	response = &CreateShardingDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
