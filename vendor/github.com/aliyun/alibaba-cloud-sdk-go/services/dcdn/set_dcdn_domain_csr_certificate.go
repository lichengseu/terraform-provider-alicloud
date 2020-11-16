package dcdn

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

// SetDcdnDomainCSRCertificate invokes the dcdn.SetDcdnDomainCSRCertificate API synchronously
func (client *Client) SetDcdnDomainCSRCertificate(request *SetDcdnDomainCSRCertificateRequest) (response *SetDcdnDomainCSRCertificateResponse, err error) {
	response = CreateSetDcdnDomainCSRCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetDcdnDomainCSRCertificateWithChan invokes the dcdn.SetDcdnDomainCSRCertificate API asynchronously
func (client *Client) SetDcdnDomainCSRCertificateWithChan(request *SetDcdnDomainCSRCertificateRequest) (<-chan *SetDcdnDomainCSRCertificateResponse, <-chan error) {
	responseChan := make(chan *SetDcdnDomainCSRCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDcdnDomainCSRCertificate(request)
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

// SetDcdnDomainCSRCertificateWithCallback invokes the dcdn.SetDcdnDomainCSRCertificate API asynchronously
func (client *Client) SetDcdnDomainCSRCertificateWithCallback(request *SetDcdnDomainCSRCertificateRequest, callback func(response *SetDcdnDomainCSRCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDcdnDomainCSRCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetDcdnDomainCSRCertificate(request)
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

// SetDcdnDomainCSRCertificateRequest is the request struct for api SetDcdnDomainCSRCertificate
type SetDcdnDomainCSRCertificateRequest struct {
	*requests.RpcRequest
	ServerCertificate string           `position:"Query" name:"ServerCertificate"`
	DomainName        string           `position:"Query" name:"DomainName"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
}

// SetDcdnDomainCSRCertificateResponse is the response struct for api SetDcdnDomainCSRCertificate
type SetDcdnDomainCSRCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDcdnDomainCSRCertificateRequest creates a request to invoke SetDcdnDomainCSRCertificate API
func CreateSetDcdnDomainCSRCertificateRequest() (request *SetDcdnDomainCSRCertificateRequest) {
	request = &SetDcdnDomainCSRCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnDomainCSRCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDcdnDomainCSRCertificateResponse creates a response to parse from SetDcdnDomainCSRCertificate response
func CreateSetDcdnDomainCSRCertificateResponse() (response *SetDcdnDomainCSRCertificateResponse) {
	response = &SetDcdnDomainCSRCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}