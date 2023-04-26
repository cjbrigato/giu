package giu

import "github.com/AllenDang/cimgui-go"

// Here are the style IDs for styling imgui apps.
// For details about each of attributes read comment above them.

// go-generate String() andStringToEnum methods
//go:generate stringer -type=StyleColorID,StyleVarID -output=StyleIDs_string.go -linecomment
// NOTE: string2enum is https://github.com/mewspring/tools
//go:generate string2enum -samepkg -type=StyleColorID,StyleVarID -output=StyleIDs_string2enum.go -linecomment

// StyleColorID identifies a color in the UI style.
type StyleColorID imgui.Col

// StyleColor identifier.
// NOTE: comments are used for CSS conversion and are generated by stringer and string2enum.
const (
	StyleColorText                  = StyleColorID(imgui.ColText)                  // color
	StyleColorTextDisabled          = StyleColorID(imgui.ColTextDisabled)          // disabled-color
	StyleColorWindowBg              = StyleColorID(imgui.ColWindowBg)              // background-color
	StyleColorChildBg               = StyleColorID(imgui.ColChildBg)               // child-background-color
	StyleColorPopupBg               = StyleColorID(imgui.ColPopupBg)               // popup-background-color
	StyleColorBorder                = StyleColorID(imgui.ColBorder)                // border-color
	StyleColorBorderShadow          = StyleColorID(imgui.ColBorderShadow)          // border-shadow-color
	StyleColorFrameBg               = StyleColorID(imgui.ColFrameBg)               // frame-background-color
	StyleColorFrameBgHovered        = StyleColorID(imgui.ColFrameBgHovered)        // frame-background-hovered-color
	StyleColorFrameBgActive         = StyleColorID(imgui.ColFrameBgActive)         // frame-background-active-color
	StyleColorTitleBg               = StyleColorID(imgui.ColTitleBg)               // title-background-color
	StyleColorTitleBgActive         = StyleColorID(imgui.ColTitleBgActive)         // title-background-active-color
	StyleColorTitleBgCollapsed      = StyleColorID(imgui.ColTitleBgCollapsed)      // title-background-collapsed-color
	StyleColorMenuBarBg             = StyleColorID(imgui.ColMenuBarBg)             // menu-bar-background-color
	StyleColorScrollbarBg           = StyleColorID(imgui.ColScrollbarBg)           // scrollbar-background-color
	StyleColorScrollbarGrab         = StyleColorID(imgui.ColScrollbarGrab)         // scrollbar-grab-color
	StyleColorScrollbarGrabHovered  = StyleColorID(imgui.ColScrollbarGrabHovered)  // scrollbar-grab-hovered-color
	StyleColorScrollbarGrabActive   = StyleColorID(imgui.ColScrollbarGrabActive)   // scrollbar-grab-active-color
	StyleColorCheckMark             = StyleColorID(imgui.ColCheckMark)             // checkmark-color
	StyleColorSliderGrab            = StyleColorID(imgui.ColSliderGrab)            // slider-grab-color
	StyleColorSliderGrabActive      = StyleColorID(imgui.ColSliderGrabActive)      // slider-grab-active-color
	StyleColorButton                = StyleColorID(imgui.ColButton)                // button-color
	StyleColorButtonHovered         = StyleColorID(imgui.ColButtonHovered)         // button-hovered-color
	StyleColorButtonActive          = StyleColorID(imgui.ColButtonActive)          // button-active-color
	StyleColorHeader                = StyleColorID(imgui.ColHeader)                // header-color
	StyleColorHeaderHovered         = StyleColorID(imgui.ColHeaderHovered)         // header-hovered-color
	StyleColorHeaderActive          = StyleColorID(imgui.ColHeaderActive)          // header-active-color
	StyleColorSeparator             = StyleColorID(imgui.ColSeparator)             // separator-color
	StyleColorSeparatorHovered      = StyleColorID(imgui.ColSeparatorHovered)      // separator-hovered-color
	StyleColorSeparatorActive       = StyleColorID(imgui.ColSeparatorActive)       // separator-active-color
	StyleColorResizeGrip            = StyleColorID(imgui.ColResizeGrip)            // resize-grip-color
	StyleColorResizeGripHovered     = StyleColorID(imgui.ColResizeGripHovered)     // resize-grip-hovered-color
	StyleColorResizeGripActive      = StyleColorID(imgui.ColResizeGripActive)      // resize-grip-active-color
	StyleColorTab                   = StyleColorID(imgui.ColTab)                   // tab-color
	StyleColorTabHovered            = StyleColorID(imgui.ColTabHovered)            // tab-hovered-color
	StyleColorTabActive             = StyleColorID(imgui.ColTabActive)             // tab-active-color
	StyleColorTabUnfocused          = StyleColorID(imgui.ColTabUnfocused)          // tab-unfocused-color
	StyleColorTabUnfocusedActive    = StyleColorID(imgui.ColTabUnfocusedActive)    // tab-unfocused-active-color
	StyleColorPlotLines             = StyleColorID(imgui.ColPlotLines)             // plot-lines-color
	StyleColorPlotLinesHovered      = StyleColorID(imgui.ColPlotLinesHovered)      // plot-lines-hovered-color
	StyleColorProgressBarActive     = StyleColorPlotLinesHovered                   // progress-bar-active-color
	StyleColorPlotHistogram         = StyleColorID(imgui.ColPlotHistogram)         // plot-histogram-color
	StyleColorPlotHistogramHovered  = StyleColorID(imgui.ColPlotHistogramHovered)  // plot-histogram-hovered-color
	StyleColorTableHeaderBg         = StyleColorID(imgui.ColTableHeaderBg)         // table-header-background-color
	StyleColorTableBorderStrong     = StyleColorID(imgui.ColTableBorderStrong)     // table-border-strong-color
	StyleColorTableBorderLight      = StyleColorID(imgui.ColTableBorderLight)      // table-border-light-color
	StyleColorTableRowBg            = StyleColorID(imgui.ColTableRowBg)            // table-row-background-color
	StyleColorTableRowBgAlt         = StyleColorID(imgui.ColTableRowBgAlt)         // table-row-alternate-background-color
	StyleColorTextSelectedBg        = StyleColorID(imgui.ColTextSelectedBg)        // text-selected-background-color
	StyleColorDragDropTarget        = StyleColorID(imgui.ColDragDropTarget)        // drag-drop-target-color
	StyleColorNavHighlight          = StyleColorID(imgui.ColNavHighlight)          // navigation-highlight-color
	StyleColorNavWindowingHighlight = StyleColorID(imgui.ColNavWindowingHighlight) // windowing-highlight-color
	StyleColorNavWindowingDimBg     = StyleColorID(imgui.ColNavWindowingDimBg)     // windowing-dim-background-color
	StyleColorModalWindowDimBg      = StyleColorID(imgui.ColModalWindowDimBg)      // modal-window-dim-background-color
)

// StyleVarID identifies a style variable in the UI style.
type StyleVarID imgui.StyleVar

// Style IDs.
// comments at same line is a CSS name.
const (
	// StyleVarAlpha is a float.
	StyleVarAlpha = StyleVarID(imgui.StyleVarAlpha) // alpha
	// StyleVarDisabledAlpha is a float.
	StyleVarDisabledAlpha = StyleVarID(imgui.StyleVarDisabledAlpha) // disabled-alpha
	// StyleVarWindowPadding is a Vec2.
	StyleVarWindowPadding = StyleVarID(imgui.StyleVarWindowPadding) // window-padding
	// StyleVarWindowRounding is a float.
	StyleVarWindowRounding = StyleVarID(imgui.StyleVarWindowRounding) // window-rounding
	// StyleVarWindowBorderSize is a float.
	StyleVarWindowBorderSize = StyleVarID(imgui.StyleVarWindowBorderSize) // window-border-size
	// StyleVarWindowMinSize is a Vec2.
	StyleVarWindowMinSize = StyleVarID(imgui.StyleVarWindowMinSize) // window-min-size
	// StyleVarWindowTitleAlign is a Vec2.
	StyleVarWindowTitleAlign = StyleVarID(imgui.StyleVarWindowTitleAlign) // window-title-align
	// StyleVarChildRounding is a float.
	StyleVarChildRounding = StyleVarID(imgui.StyleVarChildRounding) // child-rounding
	// StyleVarChildBorderSize is a float.
	StyleVarChildBorderSize = StyleVarID(imgui.StyleVarChildBorderSize) // child-border-size
	// StyleVarPopupRounding is a float.
	StyleVarPopupRounding = StyleVarID(imgui.StyleVarPopupRounding) // popup-rounding
	// StyleVarPopupBorderSize is a float.
	StyleVarPopupBorderSize = StyleVarID(imgui.StyleVarPopupBorderSize) // popup-border-size
	// StyleVarFramePadding is a Vec2.
	StyleVarFramePadding = StyleVarID(imgui.StyleVarFramePadding) // frame-padding
	// StyleVarFrameRounding is a float.
	StyleVarFrameRounding = StyleVarID(imgui.StyleVarFrameRounding) // frame-rounding
	// StyleVarFrameBorderSize is a float.
	StyleVarFrameBorderSize = StyleVarID(imgui.StyleVarFrameBorderSize) // frame-border-size
	// StyleVarItemSpacing is a Vec2.
	StyleVarItemSpacing = StyleVarID(imgui.StyleVarItemSpacing) // item-spacing
	// StyleVarItemInnerSpacing is a Vec2.
	StyleVarItemInnerSpacing = StyleVarID(imgui.StyleVarItemInnerSpacing) // item-inner-spacing
	// StyleVarIndentSpacing is a float.
	StyleVarIndentSpacing = StyleVarID(imgui.StyleVarIndentSpacing) // indent-spacing
	// StyleVarScrollbarSize is a float.
	StyleVarScrollbarSize = StyleVarID(imgui.StyleVarScrollbarSize) // scrollbar-size
	// StyleVarScrollbarRounding is a float.
	StyleVarScrollbarRounding = StyleVarID(imgui.StyleVarScrollbarRounding) // scrollbar-rounding
	// StyleVarGrabMinSize is a float.
	StyleVarGrabMinSize = StyleVarID(imgui.StyleVarGrabMinSize) // grab-min-size
	// StyleVarGrabRounding is a float.
	StyleVarGrabRounding = StyleVarID(imgui.StyleVarGrabRounding) // grab-rounding
	// StyleVarTabRounding is a float.
	StyleVarTabRounding = StyleVarID(imgui.StyleVarTabRounding) // tab-rounding
	// StyleVarButtonTextAlign is a Vec2.
	StyleVarButtonTextAlign = StyleVarID(imgui.StyleVarButtonTextAlign) // button-text-align
	// StyleVarSelectableTextAlign is a Vec2.
	StyleVarSelectableTextAlign = StyleVarID(imgui.StyleVarSelectableTextAlign) // selectable-text-align
)

// IsVec2 returns true if the style var id should be processed as imgui.Vec2
// if not, it is interpreted as float32.
func (i StyleVarID) IsVec2() bool {
	lookup := map[StyleVarID]bool{
		// StyleVarWindowPadding is a Vec2.
		StyleVarWindowPadding:    true,
		StyleVarWindowMinSize:    true,
		StyleVarWindowTitleAlign: true,
		StyleVarFramePadding:     true,
		StyleVarItemSpacing:      true,
		// StyleVarItemInnerSpacing is a Vec2.
		StyleVarItemInnerSpacing:    true,
		StyleVarButtonTextAlign:     true,
		StyleVarSelectableTextAlign: true,
	}

	result, ok := lookup[i]

	return result && ok
}
