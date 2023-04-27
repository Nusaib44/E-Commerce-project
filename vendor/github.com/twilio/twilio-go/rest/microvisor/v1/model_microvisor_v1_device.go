/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Microvisor
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

// MicrovisorV1Device struct for MicrovisorV1Device
type MicrovisorV1Device struct {
	// A string that uniquely identifies this Device.
	Sid *string `json:"sid,omitempty"`
	// A developer-defined string that uniquely identifies the Device.
	UniqueName *string `json:"unique_name,omitempty"`
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// Information about the target App and the App reported by this Device.
	App *interface{} `json:"app,omitempty"`
	// Object specifying whether application logging is enabled for this Device.
	Logging *interface{} `json:"logging,omitempty"`
	// The date that this Device was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Device was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// The absolute URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
}
