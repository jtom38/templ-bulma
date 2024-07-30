package v1

const (
	TriggerOnce = ""
)

type HxParam struct {
	HxGet     string
	HxPost    string
	HxPut     string
	HxPatch   string
	HxDelete  string

	// https://v1.htmx.org/attributes/hx-trigger/
	HxTrigger string

	// https://v1.htmx.org/attributes/hx-indicator/
	HxIndicator string
}
