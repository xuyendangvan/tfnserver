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

type Teacher struct {
	Id int64 `json:"id,omitempty"`

	UserID int64 `json:"userID,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	LoginName string `json:"loginName,omitempty"`

	Address string `json:"address,omitempty"`

	AddressCity string `json:"addressCity,omitempty"`

	AddressDistrict string `json:"addressDistrict,omitempty"`

	AddressStreet string `json:"addressStreet,omitempty"`

	Phone string `json:"phone,omitempty"`

	Photo string `json:"photo,omitempty"`

	Status int32 `json:"status,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`
}
