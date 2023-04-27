/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Fetch a specific verification attempt.
func (c *ApiService) FetchVerificationAttempt(Sid string) (*VerifyV2VerificationAttempt, error) {
	path := "/v2/Attempts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2VerificationAttempt{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVerificationAttempt'
type ListVerificationAttemptParams struct {
	// Datetime filter used to query Verification Attempts created after this datetime. Given as GMT in RFC 2822 format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Datetime filter used to query Verification Attempts created before this datetime. Given as GMT in RFC 2822 format.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Destination of a verification. It is phone number in E.164 format.
	ChannelDataTo *string `json:"ChannelData.To,omitempty"`
	// Filter used to query Verification Attempts sent to the specified destination country.
	Country *string `json:"Country,omitempty"`
	// Filter used to query Verification Attempts by communication channel. Valid values are `SMS` and `CALL`
	Channel *string `json:"Channel,omitempty"`
	// Filter used to query Verification Attempts by verify service. Only attempts of the provided SID will be returned.
	VerifyServiceSid *string `json:"VerifyServiceSid,omitempty"`
	// Filter used to return all the Verification Attempts of a single verification. Only attempts of the provided verification SID will be returned.
	VerificationSid *string `json:"VerificationSid,omitempty"`
	// Filter used to query Verification Attempts by conversion status. Valid values are `UNCONVERTED`, for attempts that were not converted, and `CONVERTED`, for attempts that were confirmed.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListVerificationAttemptParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListVerificationAttemptParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListVerificationAttemptParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListVerificationAttemptParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListVerificationAttemptParams) SetChannelDataTo(ChannelDataTo string) *ListVerificationAttemptParams {
	params.ChannelDataTo = &ChannelDataTo
	return params
}
func (params *ListVerificationAttemptParams) SetCountry(Country string) *ListVerificationAttemptParams {
	params.Country = &Country
	return params
}
func (params *ListVerificationAttemptParams) SetChannel(Channel string) *ListVerificationAttemptParams {
	params.Channel = &Channel
	return params
}
func (params *ListVerificationAttemptParams) SetVerifyServiceSid(VerifyServiceSid string) *ListVerificationAttemptParams {
	params.VerifyServiceSid = &VerifyServiceSid
	return params
}
func (params *ListVerificationAttemptParams) SetVerificationSid(VerificationSid string) *ListVerificationAttemptParams {
	params.VerificationSid = &VerificationSid
	return params
}
func (params *ListVerificationAttemptParams) SetStatus(Status string) *ListVerificationAttemptParams {
	params.Status = &Status
	return params
}
func (params *ListVerificationAttemptParams) SetPageSize(PageSize int) *ListVerificationAttemptParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListVerificationAttemptParams) SetLimit(Limit int) *ListVerificationAttemptParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of VerificationAttempt records from the API. Request is executed immediately.
func (c *ApiService) PageVerificationAttempt(params *ListVerificationAttemptParams, pageToken, pageNumber string) (*ListVerificationAttemptResponse, error) {
	path := "/v2/Attempts"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.ChannelDataTo != nil {
		data.Set("ChannelData.To", *params.ChannelDataTo)
	}
	if params != nil && params.Country != nil {
		data.Set("Country", *params.Country)
	}
	if params != nil && params.Channel != nil {
		data.Set("Channel", *params.Channel)
	}
	if params != nil && params.VerifyServiceSid != nil {
		data.Set("VerifyServiceSid", *params.VerifyServiceSid)
	}
	if params != nil && params.VerificationSid != nil {
		data.Set("VerificationSid", *params.VerificationSid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVerificationAttemptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists VerificationAttempt records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListVerificationAttempt(params *ListVerificationAttemptParams) ([]VerifyV2VerificationAttempt, error) {
	response, errors := c.StreamVerificationAttempt(params)

	records := make([]VerifyV2VerificationAttempt, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams VerificationAttempt records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamVerificationAttempt(params *ListVerificationAttemptParams) (chan VerifyV2VerificationAttempt, chan error) {
	if params == nil {
		params = &ListVerificationAttemptParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VerifyV2VerificationAttempt, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageVerificationAttempt(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamVerificationAttempt(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamVerificationAttempt(response *ListVerificationAttemptResponse, params *ListVerificationAttemptParams, recordChannel chan VerifyV2VerificationAttempt, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Attempts
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListVerificationAttemptResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListVerificationAttemptResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListVerificationAttemptResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVerificationAttemptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
