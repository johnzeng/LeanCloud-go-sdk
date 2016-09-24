package lean

type Installation struct {
	LeanClassesBase
	Badage          string   `json:"badge"`
	Channels        []string `json:"channels"`
	DeviceProfile   string   `json:"deviceProfile"`
	DeviceToken     string   `json:"deviceToken"`
	DeviceType      string   `json:"deviceType"`
	ID              string   `json:"ID"`
	InstallationId  string   `json:"installationId"`
	SubscriptionUri string   `json:"subscriptionUri"`
	TimeZone        string   `json:"timeZone"`
}

func (this *Installation) GetClassName() string {
	return "installation"
}
