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

type StudentStatistic struct {

	Id int32 `json:"id,omitempty"`

	Quater int32 `json:"quater,omitempty"`

	StudentID int32 `json:"studentID,omitempty"`

	ClassID int32 `json:"classID,omitempty"`

	Total int32 `json:"total,omitempty"`

	// Nghi khong phep
	Absences int32 `json:"absences,omitempty"`

	// Nghi co phep
	RequestedAbsences int32 `json:"requestedAbsences,omitempty"`

	UsedMeals int32 `json:"usedMeals,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`
}
