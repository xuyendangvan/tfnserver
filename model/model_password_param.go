package swagger

type PasswordParam struct {
	UserID int32 `json:"userID,omitempty"`

	Type int32 `json:"type,omitempty"`

	OldPassword string `json:"oldPassword,omitempty"`

	NewPassword string `json:"newPassword,omitempty"`
}
