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

import "time"

type MenuDetail struct {
	Id int32 `json:"id,omitempty"`

	MenuID int32 `json:"menuID,omitempty"`

	SessionID int32 `json:"sessionID,omitempty"`

	FoodName string `json:"foodName,omitempty"`

	Note string `json:"note,omitempty"`

	CanteenID int32 `json:"canteenID,omitempty"`

	DateCreated time.Time `json:"dateCreated,omitempty"`
}
