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

type User struct {
	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	// User Role: 1- admin, 2- teacher, 3- parent
	Role int32 `json:"role,omitempty"`

	// User Status: 1- active, 2- inactive
	UserStatus int32 `json:"userStatus,omitempty"`

	Token string `json:"token,omitempty"`
}
