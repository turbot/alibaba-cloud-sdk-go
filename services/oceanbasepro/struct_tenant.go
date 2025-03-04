package oceanbasepro

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

// Tenant is a nested struct in oceanbasepro response
type Tenant struct {
	TenantId                     string                  `json:"TenantId" xml:"TenantId"`
	TenantName                   string                  `json:"TenantName" xml:"TenantName"`
	TenantMode                   string                  `json:"TenantMode" xml:"TenantMode"`
	VpcId                        string                  `json:"VpcId" xml:"VpcId"`
	Status                       string                  `json:"Status" xml:"Status"`
	EnableInternetAddressService bool                    `json:"EnableInternetAddressService" xml:"EnableInternetAddressService"`
	PrimaryZone                  string                  `json:"PrimaryZone" xml:"PrimaryZone"`
	DeployType                   string                  `json:"DeployType" xml:"DeployType"`
	DeployMode                   string                  `json:"DeployMode" xml:"DeployMode"`
	Description                  string                  `json:"Description" xml:"Description"`
	CreateTime                   string                  `json:"CreateTime" xml:"CreateTime"`
	ClogServiceStatus            string                  `json:"ClogServiceStatus" xml:"ClogServiceStatus"`
	EnableClogService            bool                    `json:"EnableClogService" xml:"EnableClogService"`
	Charset                      string                  `json:"Charset" xml:"Charset"`
	Collation                    string                  `json:"Collation" xml:"Collation"`
	PrimaryZoneDeployType        string                  `json:"PrimaryZoneDeployType" xml:"PrimaryZoneDeployType"`
	MasterIntranetAddressZone    string                  `json:"MasterIntranetAddressZone" xml:"MasterIntranetAddressZone"`
	PayType                      string                  `json:"PayType" xml:"PayType"`
	InstanceType                 string                  `json:"InstanceType" xml:"InstanceType"`
	Series                       string                  `json:"Series" xml:"Series"`
	DiskType                     string                  `json:"DiskType" xml:"DiskType"`
	EnableReadWriteSplit         bool                    `json:"EnableReadWriteSplit" xml:"EnableReadWriteSplit"`
	AvailableZones               []string                `json:"AvailableZones" xml:"AvailableZones"`
	TenantResource               TenantResource          `json:"TenantResource" xml:"TenantResource"`
	TenantConnections            []TenantConnectionsItem `json:"TenantConnections" xml:"TenantConnections"`
	TenantZones                  []TenantZonesItem       `json:"TenantZones" xml:"TenantZones"`
}
