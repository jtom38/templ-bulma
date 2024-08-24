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
	HxGet     string
	HxPost    string
	HxTrigger string
	HxTarget  string
	HxSwap    string
}
