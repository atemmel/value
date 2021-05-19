package main

import (
	"fmt"
	"github.com/atemmel/value"
)

func show(value value.Value) {
	if val := value.AsBool(); val != nil {
		fmt.Print(*val)
	} else if val := value.AsInt(); val != nil {
		fmt.Print(*val)
	} else if val := value.AsString(); val != nil {
		fmt.Print(*val)
	} else if val := value.AsStringArray(); val != nil {
		fmt.Print(val)
	} else if val := value.AsValueArray(); val != nil {
		print("[")
		for i := range val {
			show(val[i])
			if i + 1 != len(val) {
				print(", ")
			}
		}
		print("]")
	}
}

func main() {

	table := value.ValueTable()

	table["health"] = value.Int(100)
	table["name"] = value.String("Rex")
	table["isRunning"] = value.Bool(false)
	table["friends"] = value.StringArray([]string{
		"Henry",
		"Jessie",
		"Stacy",
	})
	table["other"] = value.Array([]value.Value{
		value.Int(3),
		value.Bool(true),
		value.String("red"),
	})

	for key, value := range table {
		fmt.Printf("The key %s holds the value(s) ", key)
		show(value)
		fmt.Println()
	}
}
