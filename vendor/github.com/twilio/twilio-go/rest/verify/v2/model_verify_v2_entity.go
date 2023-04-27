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
	"time"
)

// VerifyV2Entity struct for VerifyV2Entity
type VerifyV2Entity struct {
	// A string that uniquely identifies this Entity.
	Sid *string `json:"sid,omitempty"`
	// Unique external identifier of the Entity
	Identity *string `json:"identity,omitempty"`
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The date this Entity was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Entity was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
}