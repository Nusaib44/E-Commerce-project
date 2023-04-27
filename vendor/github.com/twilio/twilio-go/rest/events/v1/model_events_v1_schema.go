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

// EventsV1Schema struct for EventsV1Schema
type EventsV1Schema struct {
	// Schema Identifier.
	Id *string `json:"id,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The date that the latest schema version was created.
	LatestVersionDateCreated *time.Time `json:"latest_version_date_created,omitempty"`
	// Latest schema version.
	LatestVersion *int `json:"latest_version,omitempty"`
}
