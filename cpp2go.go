package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var g_isTest = flag.Bool("t", false, "is testing")
var g_isListHint = flag.Bool("n", false, "is listing hint numbers")

var g_hints = []string{
	`semicolons omit and multiple line rule
  func main() <-- compile error
  {
    ...
  }
  func main() { <-- OK
    ...
  }
  if x {
  }       <-- compile error
  else {
  }
  s := "line1 \  <-- compile error
  line2"
  s := "line1"  <-- compile error
   + "line2"
  s := "line1"+  <-- OK
   "line2"

  http://weekly.golang.org/doc/go_for_cpp_programmers.html#Syntax
  http://golang.org/doc/go_spec.html#Semicolons
`,
	"multi line string use backquote" +
		"\n  `line 1" +
		"\n  line 2`" +
		"\n  for string include backquote use strings.Replace(`.X..`,\"X\",\"'\",-1)" +
		"\n  or use \"line 1\"+" +
		"\n         \"line 2\"",
	`type cast
  convert float to []byte http://weekly.golang.org/pkg/encoding/binary/#Write
`,
	`array declaration
  http://weekly.golang.org/doc/go_for_cpp_programmers.html#Syntax
  for multiple dimension use following trick for speedup
  FIXME
`,
}
var g_dic = map[string]string{
	// Reserved Words http://cs.smu.ca/~porter/csc/ref/cpp_keywords.html
	"auto":     "X",
	"break":    "break [Label] to outer loop",
	"case":     "case",
	"char":     "int8/uint8/byte",
	"const":    "only for define constant variable in global/function",
	"continue": "continue [Label] to outer loop",
	"default":  "default",
	"do":       "http://weekly.golang.org/doc/go_spec.html#For_statements",
	"double":   "float64",
	"else":     "else",
	"enum": `const+iota
  const (
    a = 1 << iota  // a == 1 (iota has been reset)
    b = 1 << iota  // b == 2
    c = 1 << iota  // c == 4
  )
`,
	"extern":   "all global declaration which name lead with uppercase",
	"float":    "float32",
	"for":      "http://weekly.golang.org/doc/go_spec.html#For_statements",
	"goto":     "goto",
	"if":       "if",
	"int":      "Go's int is unlimited digits, there are int8,uint16,...",
	"long":     "int32 for 4 bytes, int64 for 8 bytes",
	"register": "X",
	"return":   "could return multiple values",
	"short":    "int16",
	"signed":   "use uint,uint8,uint16,...",
	"sizeof":   "unsafe.Sizeof",
	"static":   "X",
	"struct":   "http://weekly.golang.org/doc/go_spec.html#Struct_types",
	"switch": `there is no default fall through rule
  switch i {  // i=1/2/3 will only produce one line output
  case 1:
  case 2:
    println("It is", i)
  case 3:
    println("It's three")
  }
  is equal to
  switch i {
  case 1, 2:
    println("It is", i)
  case 3:
    println("It's three")
  }
  Here is the code with same logic as C/C++ code
  switch i {
  case 1: case 2: println("It is", i); fallthrough
  case 3: println("It's three")
  }
  more http://weekly.golang.org/doc/go_spec.html#Switch_statements
 `,
	"typedef":  "http://weekly.golang.org/doc/go_spec.html#Types",
	"union":    "X",
	"unsigned": "uint8,uint16,...",
	"void":     "X",
	"volatile": "X",
	"while":    "http://weekly.golang.org/doc/go_spec.html#For_statements",

	// C++ Reserved words
	"asm":              "try to import C, http://weekly.golang.org/cmd/cgo/",
	"dynamic_cast":     "X",
	"namespace":        "X",
	"reinterpret_cast": "X",
	"try":              "http://weekly.golang.org/doc/go_spec.html#Handling_panics",
	"bool":             "bool",
	"explicit":         "X",
	"new":              "new/make http://weekly.golang.org/doc/go_spec.html#Allocation",
	"static_cast":      "X",
	"typeid":           "X",
	"catch":            "http://weekly.golang.org/doc/go_spec.html#Handling_panics",
	"false":            "false",
	"operator":         "X",
	"template":         "X",
	"typename":         "X",
	"class":            "interface http://weekly.golang.org/doc/go_for_cpp_programmers.html#Interfaces",
	"friend":           "X",
	"private":          "name's first char prefix with lowercase",
	"this":             "X",
	"using":            "X",
	"const_cast":       "X",
	"inline":           "X",
	"public":           "name's first char prefix with uppercase",
	"throw":            "X",
	"virtual":          "X",
	"delete":           "automatic garbage collection, delete(m,k) for del map item",
	"mutable":          "X",
	"protected":        "X",
	"true":             "true",
	"wchar_t":          "rune",

	// some predefined identifiers
	"cin":      "os.Stdin",
	"cout":     "os.Stdout",
	"endl":     "use Println(), xxxln(), or \"\\n\"",
	"include":  "import, see also \"go help importpath\"",
	"INT_MIN":  "",
	"INT_MAX":  "",
	"iomanip":  "godoc fmt",
	"iostream": "godoc fmt",
	"main": `http://weekly.golang.org/doc/go_spec.html#Packages
To return error code to OS, use os.Exit() instead of return from main()
`,
	"MAX_RAND": "X",
	"NULL":     "nil",
	"string": `maintain string by package strings/strconv"
  http://weekly.golang.org/doc/go_spec.html#String_types
`,

	// common programing behavior
	"argv": "os.Args",
	"argc": "len(os.Args",

	// common functions
	"getopt": "godoc flag",
	"printf": `fmt.Println/Print/Printf
Go's Print is powerful, it could print any type with String() member.
For xxxf() function, we can use "go tool vet *.go" to validate % flags
`,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`Usage: cpp2go [options] <keyword_in_c/c++>
This utility is to let you quick map concept of C/C++ to Go.
eg.
  cpp2go sprintf explicit # search multiple keywords
  cpp2go 12 # list No.12 hint
  cpp2go -h # list option help
Reference:
  http://weekly.golang.org/doc/go_for_cpp_programmers.html
  Go syntax          http://golang.org/doc/go_spec.html
  Go packages        http://golang.org/pkg/
  Test code on web   http://play.golang.org/
  3rd party packages
    http://godashboard.appspot.com/package
    http://godashboard.appspot.com/project
`)
		os.Exit(1)
	}
	flag.Parse()
	if *g_isTest {
		fmt.Println("Err: TODO")
		os.Exit(1)
	}
	if *g_isListHint {
		showHintList()
	}
	for _, s := range flag.Args() {
		n, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			fmt.Println(s, "->", g_dic[s])
		} else if int(n) >= len(g_hints) {
			fmt.Println("Err: over max index, use -n option to list index")
		} else {
			fmt.Print(s, ":", g_hints[n], "\n")
		}
	}
}
func showHintList() {
	for i, s := range g_hints {
		lines := strings.SplitN(s, "\n", 2)
		fmt.Printf("%3d %s\n", i, lines[0])
	}
}
