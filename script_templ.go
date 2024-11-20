// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package bulma

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func UseBulmaCdn() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bulma@1.0.1/css/bulma.min.css\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// This is required for models to work correctly.
// This needs to be placed at the end of your body tag.
//
// https://bulma.io/documentation/components/modal/#javascript-implementation-example
func RequiredModelScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n    document.addEventListener('DOMContentLoaded', () => {\n    \n        // Functions to open and close a modal\n        function openModal($el) {\n            $el.classList.add('is-active');\n        }\n\n        function closeModal($el) {\n            $el.classList.remove('is-active');\n        }\n\n        function closeAllModals() {\n            (document.querySelectorAll('.modal') || []).forEach(($modal) => {\n                closeModal($modal);\n            });\n        }\n\n        // Add a click event on buttons to open a specific modal\n        (document.querySelectorAll('.js-modal-trigger') || []).forEach(($trigger) => {\n            const modal = $trigger.dataset.target;\n            const $target = document.getElementById(modal);\n\n            $trigger.addEventListener('click', () => {\n            openModal($target);\n            });\n        });\n\n        // Add a click event on various child elements to close the parent modal\n        (document.querySelectorAll('.modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button') || []).forEach(($close) => {\n            const $target = $close.closest('.modal');\n\n            $close.addEventListener('click', () => {\n            closeModal($target);\n            });\n        });\n\n        // Add a keyboard event to close all modals\n        document.addEventListener('keydown', (event) => {\n            if(event.key === \"Escape\") {\n                closeAllModals();\n            }\n        });\n    });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// This script will toggle the navbar on mobile.
// This needs to be placed at the end of your body tag.
//
// https://bulma.io/documentation/components/navbar/#navbar-menu
func RequiredNavbarToggleScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        document.addEventListener('DOMContentLoaded', () => {\n            // Get all \"navbar-burger\" elements\n            const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);\n\n            // Add a click event on each of them\n            $navbarBurgers.forEach( el => {\n                el.addEventListener('click', () => {\n                // Get the target from the \"data-target\" attribute\n                const target = el.dataset.target; \n                const $target = document.getElementById(target);\n\n                // Toggle the \"is-active\" class on both the \"navbar-burger\" and the \"navbar-menu\"\n                el.classList.toggle('is-active');\n                $target.classList.toggle('is-active');\n                });\n            });\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
