package rds

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

// ModifyBackupPolicy invokes the rds.ModifyBackupPolicy API synchronously
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (response *ModifyBackupPolicyResponse, err error) {
	response = CreateModifyBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupPolicyWithChan invokes the rds.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithChan(request *ModifyBackupPolicyRequest) (<-chan *ModifyBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyWithCallback invokes the rds.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithCallback(request *ModifyBackupPolicyRequest, callback func(response *ModifyBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyRequest is the request struct for api ModifyBackupPolicy
type ModifyBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LocalLogRetentionHours        string           `position:"Query" name:"LocalLogRetentionHours"`
	LogBackupFrequency            string           `position:"Query" name:"LogBackupFrequency"`
	ArchiveBackupKeepCount        requests.Integer `position:"Query" name:"ArchiveBackupKeepCount"`
	BackupLog                     string           `position:"Query" name:"BackupLog"`
	BackupInterval                string           `position:"Query" name:"BackupInterval"`
	DuplicationContent            string           `position:"Query" name:"DuplicationContent"`
	HighSpaceUsageProtection      string           `position:"Query" name:"HighSpaceUsageProtection"`
	LogBackupLocalRetentionNumber requests.Integer `position:"Query" name:"LogBackupLocalRetentionNumber"`
	DBInstanceId                  string           `position:"Query" name:"DBInstanceId"`
	EnableBackupLog               string           `position:"Query" name:"EnableBackupLog"`
	BackupPolicyMode              string           `position:"Query" name:"BackupPolicyMode"`
	PreferredBackupPeriod         string           `position:"Query" name:"PreferredBackupPeriod"`
	EnableIncrementDataBackup     requests.Boolean `position:"Query" name:"EnableIncrementDataBackup"`
	ReleasedKeepPolicy            string           `position:"Query" name:"ReleasedKeepPolicy"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	CompressType                  string           `position:"Query" name:"CompressType"`
	LocalLogRetentionSpace        string           `position:"Query" name:"LocalLogRetentionSpace"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	ArchiveBackupKeepPolicy       string           `position:"Query" name:"ArchiveBackupKeepPolicy"`
	Duplication                   string           `position:"Query" name:"Duplication"`
	PreferredBackupTime           string           `position:"Query" name:"PreferredBackupTime"`
	BackupRetentionPeriod         string           `position:"Query" name:"BackupRetentionPeriod"`
	BackupMethod                  string           `position:"Query" name:"BackupMethod"`
	DuplicationLocation           string           `position:"Query" name:"DuplicationLocation"`
	ArchiveBackupRetentionPeriod  string           `position:"Query" name:"ArchiveBackupRetentionPeriod"`
	Category                      string           `position:"Query" name:"Category"`
	LogBackupRetentionPeriod      string           `position:"Query" name:"LogBackupRetentionPeriod"`
}

// ModifyBackupPolicyResponse is the response struct for api ModifyBackupPolicy
type ModifyBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId                     string `json:"RequestId" xml:"RequestId"`
	CompressType                  string `json:"CompressType" xml:"CompressType"`
	LocalLogRetentionSpace        string `json:"LocalLogRetentionSpace" xml:"LocalLogRetentionSpace"`
	LogBackupLocalRetentionNumber int    `json:"LogBackupLocalRetentionNumber" xml:"LogBackupLocalRetentionNumber"`
	DBInstanceID                  string `json:"DBInstanceID" xml:"DBInstanceID"`
	EnableBackupLog               string `json:"EnableBackupLog" xml:"EnableBackupLog"`
	LocalLogRetentionHours        int    `json:"LocalLogRetentionHours" xml:"LocalLogRetentionHours"`
	HighSpaceUsageProtection      string `json:"HighSpaceUsageProtection" xml:"HighSpaceUsageProtection"`
}

// CreateModifyBackupPolicyRequest creates a request to invoke ModifyBackupPolicy API
func CreateModifyBackupPolicyRequest() (request *ModifyBackupPolicyRequest) {
	request = &ModifyBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyBackupPolicy", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackupPolicyResponse creates a response to parse from ModifyBackupPolicy response
func CreateModifyBackupPolicyResponse() (response *ModifyBackupPolicyResponse) {
	response = &ModifyBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
