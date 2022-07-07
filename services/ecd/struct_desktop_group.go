package ecd

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

// DesktopGroup is a nested struct in ecd response
type DesktopGroup struct {
	CreateTime              string  `json:"CreateTime" xml:"CreateTime"`
	PayType                 string  `json:"PayType" xml:"PayType"`
	PolicyGroupName         string  `json:"PolicyGroupName" xml:"PolicyGroupName"`
	Creator                 string  `json:"Creator" xml:"Creator"`
	MaxDesktopsCount        int     `json:"MaxDesktopsCount" xml:"MaxDesktopsCount"`
	SystemDiskSize          int     `json:"SystemDiskSize" xml:"SystemDiskSize"`
	PolicyGroupId           string  `json:"PolicyGroupId" xml:"PolicyGroupId"`
	OwnBundleId             string  `json:"OwnBundleId" xml:"OwnBundleId"`
	GpuCount                float64 `json:"GpuCount" xml:"GpuCount"`
	Memory                  int64   `json:"Memory" xml:"Memory"`
	GpuSpec                 string  `json:"GpuSpec" xml:"GpuSpec"`
	DirectoryId             string  `json:"DirectoryId" xml:"DirectoryId"`
	OwnBundleName           string  `json:"OwnBundleName" xml:"OwnBundleName"`
	DataDiskCategory        string  `json:"DataDiskCategory" xml:"DataDiskCategory"`
	DesktopGroupName        string  `json:"DesktopGroupName" xml:"DesktopGroupName"`
	SystemDiskCategory      string  `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	OfficeSiteId            string  `json:"OfficeSiteId" xml:"OfficeSiteId"`
	KeepDuration            int64   `json:"KeepDuration" xml:"KeepDuration"`
	MinDesktopsCount        int     `json:"MinDesktopsCount" xml:"MinDesktopsCount"`
	EndUserCount            int     `json:"EndUserCount" xml:"EndUserCount"`
	DataDiskSize            string  `json:"DataDiskSize" xml:"DataDiskSize"`
	DesktopGroupId          string  `json:"DesktopGroupId" xml:"DesktopGroupId"`
	OfficeSiteName          string  `json:"OfficeSiteName" xml:"OfficeSiteName"`
	DirectoryType           string  `json:"DirectoryType" xml:"DirectoryType"`
	Cpu                     int     `json:"Cpu" xml:"Cpu"`
	ExpiredTime             string  `json:"ExpiredTime" xml:"ExpiredTime"`
	Comments                string  `json:"Comments" xml:"Comments"`
	OfficeSiteType          string  `json:"OfficeSiteType" xml:"OfficeSiteType"`
	Status                  int     `json:"Status" xml:"Status"`
	ResetType               int64   `json:"ResetType" xml:"ResetType"`
	LoadPolicy              int64   `json:"LoadPolicy" xml:"LoadPolicy"`
	BindAmount              int64   `json:"BindAmount" xml:"BindAmount"`
	OwnType                 int64   `json:"OwnType" xml:"OwnType"`
	ImageId                 string  `json:"ImageId" xml:"ImageId"`
	VolumeEncryptionEnabled bool    `json:"VolumeEncryptionEnabled" xml:"VolumeEncryptionEnabled"`
	VolumeEncryptionKey     string  `json:"VolumeEncryptionKey" xml:"VolumeEncryptionKey"`
}
