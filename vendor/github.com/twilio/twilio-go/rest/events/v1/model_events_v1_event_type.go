/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
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

// EventsV1EventType struct for EventsV1EventType
type EventsV1EventType struct {
	// The Event Type identifier.
	Type *string `json:"type,omitempty"`
	// The Schema identifier for this Event Type.
	SchemaId *string `json:"schema_id,omitempty"`
	// The date this Event Type was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Event Type was updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Event Type description.
	Description *string `json:"description,omitempty"`
	// The URL of this resource.
	Url   *string                 `json:"url,omitempty"`
	Links *map[string]interface{} `json:"links,omitempty"`
}