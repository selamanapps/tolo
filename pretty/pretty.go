package pretty

import "fmt"

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	dim   = "\033[2m"

	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[90m"
)

func ResetString() string {
	return reset
}

func DimString(msg string) string {
	return dim + msg + reset
}

func BoldString(msg string) string {
	return bold + msg + reset
}

func CyanString(msg string) string {
	return cyan + msg + reset
}

func GreenString(msg string) string {
	return green + msg + reset
}

func RedString(msg string) string {
	return red + msg + reset
}

func YellowString(msg string) string {
	return yellow + msg + reset
}

func MagentaString(msg string) string {
	return magenta + msg + reset
}

var icons = struct {
	Success string
	Error   string
	Warning string
	Info    string
	Running string
	Saved   string
	Deleted string
	Updated string
	List    string
	Search  string
}{
	Success: "✓",
	Error:   "✗",
	Warning: "⚠",
	Info:    "ℹ",
	Running: "▶",
	Saved:   "💾",
	Deleted: "🗑",
	Updated: "🔄",
	List:    "📋",
	Search:  "🔍",
}

func Success(msg string) {
	fmt.Printf("%s%s %s%s %s\n", green, icons.Success, reset, bold, msg)
}

func Error(msg string) {
	fmt.Printf("%s%s %s%s %s\n", red, icons.Error, reset, bold, msg)
}

func Warning(msg string) {
	fmt.Printf("%s%s %s%s %s\n", yellow, icons.Warning, reset, bold, msg)
}

func Info(msg string) {
	fmt.Printf("%s%s %s%s %s\n", blue, icons.Info, reset, bold, msg)
}

func Running(msg string) {
	fmt.Printf("%s%s %s%s %s\n", cyan, icons.Running, reset, bold, msg)
}

func Saved(msg string) {
	fmt.Printf("%s%s %s%s %s\n", green, icons.Saved, reset, bold, msg)
}

func Deleted(msg string) {
	fmt.Printf("%s%s %s%s %s\n", red, icons.Deleted, reset, bold, msg)
}

func Updated(msg string) {
	fmt.Printf("%s%s %s%s %s\n", blue, icons.Updated, reset, bold, msg)
}

func List(msg string) {
	fmt.Printf("%s%s %s%s %s\n", magenta, icons.List, reset, bold, msg)
}

func Search(msg string) {
	fmt.Printf("%s%s %s%s %s\n", cyan, icons.Search, reset, bold, msg)
}

func Header(msg string) {
	fmt.Printf("\n%s%s═══ %s %s%s\n\n", bold, magenta, msg, reset, bold)
}

func Separator() {
	fmt.Printf("%s%s%s\n", gray, "─", reset)
}

func Dim(msg string) {
	fmt.Printf("%s%s%s\n", dim, msg, reset)
}

func Label(msg string) {
	fmt.Printf("%s%s%s", dim, msg, reset)
}

func Value(msg string) {
	fmt.Printf("%s%s%s\n", cyan, msg, reset)
}

func Command(msg string) {
	fmt.Printf("%s%s%s\n", green, msg, reset)
}

func Alias(name string) {
	fmt.Printf("%s%s%s", bold, name, reset)
}

func Count(count int) {
	fmt.Printf("%s%s%d%s\n", bold, cyan, count, reset)
}

func Newline() {
	fmt.Println()
}
