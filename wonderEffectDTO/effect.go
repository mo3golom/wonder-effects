package wonderEffectDTO

type Effect struct {
	Id       string            `json:"id"`
	Settings map[string]string `json:"settings"`
}
