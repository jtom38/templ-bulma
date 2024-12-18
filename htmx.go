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
	HxGet string

	// https://htmx.org/attributes/hx-post/
	HxPost string

	// https://htmx.org/attributes/hx-put/
	HxPut string

	// https://htmx.org/attributes/hx-patch/
	HxPatch string

	// https://htmx.org/attributes/hx-delete/
	HxDelete string

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

	// if set client-side will do a full refresh of the page
	//HxRefresh bool
}
