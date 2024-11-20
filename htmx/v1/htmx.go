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

type HtmxProperties struct {
	// The hx-get attribute will cause an element to issue a GET to the specified URL and swap the HTML into the DOM using a swap strategy
	//
	// https://htmx.org/attributes/hx-get/
	HxGet string

	HxPost    string
	HxTrigger string
	HxTarget  string
	HxSwap    string

	// The hx-boost attribute allows you to “boost” normal anchors and form tags to use AJAX instead
	//
	// https://htmx.org/attributes/hx-boost/
	HxBoost bool

	// The hx-push-url attribute allows you to push a URL into the browser location history.
	//
	// https://htmx.org/attributes/hx-push-url/
	HxPushUrl bool

	// The hx-confirm attribute allows you to confirm an action before issuing a request.
	// This can be useful in cases where the action is destructive and you want to ensure that the user really wants to do it.
	//
	//https://htmx.org/attributes/hx-confirm/
	HxConfirm string
}