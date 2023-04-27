# twilio-go

[![Tests](https://github.com/twilio/twilio-go/actions/workflows/test-and-deploy.yml/badge.svg)](https://github.com/twilio/twilio-go/actions/workflows/test-and-deploy.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/twilio/twilio-go)](https://pkg.go.dev/github.com/twilio/twilio-go)
[![Release](https://img.shields.io/github/release/twilio/twilio-go.svg)](https://github.com/twilio/twilio-go/releases/latest)
[![Learn OSS Contribution in TwilioQuest](https://img.shields.io/static/v1?label=TwilioQuest&message=Learn%20to%20contribute%21&color=F22F46&labelColor=1f243c&style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAASFBMVEUAAAAZGRkcHBwjIyMoKCgAAABgYGBoaGiAgICMjIyzs7PJycnMzMzNzc3UoBfd3d3m5ubqrhfrMEDu7u739/f4vSb/3AD///9tbdyEAAAABXRSTlMAAAAAAMJrBrEAAAKoSURBVHgB7ZrRcuI6EESdyxXGYoNFvMD//+l2bSszRgyUYpFAsXOeiJGmj4NkuWx1Qeh+Ekl9DgEXOBwOx+Px5xyQhDykfgq4wG63MxxaR4ddIkg6Ul3g84vCIcjPBA5gmUMeXESrlukuoK33+33uID8TWeLAdOWsKpJYzwVMB7bOzYSGOciyUlXSn0/ABXTosJ1M1SbypZ4O4MbZuIDMU02PMbauhhHMHXbmebmALIiEbbbbbUrpF1gwE9kFfRNAJaP+FQEXCCTGyJ4ngDrjOFo3jEL5JdqjF/pueR4cCeCGgAtwmuRS6gDwaRiGvu+DMFwSBLTE3+jF8JyuV1okPZ+AC4hDFhCHyHQjdjPHUKFDlHSJkHQXMB3KpSwXNGJPcwwTdZiXlRN0gSp0zpWxNtM0beYE0nRH6QIbO7rawwXaBYz0j78gxjokDuv12gVeUuBD0MDi0OQCLvDaAho4juP1Q/jkAncXqIcCfd+7gAu4QLMACCLxpRsSuQh0igu0C9Svhi7weAGZg50L3IE3cai4IfkNZAC8dfdhsUD3CgKBVC9JE5ABAFzg4QL/taYPAAWrHdYcgfLaIgAXWJ7OV38n1LEF8tt2TH29E+QAoDoO5Ve/LtCQDmKM9kPbvCEBApK+IXzbcSJ0cIGF6e8gpcRhUDogWZ8JnaWjPXc/fNnBBUKRngiHgTUSivSzDRDgHZQOLvBQgf8rRt+VdBUUhwkU6VpJ+xcOwQUqZr+mR0kvBUgv6cB4+37hQAkXqE8PwGisGhJtN4xAHMzrsgvI7rccXqSvKh6jltGlrOHA3Xk1At3LC4QiPdX9/0ndHpGVvTjR4bZA1ypAKgVcwE5vx74ulwIugDt8e/X7JgfkucBMIAr26ndnB4UCLnDOqvteQsHlgX9N4A+c4cW3DXSPbwAAAABJRU5ErkJggg==)](https://twil.io/learn-open-source)

All the code [here](./rest) was generated by [twilio-oai-generator](https://github.com/twilio/twilio-oai-generator) by
leveraging [openapi-generator](https://github.com/OpenAPITools/openapi-generator)
and [twilio-oai](https://github.com/twilio/twilio-oai). If you find an issue with the generation or the OpenAPI specs,
please go ahead and open an issue or a PR against the relevant repositories.

## Documentation

The documentation for the Twilio API can be found [here][apidocs].

The Go library documentation can be found [here][libdocs].

### Supported Go Versions

This library supports the following Go implementations:

* Go 1.15
* Go 1.16
* Go 1.17
* Go 1.18
* Go 1.19

## Installation

To use twilio-go in your project initialize go modules then run:

```bash
go get github.com/twilio/twilio-go
```

## Getting Started

Getting started with the Twilio API couldn't be easier. Create a
`RestClient` and you're ready to go.

### API Credentials

The Twilio `RestClient` needs your Twilio credentials. We recommend storing them as environment variables, so that you don't have to worry about committing and accidentally posting them somewhere public. See http://twil.io/secure for more details on how to store environment variables.

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	// This will look for `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` variables inside the current environment to initialize the constructor
	// You can find your Account SID and Auth Token at twilio.com/console
	// `TWILIO_ACCOUNT_SID` should be in format "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	client := twilio.NewRestClient()
}
```

If you don't want to use environment variables, you can also pass the credentials directly to the constructor as below.

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}
```

### Using Subaccount

Subaccounts in Twilio are just accounts that are "owned" by your account. Twilio users can create subaccounts to help separate Twilio account usage into different buckets.

If you wish to make API calls with a Subaccount, you can do so by setting the `AccountSid` field in the `twilio.ClientParams`:

```go
package main

import "github.com/twilio/twilio-go"

func main() {
	// subaccountSid should also be in format "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	subaccountSid := os.Getenv("TWILIO_SUBACCOUNT_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: subaccountSid,
	})
}
```

### Specify a Region and/or Edge

```go
package main

import (
	"github.com/twilio/twilio-go"
)

func main() {
	client := twilio.NewRestClient()
	client.SetRegion("au1")
	client.SetEdge("sydney")
}
```

This will result in the `hostname` transforming from `api.twilio.com` to `api.sydney.au1.twilio.com`.

A Twilio client constructed without these parameters will also look for `TWILIO_REGION` and `TWILIO_EDGE` variables
inside the current environment.

### Buy a phone number

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	phoneNumber := "AVAILABLE_TWILIO_PHONE_NUMBER"

	client := twilio.NewRestClient()

	params := &twilioApi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Phone Number Status: " + *resp.Status)
	}
}
```

### Send a text message

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
```

### Make a call

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetUrl("http://twimlets.com/holdmusic?Bucket=com.twilio.music.ambient")

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Call Status: " + *resp.Status)
		fmt.Println("Call Sid: " + *resp.Sid)
		fmt.Println("Call Direction: " + *resp.Direction)
	}
}
```

### Create a Serverless Function

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/serverless/v1"
)

func main() {
	// serviceSid should be in format "ZSxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	serviceSid := os.Getenv("TWILIO_SERVICE_SID")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateFunctionParams{}
	params.SetFriendlyName("My Serverless func")

	resp, err := client.ServerlessV1.CreateFunction(serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*resp.Sid)
	}
}
```

### Create a Studio Flow

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/studio/v2"
)

func main() {
	var jsonStr = `{
		"description":"Twilio Studio flow service",
		"initial_state":"Trigger",
		"states":[
			{
				"properties":{
					"offset":{
						"y":0,
						"x":0
					}
				},
				"transitions":[

				],
				"name":"Trigger",
				"type":"trigger"
			}
		]
	}`

	var definition interface{}
	_ = json.Unmarshal([]byte(jsonStr), &definition)

	client := twilio.NewRestClient()
	params := &twilioApi.CreateFlowParams{
		Definition: &definition,
	}
	params.SetCommitMessage("commit")
	params.SetFriendlyName("Studio flow from Go")
	params.SetStatus("draft")

	resp, err := client.StudioV2.CreateFlow(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*resp.Sid)
	}
}
```

### Using Paging

This library also offers paging functionality. Collections such as calls and messages have `ListXxx` and `StreamXxx`
functions that page under the hood. With both list and stream, you can specify the number of records you want to
receive (limit) and the maximum size you want each page fetch to be (pageSize). The library will then handle the task
for you.

`List` eagerly fetches all records and returns them as a list, whereas `Stream` streams the records and lazily retrieves
the pages as you iterate over the collection. Also, `List` returns no records if any errors are encountered while paging,
whereas `Stream` returns all records up until encountering an error. You can also page manually using the `PageXxx`
function in each of the apis.

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.ListMessageParams{}
	params.SetFrom(from)
	params.SetPageSize(20)
	params.SetLimit(100)

	resp, _ := client.Api.ListMessage(params)
	for record := range resp {
		fmt.Println("Body: ", *resp[record].Body)
	}

	channel, _ := client.Api.StreamMessage(params)
	for record := range channel {
		fmt.Println("Body: ", *record.Body)
	}
}
```

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"net/url"
	"os"
)

func main() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.ListMessageParams{}
	params.SetFrom(from)
	params.SetPageSize(20)

	var pageToken string
	var pageNumber string
	resp, err = client.Api.PageMessage(params, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.NextPageUri)
		u, _ := url.Parse(resp.NextPageUri)
		q := u.Query()
		pageToken = q.Get("PageToken")
		pageNumber = q.Get("Page")
	}

	resp, err := client.Api.PageMessage(params, pageToken, pageNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		if resp != nil {
			fmt.Println(*resp.Messages[0].Body)
		}
	}
}
```

### Handling Exceptions

```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioclient "github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &twilioApi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)

	resp, err := client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		twilioError := err.(*twilioclient.TwilioRestError)
		fmt.Println(twilioError.Error())
	}
}
```

For more descriptive exception types, please see
the [Twilio documentation](https://www.twilio.com/docs/libraries/go/usage-guide#exceptions).

### Generating TwiML

To control phone calls, your application needs to output [TwiML](https://www.twilio.com/docs/voice/twiml).

Use the `twiml` package to easily create such responses.

```go
package main

import (
	"fmt"
	"github.com/twilio/twilio-go/twiml"
)

func main() {
	//Construct Verbs
	dial := &twiml.VoiceDial{}
	say := &twiml.VoiceSay{
		Message:            "Welcome to Twilio!",
		Voice:              "woman",
		Language:           "en-gb",
		OptionalAttributes: map[string]string{"input": "test"},
	}
	pause := &twiml.VoicePause{
		Length: "10",
	}
	//Construct Noun
	queue := &twiml.VoiceQueue{
		Url: "www.twilio.com",
	}
	//Adding Queue to Dial
	dial.InnerElements = []twiml.Element{queue}

	//Adding all Verbs to twiml.Voice
	verbList := []twiml.Element{dial, say, pause}
	twimlResult, err := twiml.Voice(verbList)
	if err == nil {
		fmt.Println(twimlResult)
	} else {
		fmt.Println(err)
	}
}
```
This will print the following:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Dial>
        <Queue url="www.twilio.com"/>
    </Dial>
    <Say voice="woman" language="en-gb" input="test">Welcome to Twilio!</Say>
    <Pause length="10"/>
</Response>
```
## Advanced Usage

### Using Request Validator

Validating GET/POST Requests are coming from Twilio:
```go
package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go/client"
)

func main() {
	// You can find your Auth Token at twilio.com/console
	// For this example: authToken := "12345"
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	
	requestValidator := client.NewRequestValidator(authToken)

	// Twilio's request URL
	url := "https://mycompany.com/myapp.php?foo=1&bar=2"
	
	// Post variables in Twilio's request
	params := map[string]string{
		"CallSid": "CA1234567890ABCDE",
		"Caller":  "+12349013030",
		"Digits":  "1234",
		"From":    "+12349013030",
		"To":      "+18005551212",
	}
	
	// X-Twilio-Signature header attached to the request
	signature := "0/KCTR6DLpKmkAf8muzZqo1nDgQ="
    
	// Validate GET request
	fmt.Println(requestValidator.Validate(url, params, signature))

	// Example of the POST request
	Body := []byte(`{"property": "value", "boolean": true}`)
	theUrl := "https://mycompany.com/myapp.php?bodySHA256=0a1ff7634d9ab3b95db5c9a2dfe9416e41502b283a80c7cf19632632f96e6620"
	theSignature := "y77kIzt2vzLz71DgmJGsen2scGs="
    
	// Validate POST request
	fmt.Println(requestValidator.ValidateBody(theUrl, Body, theSignature))
}
```

### Using Standalone Products

Don't want to import the top-level Twilio RestClient with access to the full suite of Twilio products? Use standalone
product services instead:

```go
package main

import (
	"github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	serverless "github.com/twilio/twilio-go/rest/serverless/v1"
	"os"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Create an instance of our default BaseClient implementation
	// You will need to provide your API credentials to the Client manually
	defaultClient := &client.Client{
		Credentials: client.NewCredentials(accountSid, authToken),
	}
	defaultClient.SetAccountSid(accountSid)

	coreApiService := twilioApi.NewApiServiceWithClient(defaultClient)
	serverlessApiService := serverless.NewApiServiceWithClient(defaultClient)
}
```

### Using a Custom Client

```go
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/client"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type MyClient struct {
	client.Client
}

func (c *MyClient) SendRequest(method string, rawURL string, data url.Values, headers map[string]interface{}) (*http.Response, error) {
	// Custom code to pre-process request here
	resp, err := c.Client.SendRequest(method, rawURL, data, headers)
	// Custom code to pre-process response here
	fmt.Println(resp.StatusCode)
	return resp, err
}

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	customClient := &MyClient{
		Client: client.Client{
			Credentials: client.NewCredentials(accountSid, authToken),
		},
	}
	customClient.SetAccountSid(accountSid)

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{Client: customClient})

	// You may also use custom clients with standalone product services
	twilioApiV2010 := twilioApi.NewApiServiceWithClient(customClient)
}
```

## Building Access Tokens
This library supports [access token](https://www.twilio.com/docs/iam/access-tokens) generation for use in the Twilio Client SDKs.

Here's how you would generate a token for the Voice SDK:
```go
package main

import (
	"os"
	"github.com/twilio/twilio-go/client/jwt"
)

accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
applicationSid := os.Getenv("TWILIO_TWIML_APP_SID")
apiKey := os.Getenv("TWILIO_API_KEY")
apiSecret := os.Getenv("TWILIO_API_SECRET")
identity := "fake123"

params := jwt.AccessTokenParams{
	AccountSid: accountSid,
	SigningKeySid: apiKey,
	Secret: apiSecret,
	Identity: identity,
}

jwtToken := jwt.CreateAccessToken(params)
voiceGrant := &jwt.VoiceGrant{
	Incoming: jwt.Incoming{Allow: true},
	Outgoing: jwt.Outgoing{
		ApplicationSid: applicationSid,
	},
}

jwtToken.AddGrant(voiceGrant)
token, err := jwtToken.ToJwt()
```

Creating Capability Token for TaskRouter v1:
```go
package main

import (
	"os"
	"github.com/twilio/twilio-go/client/jwt/taskrouter"
)

AccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
AuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
WorkspaceSid := os.Getenv("TWILIO_WORKSPACE_SID")
ChannelID := os.Getenv("TWILIO_CHANNEL_ID")

Params = taskrouter.CapabilityTokenParams{
	AccountSid: AccountSid,
	AuthToken: AuthToken,
	WorkspaceSid: WorkspaceSid,
	ChannelID: ChannelID,
}

capabilityToken := taskrouter.CreateCapabilityToken(Params)
token, err := capabilityToken.ToJwt()
```

## Local Usage

### Building

To build *twilio-go* run:

```bash
go build ./...
```

### Testing

To execute the test suite run:

```bash
go test ./...
```

### Generating Local Documentation

To generate documentation, from the root directory:

```bash
godoc -http=localhost:{port number}
```

Then, navigate to `http://localhost:{port number}/pkg/github.com/twilio/twilio-go` in your local browser.

Example:

```bash
godoc -http=localhost:6060
```

http://localhost:6060/pkg/github.com/twilio/twilio-go

## Docker Image

The `Dockerfile` present in this repository and its respective `twilio/twilio-go` Docker image are currently used by Twilio for testing purposes only.

## Getting help

If you need help installing or using the library, please check the [Twilio Support Help Center](https://support.twilio.com) first, and [file a support ticket](https://twilio.com/help/contact) if you don't find an answer to your question.

If you've instead found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!

[apidocs]: https://www.twilio.com/docs/api

[libdocs]: https://pkg.go.dev/github.com/twilio/twilio-go?tab=versions