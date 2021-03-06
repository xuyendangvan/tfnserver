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

import (
	"time"
)

type Student struct {

	Id int64 `json:"id,omitempty"`

	ClassID int32 `json:"classID,omitempty"`

	ParentID int32 `json:"parentID,omitempty"`

	Name string `json:"name,omitempty"`

	Birthday time.Time `json:"birthday,omitempty"`

	Phone string `json:"phone,omitempty"`

	// User Status: 1- presence, 2- absence
	StudentStatus int32 `json:"studentStatus,omitempty"`
}
