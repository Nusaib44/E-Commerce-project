/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// TrusthubV1TrustProduct struct for TrusthubV1TrustProduct
type TrusthubV1TrustProduct struct {
	// The unique string that identifies the resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique string of a policy.
	PolicySid *string `json:"policy_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	Status       *string `json:"status,omitempty"`
	// The ISO 8601 date and time in GMT when the resource will be valid until.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// The email address
	Email *string `json:"email,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"status_callback,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Customer-Profile resource
	Url *string `json:"url,omitempty"`
	// The URLs of the Assigned Items of the Customer-Profile resource
	Links *map[string]interface{} `json:"links,omitempty"`
}
