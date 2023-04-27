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

// VerifyV2VerificationAttempt struct for VerifyV2VerificationAttempt
type VerifyV2VerificationAttempt struct {
	// The SID that uniquely identifies the verification attempt.
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the verification.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the verify service that generated this attempt.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the verification that generated this attempt.
	VerificationSid *string `json:"verification_sid,omitempty"`
	// The date this Attempt was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Attempt was updated
	DateUpdated      *time.Time `json:"date_updated,omitempty"`
	ConversionStatus *string    `json:"conversion_status,omitempty"`
	Channel          *string    `json:"channel,omitempty"`
	// An object containing the charge for this verification attempt.
	Price *interface{} `json:"price,omitempty"`
	// An object containing the channel specific information for an attempt.
	ChannelData *interface{} `json:"channel_data,omitempty"`
	Url         *string      `json:"url,omitempty"`
}