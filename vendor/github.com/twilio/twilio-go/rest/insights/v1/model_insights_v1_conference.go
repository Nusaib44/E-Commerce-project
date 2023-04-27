/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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

// InsightsV1Conference struct for InsightsV1Conference
type InsightsV1Conference struct {
	// Conference SID.
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// Custom label for the conference.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Conference creation date/time.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Timestamp in ISO 8601 format when the conference started.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Conference end date/time.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Conference duration in seconds.
	DurationSeconds *int `json:"duration_seconds,omitempty"`
	// Duration of the conference in seconds.
	ConnectDurationSeconds *int    `json:"connect_duration_seconds,omitempty"`
	Status                 *string `json:"status,omitempty"`
	// Max participants specified in config.
	MaxParticipants *int `json:"max_participants,omitempty"`
	// Actual maximum concurrent participants.
	MaxConcurrentParticipants *int `json:"max_concurrent_participants,omitempty"`
	// Unique conference participants.
	UniqueParticipants *int    `json:"unique_participants,omitempty"`
	EndReason          *string `json:"end_reason,omitempty"`
	// Call SID that ended the conference.
	EndedBy              *string `json:"ended_by,omitempty"`
	MixerRegion          *string `json:"mixer_region,omitempty"`
	MixerRegionRequested *string `json:"mixer_region_requested,omitempty"`
	// Boolean. Indicates whether recording was enabled.
	RecordingEnabled *bool `json:"recording_enabled,omitempty"`
	// Potential issues detected during the conference.
	DetectedIssues *interface{} `json:"detected_issues,omitempty"`
	// Tags for detected conference conditions and participant behaviors.
	Tags *[]string `json:"tags,omitempty"`
	// Object. Contains details about conference tags.
	TagInfo         *interface{} `json:"tag_info,omitempty"`
	ProcessingState *string      `json:"processing_state,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
}