package text

import "strings"

const (
	black      = "§0"
	darkBlue   = "§1"
	darkGreen  = "§2"
	darkAqua   = "§3"
	darkRed    = "§4"
	darkPurple = "§5"
	gold       = "§6"
	grey       = "§7"
	darkGrey   = "§8"
	blue       = "§9"
	green      = "§a"
	aqua       = "§b"
	red        = "§c"
	purple     = "§d"
	yellow     = "§e"
	white      = "§f"
	darkYellow = "§g"

	obfuscated    = "§k"
	bold          = "§l"
	strikethrough = "§m"
	underline     = "§n"
	italic        = "§o"
	reset         = "§r"
)

const (
	ansiBlack      = "\x1b[38;5;16m"
	ansiDarkBlue   = "\x1b[38;5;19m"
	ansiDarkGreen  = "\x1b[38;5;34m"
	ansiDarkAqua   = "\x1b[38;5;37m"
	ansiDarkRed    = "\x1b[38;5;124m"
	ansiDarkPurple = "\x1b[38;5;127m"
	ansiGold       = "\x1b[38;5;214m"
	ansiGrey       = "\x1b[38;5;145m"
	ansiDarkGrey   = "\x1b[38;5;59m"
	ansiBlue       = "\x1b[38;5;63m"
	ansiGreen      = "\x1b[38;5;83m"
	ansiAqua       = "\x1b[38;5;87m"
	ansiRed        = "\x1b[38;5;203m"
	ansiPurple     = "\x1b[38;5;207m"
	ansiYellow     = "\x1b[38;5;227m"
	ansiWhite      = "\x1b[38;5;231m"
	ansiDarkYellow = "\x1b[38;5;226m"

	ansiObfuscated    = ""
	ansiBold          = "\x1b[1m"
	ansiStrikethrough = "\x1b[9m"
	ansiUnderline     = "\x1b[4m"
	ansiItalic        = "\x1b[3m"
	ansiReset         = "\x1b[m"
)

var m = map[string]string{
	black:         ansiBlack,
	darkBlue:      ansiDarkBlue,
	darkGreen:     ansiDarkGreen,
	darkAqua:      ansiDarkAqua,
	darkRed:       ansiDarkRed,
	darkPurple:    ansiDarkPurple,
	gold:          ansiGold,
	grey:          ansiGrey,
	darkGrey:      ansiDarkGrey,
	blue:          ansiBlue,
	green:         ansiGreen,
	aqua:          ansiAqua,
	red:           ansiRed,
	purple:        ansiPurple,
	yellow:        ansiYellow,
	white:         ansiWhite,
	darkYellow:    ansiDarkYellow,
	obfuscated:    ansiObfuscated,
	bold:          ansiBold,
	strikethrough: ansiStrikethrough,
	underline:     ansiUnderline,
	reset:         ansiReset,
	italic:        ansiItalic,
}

var strMap = map[string]string{
	"black":         black,
	"dark-blue":     darkBlue,
	"dark-green":    darkGreen,
	"dark-aqua":     darkAqua,
	"dark-red":      darkRed,
	"dark-purple":   darkPurple,
	"gold":          gold,
	"grey":          grey,
	"dark-grey":     darkGrey,
	"blue":          darkBlue,
	"green":         green,
	"aqua":          aqua,
	"red":           red,
	"purple":        purple,
	"yellow":        yellow,
	"white":         white,
	"dark-yellow":   darkYellow,
	"obfuscated":    obfuscated,
	"bold":          bold,
	"strikethrough": strikethrough,
	"underline":     underline,
	"italic":        italic,
}

// minecraftReplacer and ansiReplacer are used to translate ANSI formatting codes to Minecraft formatting
// codes and vice versa.
var minecraftReplacer, ansiReplacer *strings.Replacer

func init() {
	var minecraftToANSI []string
	var ansiToMinecraft []string
	for minecraftCode, ansiCode := range m {
		if ansiCode != ansiObfuscated {
			ansiToMinecraft = append(ansiToMinecraft, ansiCode, minecraftCode)
		}
		minecraftToANSI = append(minecraftToANSI, minecraftCode, ansiCode)
	}
	minecraftReplacer, ansiReplacer = strings.NewReplacer(minecraftToANSI...), strings.NewReplacer(ansiToMinecraft...)
}
