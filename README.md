# SOAP
Simple HTTP SOAP client library for Go. We solve common problem of consuming a SOAP-based web service, in distributed systems and application architecture so you can focus on delivering business value.

## News
* v1.0 released on Nov 19th, 2019.

## Features
* SOAP Client.
* Simple way to chain methods for settings and request.
* Easy to use.
* Well tested client library.

## Installation
```go
go get github.com/mencosk/soap
```

## Usage
The following samples will assist you to consuming a SOAP-based web service.

```go
// Import soap into your code.
import "github.com/mencosk/soap"
```

#### Example

```go
// Create a SOAP Client
client := soap.New()

// Create a SOAP Request
resp, err := client.R().
		SetUrl("https://www.dataaccess.com/webservicesserver/numberconversion.wso?op=NumberToWords").
		SetHeader("Content-Type", "text/xml; charset=utf-8").
		SetHeader("SOAPAction", "https://www.dataaccess.com/webservicesserver/NumberConversion.wso?op=NumberToWords").
		SetPayloadRequest(&request).
		SetPayloadResponse(&NumberToWordsResponse{}).
		SetPayloadFault(&NumberToWordsFault{}).
		Call()
if err != nil {
	log.Fatalf("An error occurred in the SOAP service call %s", err)
}

if resp.RawResponse.StatusCode != http.StatusOK {
	fmt.Println(resp.PayloadResultError().(*NumberToWordsFault).Body.Fault.Faultstring)
}

response := resp.PayloadResult().(*NumberToWordsResponse)
fmt.Println(response.Body.NumberToWordsResponse.NumberToWordsResult)



// NOTE: Generate the Request, Response and Fault
// Go Struct using onlinetools like xmltogo
// https://www.onlinetool.io/xmltogo/

// Request
type NumberToWordsRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Body    struct {
		Text          string `xml:",chardata"`
		NumberToWords struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			UbiNum string `xml:"ubiNum"`
		} `xml:"NumberToWords"`
	} `xml:"soap:Body"`
}

// Response
type NumberToWordsResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                  string `xml:",chardata"`
		NumberToWordsResponse struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			NumberToWordsResult string `xml:"NumberToWordsResult"`
		} `xml:"NumberToWordsResponse"`
	} `xml:"Body"`
}

// Fault
type NumberToWordsFault struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text  string `xml:",chardata"`
		Fault struct {
			Text        string `xml:",chardata"`
			Faultcode   string `xml:"faultcode"`
			Faultstring string `xml:"faultstring"`
			Detail      string `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

```

## Contribution
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Creator
[Kevin Mencos](https://github.com/mencosk)(mencosk@gmail.com)


## License
soap released under [MIT](https://github.com/mencosk/soap) license, refer [LICENSE](LICENSE) file.