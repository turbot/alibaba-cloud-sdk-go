package imm

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

// FacesItemInSearchImagesByTagNames is a nested struct in imm response
type FacesItemInSearchImagesByTagNames struct {
	Age               int            `json:"Age" xml:"Age"`
	GenderConfidence  float64        `json:"GenderConfidence" xml:"GenderConfidence"`
	Attractive        float64        `json:"Attractive" xml:"Attractive"`
	Gender            string         `json:"Gender" xml:"Gender"`
	FaceConfidence    float64        `json:"FaceConfidence" xml:"FaceConfidence"`
	Emotion           string         `json:"Emotion" xml:"Emotion"`
	FaceId            string         `json:"FaceId" xml:"FaceId"`
	EmotionConfidence float64        `json:"EmotionConfidence" xml:"EmotionConfidence"`
	EmotionDetails    EmotionDetails `json:"EmotionDetails" xml:"EmotionDetails"`
	FaceAttributes    FaceAttributes `json:"FaceAttributes" xml:"FaceAttributes"`
}
