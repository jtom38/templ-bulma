package bulma

const (
	HxTriggerClick      = "clicked"
	HxTriggerMouseEnter = "mouseenter"

	HxSwapInnerHtml = "innerHtml"
	HxSwapOuterHtml = "outerHtml"
	HxSwapDefault   = HxSwapInnerHtml
	HxSwapBlank     = ""
)

type HtmxProperties struct {
	// The hx-get attribute will cause an element to issue a GET to the specified URL and swap the HTML into the DOM using a swap strategy
	//
	// https://htmx.org/attributes/hx-get/
	HxGet     string

	
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
}
