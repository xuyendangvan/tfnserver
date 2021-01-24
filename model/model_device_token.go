package swagger

type DeviceToken struct {
	Id int32 `json:"id,omitempty"`

	Type int32 `json:"type,omitempty"`

	DeviceToken string `json:"deviceToken,omitempty"`
}
