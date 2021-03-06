/*
 * Demo App
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Schedule struct {

	Id int32 `json:"id,omitempty"`

	Level int32 `json:"level,omitempty"`

	DayOfWeek int32 `json:"dayOfWeek,omitempty"`

	ClassID int32 `json:"classID,omitempty"`

	AssignedDate string `json:"assignedDate,omitempty"`

	Note string `json:"note,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`
}
