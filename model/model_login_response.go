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

type LoginResponse struct {

	Id int32 `json:"id,omitempty"`

	Token string `json:"token,omitempty"`

	Role string `json:"role,omitempty"`

	Message string `json:"message,omitempty"`
}