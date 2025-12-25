package enums

// ViewType defines the nature and behavior of a View entity.
type ViewType string

const (
	// ViewTypeRoot (T) is a virtual root node, used as a mount point for view trees in different scopes. It is not displayed.
	ViewTypeRoot ViewType = "T"
	// ViewTypeGroup (G) is a pure visual grouping in a menu. It is not clickable and has no link.
	ViewTypeGroup ViewType = "G"
	// ViewTypeMenu (M) is a core navigation item, typically with a link that navigates to a PAGE.
	ViewTypeMenu ViewType = "M"
	// ViewTypeLink (L) is an external link that opens in a new tab.
	ViewTypeLink ViewType = "L"
	// ViewTypePage (P) is the final destination of navigation, a container for content.
	ViewTypePage ViewType = "P"
	// ViewTypeButton (B) is an action trigger displayed on a page.
	ViewTypeButton ViewType = "B"
	// ViewTypeElement (E) is a generic UI element like a tab or a table column that requires permission control.
	ViewTypeElement ViewType = "E"
	// ViewTypeRedirect (R) is a route that performs a redirect.
	ViewTypeRedirect ViewType = "R"
	// ViewTypeUnknown (U) is an unknown or undefined type.
	ViewTypeUnknown ViewType = "U"
)
