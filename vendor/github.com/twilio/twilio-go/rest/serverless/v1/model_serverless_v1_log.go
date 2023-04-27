/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
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

// ServerlessV1Log struct for ServerlessV1Log
type ServerlessV1Log struct {
	// The unique string that identifies the Log resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the Log resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Log resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the environment in which the log occurred
	EnvironmentSid *string `json:"environment_sid,omitempty"`
	// The SID of the build that corresponds to the log
	BuildSid *string `json:"build_sid,omitempty"`
	// The SID of the deployment that corresponds to the log
	DeploymentSid *string `json:"deployment_sid,omitempty"`
	// The SID of the function whose invocation produced the log
	FunctionSid *string `json:"function_sid,omitempty"`
	// The SID of the request associated with the log
	RequestSid *string `json:"request_sid,omitempty"`
	Level      *string `json:"level,omitempty"`
	// The log message
	Message *string `json:"message,omitempty"`
	// The ISO 8601 date and time in GMT when the Log resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Log resource
	Url *string `json:"url,omitempty"`
}
