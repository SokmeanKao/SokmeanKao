package ui

type Regions struct {
	HeaderH  int
	FooterH  int
	BodyH    int
	SidebarW int
	ContentW int
}

// ComputeRegions derives frame sizes for a terminal.
// Header/footer content targets ~3 rows including border; body min 10 when possible.
func ComputeRegions(termW, termH int) Regions {
	headerH := 3
	footerH := 3
	bodyH := Max(1, termH-headerH-footerH)
	if termH >= 16 {
		bodyH = Max(10, bodyH)
		if headerH+footerH+bodyH > termH {
			bodyH = Max(1, termH-headerH-footerH)
		}
	}

	sidebarW := 22
	if termW < 60 {
		sidebarW = Max(12, termW/3)
	}
	contentW := Max(1, termW-sidebarW)
	if sidebarW+contentW > termW {
		contentW = Max(1, termW-sidebarW)
	}

	return Regions{
		HeaderH:  headerH,
		FooterH:  footerH,
		BodyH:    bodyH,
		SidebarW: sidebarW,
		ContentW: contentW,
	}
}
