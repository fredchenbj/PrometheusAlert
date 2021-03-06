package qualitycheck

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

// GetNextResultToReview invokes the qualitycheck.GetNextResultToReview API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getnextresulttoreview.html
func (client *Client) GetNextResultToReview(request *GetNextResultToReviewRequest) (response *GetNextResultToReviewResponse, err error) {
	response = CreateGetNextResultToReviewResponse()
	err = client.DoAction(request, response)
	return
}

// GetNextResultToReviewWithChan invokes the qualitycheck.GetNextResultToReview API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getnextresulttoreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNextResultToReviewWithChan(request *GetNextResultToReviewRequest) (<-chan *GetNextResultToReviewResponse, <-chan error) {
	responseChan := make(chan *GetNextResultToReviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNextResultToReview(request)
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

// GetNextResultToReviewWithCallback invokes the qualitycheck.GetNextResultToReview API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getnextresulttoreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNextResultToReviewWithCallback(request *GetNextResultToReviewRequest, callback func(response *GetNextResultToReviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNextResultToReviewResponse
		var err error
		defer close(result)
		response, err = client.GetNextResultToReview(request)
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

// GetNextResultToReviewRequest is the request struct for api GetNextResultToReview
type GetNextResultToReviewRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetNextResultToReviewResponse is the response struct for api GetNextResultToReview
type GetNextResultToReviewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetNextResultToReviewRequest creates a request to invoke GetNextResultToReview API
func CreateGetNextResultToReviewRequest() (request *GetNextResultToReviewRequest) {
	request = &GetNextResultToReviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetNextResultToReview", "", "")
	return
}

// CreateGetNextResultToReviewResponse creates a response to parse from GetNextResultToReview response
func CreateGetNextResultToReviewResponse() (response *GetNextResultToReviewResponse) {
	response = &GetNextResultToReviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
