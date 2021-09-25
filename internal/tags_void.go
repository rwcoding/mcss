package internal

var voidTag = []string{
	"area",
	"base",
	"basefont",
	"br",
	"col",
	"command",
	"embed",
	"frame",
	"hr",
	"img",
	"input",
	"isindex",
	"keygen",
	"link",
	"meta",
	"param",
	"source",
	"track",
	"wbr",
}

func isVoidTag(tag string) bool {
	for _, v := range voidTag {
		if v == tag {
			return true
		}
	}
	return false
}
