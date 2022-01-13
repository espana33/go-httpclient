package main

import (
	"fmt"

	"github.com/espana33/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

type User struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getGithubClient() gohttp.Client {
	client:=gohttp.NewBuilder().DisableTimeouts(true).SetMaxIdleConnections(5).Build()
	//builder.SetConnectionTimeout(2 * time.Second)
	//builder.SetResponseTimeout(4 * time.Second)

	//commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer ABC-123")
	//builder.SetHeaders(commonHeaders)
	return client
}

func main(){
	getUrls()
}

func getUrls(){
	response, err := githubHttpClient.Get("https://api.github.com",nil)

	if err != nil{
		panic(err)
	}
	
	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())


	// Using our custome response: 
	/*var user User
	if err := response.UnmarshalJson(&user); err != nil{
		panic(err)
	}
	fmt.Println(user.FirstName)	*/
	

	//Using default http.Response

	/*response.Body.Close()
	
	fmt.Println(response.StatusCode)
	
	bytes2, err:= ioutil.ReadAll(response.Body)


	if err != nil{
		panic(err)
	}
	var user2 User
	if err := json.Unmarshal(bytes2,&user2); err != nil{
		panic(err)
	}
	fmt.Println(user2.FirstName)	*/


}

func createUser(user User){
	response, err := githubHttpClient.Post("https://api.github.com",nil,user)

	if err != nil{
		panic(err)
	}

	fmt.Println(response.StatusCode)


	fmt.Println(response.String())	
}