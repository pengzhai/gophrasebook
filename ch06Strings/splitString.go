package main
import (
	"strings"
	"unicode"
	"fmt"
	"golang.org/x/exp/utf8string"
)

func main() {
	str := "\tthe important rôles of utf8 text\n"
	str = strings.TrimFunc(str, unicode.IsSpace)
	// the wrong way
	fmt.Printf("%s\n", str[0:len(str)/2])
	// the right way
	u8 := utf8string.NewString(str)
	FirstHalf := u8.Slice(0, u8.RuneCount()/2)
	fmt.Printf("%s\n", FirstHalf)
}
