package main

import (
	"encoding/xml"
	"fmt"
	"github.com/mencosk/soap"
	"log"
	"net/http"
)

func main() {

	// make request
	request := NumberToWordsRequest{}
	request.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.NumberToWords.Xmlns = "http://www.dataaccess.com/webservicesserver/"
	request.Body.NumberToWords.UbiNum = "777"

	client := soap.New()
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
}

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

// Fault Error
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
