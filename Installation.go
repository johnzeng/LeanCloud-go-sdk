package lean

type Installation struct {
	LeanClassesBase
	Badage          string   `json:"badge,omitempty"`
	Channels        []string `json:"channels,omitempty"`
	DeviceProfile   string   `json:"deviceProfile,omitempty"`
	DeviceToken     string   `json:"deviceToken,omitempty"`
	DeviceType      string   `json:"deviceType,omitempty"`
	ID              string   `json:"ID,omitempty"`
	InstallationId  string   `json:"installationId,omitempty"`
	SubscriptionUri string   `json:"subscriptionUri,omitempty"`
	TimeZone        string   `json:"timeZone,omitempty"`
}

func (this *Installation) GetClassName() string {
	return "installation"
}
