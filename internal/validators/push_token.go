package validators

type AddPushToken struct {
	DeviceId  string `json:"deviceId" binding:"required,max=255"`
	Name      string `json:"name" binding:"omitempty,max=255"`
	Os        string `json:"os" binding:"omitempty"`
	PushToken string `json:"pushToken" binding:"required,max=255"`
}

type DeletePushToken struct {
	PushToken string `json:"pushToken" binding:"required,max=255"`
}
