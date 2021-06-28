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

// OrderBaseInfo is a nested struct in dms_enterprise response
type OrderBaseInfo struct {
	Comment             string              `json:"Comment" xml:"Comment"`
	CreateTime          string              `json:"CreateTime" xml:"CreateTime"`
	Committer           string              `json:"Committer" xml:"Committer"`
	WorkflowInstanceId  int64               `json:"WorkflowInstanceId" xml:"WorkflowInstanceId"`
	CommitterId         int64               `json:"CommitterId" xml:"CommitterId"`
	LastModifyTime      string              `json:"LastModifyTime" xml:"LastModifyTime"`
	StatusCode          string              `json:"StatusCode" xml:"StatusCode"`
	WorkflowStatusDesc  string              `json:"WorkflowStatusDesc" xml:"WorkflowStatusDesc"`
	StatusDesc          string              `json:"StatusDesc" xml:"StatusDesc"`
	PluginType          string              `json:"PluginType" xml:"PluginType"`
	OrderId             int64               `json:"OrderId" xml:"OrderId"`
	RelatedUserNickList RelatedUserNickList `json:"RelatedUserNickList" xml:"RelatedUserNickList"`
	RelatedUserList     RelatedUserList     `json:"RelatedUserList" xml:"RelatedUserList"`
}
