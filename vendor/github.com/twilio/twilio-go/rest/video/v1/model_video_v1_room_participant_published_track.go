/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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

// VideoV1RoomParticipantPublishedTrack struct for VideoV1RoomParticipantPublishedTrack
type VideoV1RoomParticipantPublishedTrack struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Participant resource with the published track
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the Room resource where the track is published
	RoomSid *string `json:"room_sid,omitempty"`
	// The track name
	Name *string `json:"name,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the track is enabled
	Enabled *bool   `json:"enabled,omitempty"`
	Kind    *string `json:"kind,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
