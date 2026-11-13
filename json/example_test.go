package json_test

import (
	"errors"
	"fmt"

	"zeroparse.dev/alpha-parsers/json"
	"zeroparse.dev/alpha-parsers/json/types"
)

func ExampleDecode() {
	v, err := json.Decode(`{"name": "zp", "tags": ["peg", "pratt"], "stars": 42}`)
	if err != nil {
		panic(err)
	}
	m := v.(map[string]any)
	fmt.Println(m["name"], m["tags"], m["stars"])
	// Output: zp [peg pratt] 42
}

func ExampleParseValue() {
	v, err := json.ParseValue(`{"a": [1, true], "b": "x\ty"}`)
	if err != nil {
		panic(err)
	}
	for _, m := range v.(*types.Object).Members {
		fmt.Printf("%q at %d-%d: %T\n", m.Key.Value(), m.Span.Start, m.Span.End, m.Value)
	}
	// Output:
	// "a" at 1-15: *types.Array
	// "b" at 17-28: *types.String
}

func ExampleParseValue_syntaxError() {
	_, err := json.ParseValue(`{"a": [1, 2,]}`)
	var se *json.SyntaxError
	if errors.As(err, &se) {
		fmt.Println(se.Line, se.Col, se.Message())
	}
	// Output: 1 13 syntax error
}

func ExampleValid() {
	fmt.Println(json.Valid(`[1, 2, 3]`), json.Valid(`[1, 2, 3,]`))
	// Output: true false
}
