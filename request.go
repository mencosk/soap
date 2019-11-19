package soap

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Request struct {
	Url             string
	Header          http.Header
	PayloadRequest  interface{}
	PayloadResponse interface{}
	PayloadFault    interface{}
	RawRequest      *http.Request
	client          *Client
	Time            time.Time
}

//  SetUrl method is to set a webservice url field and its value in the current request.
//
// For Example: To set `http://mywebservice.com/km/add`.
// 		client.R().
//			SetUrl("http://mywebservice.com/km/add")
//
func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

// SetHeader method is to set a single header field and its value in the current request.
//
// For Example: To set `Content-Type` and `Accept` as `application/json`.
// 		client.R().
//			SetHeader("Content-Type", "text/xml; charset=utf-8").
//			SetHeader("Accept", "text/xml; charset=utf-8")
//			SetHeader("SOAPAction", "http://mywebservice.com/km/add/action")
//
// Also you can override header value, which was set at client instance level.
func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

// SetHeaders method sets multiple headers field and its values at one go in the current request.
//
// For Example: To set `Content-Type` and `Accept` as `text/xml; charset=utf-8`
// 			and `SOAPAction` as "http://mywebservice.com/km/add/action"
//
// 		client.R().
//			SetHeaders(map[string]string{
//				"Content-Type": "text/xml; charset=utf-8",
//				"Accept": "text/xml; charset=utf-8",
//				"SOAPAction" : "http://mywebservice.com/km/add/action",
//			})
// Also you can override header value, which was set at client instance level.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for h, v := range headers {
		r.SetHeader(h, v)
	}
	return r
}

//  SetPayloadRequest method is to set a PayloadRequest field and its value in the current request.
//
// For Example: To set `<?xml version="1.0" encoding="utf-8"?>
//						<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
// 						 <soap:Body>
//    						<NumberToWords xmlns="http://www.my.com/webservicesserver/">
//      						<Number>9</Number>
//    						</NumberToWords>
// 						 </soap:Body>
//                      </soap:Envelope>`.
//
//  convert xml to Go struct.   (using the following onlinetools https://www.onlinetool.io/xmltogo/)
//
//  `type NumberToWordsRequest struct {
//		XMLName xml.Name `xml:"soap:Envelope"`
//		Text    string   `xml:",chardata"`
//		Soap    string   `xml:"xmlns:soap,attr"`
//		Body    struct {
//			Text          string `xml:",chardata"`
//			NumberToWords struct {
//				Text   string `xml:",chardata"`
//				Xmlns  string `xml:"xmlns,attr"`
//				Number string `xml:"Number"`
//			} `xml:"NumberToWords"`
//		} `xml:"soap:Body"`
//	}`
//
//	payload := NumberToWordsRequest{}
//	payload.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
//	payload.Body.NumberToWords.Xmlns = "http://www.my.com/webservicesserver/"
//	payload.Body.NumberToWords.Number = "777"
//
// 		client.R().
//			SetPayloadRequest(&payload)
//
func (r *Request) SetPayloadRequest(payloadRequest interface{}) *Request {
	r.PayloadRequest = payloadRequest
	return r
}

//  SetPayloadResponse method is to set a PayloadResponse field and its value in the current request.
//
// For Example: To set `<?xml version="1.0" encoding="utf-8"?>
//						<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//  					<soap:Body>
//    						<NumberToWordsResponse xmlns="http://www.dataaccess.com/webservicesserver/">
//      						<NumberToWordsResult>string</NumberToWordsResult>
//    						</NumberToWordsResponse>
//  					</soap:Body>
//						</soap:Envelope>`.
//
//  convert xml to Go struct.   (using the following onlinetools https://www.onlinetool.io/xmltogo/)
//
//  `type NumberToWordsResponse struct {
//		XMLName xml.Name `xml:"Envelope"`
//		Text    string   `xml:",chardata"`
//		Soap    string   `xml:"soap,attr"`
//		Body    struct {
//			Text                  string `xml:",chardata"`
//			NumberToWordsResponse struct {
//				Text                string `xml:",chardata"`
//				Xmlns               string `xml:"xmlns,attr"`
//				NumberToWordsResult string `xml:"NumberToWordsResult"`
//			} `xml:"NumberToWordsResponse"`
//		} `xml:"Body"`
//	}`
//
// 		client.R().
//			SetPayloadResponse(&NumberToWordsResponse{})
//
func (r *Request) SetPayloadResponse(payloadResponse interface{}) *Request {
	r.PayloadResponse = getPointer(payloadResponse)
	return r
}

//  SetPayloadFault method is to set a PayloadFault field and its value in the current request.
//
// For Example: To set `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
//  					<soap:Body>
//    						<soap:Fault>
//      						<faultcode>soap:Server</faultcode>
//      						<faultstring>Error processing request</faultstring>
//      						<detail/>
//    						</soap:Fault>
//  					</soap:Body>
//						</soap:Envelope>`.
//
//  convert xml to Go struct.   (using the following onlinetools https://www.onlinetool.io/xmltogo/)
//
//  `type NumberToWordsFault struct {
//		XMLName xml.Name `xml:"Envelope"`
//		Text    string   `xml:",chardata"`
//		Soap    string   `xml:"soap,attr"`
//		Body    struct {
//			Text  string `xml:",chardata"`
//			Fault struct {
//				Text        string `xml:",chardata"`
//				Faultcode   string `xml:"faultcode"`
//				Faultstring string `xml:"faultstring"`
//				Detail      string `xml:"detail"`
//			} `xml:"Fault"`
//		} `xml:"Body"`
//	}`
//
// 		client.R().
//			SetPayloadFault(&NumberToWordsFault{})
//
func (r *Request) SetPayloadFault(payloadFault interface{}) *Request {
	r.PayloadFault = getPointer(payloadFault)
	return r
}

// The Call method Execute the request
func (r *Request) Call() (*Response, error) {

	marshalRequest, _ := xml.Marshal(r.PayloadRequest)
	req, err := http.NewRequest("POST", r.Url, bytes.NewReader(marshalRequest))
	if err != nil {
		log.Fatalf("failed to create POST request %s", err)
		return nil, err
	}
	// Create headers
	req.Header = r.Header
	req.Close = true

	r.Time = time.Now()
	resp, err := r.client.httpClient.Do(req)
	endTime := time.Now()
	if err != nil {
		log.Fatalf("failed to send request %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	response := &Response{
		Request:     r,
		RawResponse: resp,
		receivedAt:  endTime,
	}

	if response.payloadResponse, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := xml.Unmarshal(response.payloadResponse, response.Request.PayloadFault)
		if err != nil {
			log.Printf("filed trying to convert fault fault response %s", err)
			return response, err
		}
		return response, err
	}
	err = xml.Unmarshal(response.payloadResponse, response.Request.PayloadResponse)
	return response, err
}

func getPointer(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return v
	}
	return reflect.New(vv.Type()).Interface()
}
