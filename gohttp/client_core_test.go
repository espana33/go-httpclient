package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T){
	//Initialization
	client:= httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type","application/json")
	commonHeaders.Set("User-Agent","cool-http-client")
	client.builder.headers = commonHeaders

	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id","ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation
	if len(finalHeaders) != 3 {
		t.Error("Its expected 3 headers for this test")
	}
}

func TestGetRequestBodyNilBody(t *testing.T){
	//Initialization
	client:= httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T){
	
		//Execution
		body, err := client.getRequestBody("",nil)
		
		//Validation
		if body != nil || err != nil{
			t.Error("Its expected a nil return")
		}
	})

	t.Run("JsonResponse", func(t *testing.T){
	
		//Execution
		requestBody := []string{"one","two"}
		body, err := client.getRequestBody("application/json",requestBody)
		
		//Validation
		if err != nil{
			t.Error("No error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]`{
			t.Error("Invalid Json body response")
		}		
	})

}