package enum

type IconType string

const (
	ICON IconType = "icon"
	SVG  IconType = "svg"
)

var IconTypeValues = map[IconType]string{
	ICON: "icon",
	SVG:  "svg",
}