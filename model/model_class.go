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

type Class struct {

	Id int64 `json:"id,omitempty"`

	YearID int32 `json:"yearID,omitempty"`

	Level int32 `json:"level,omitempty"`

	TeacherID int32 `json:"teacherID,omitempty"`

	Name string `json:"name,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`
}