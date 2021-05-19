# value

**A (moderately) generic value type implementation for Go**

## Usage

```go
import (
	"fmt"
	"github.com/atemmel/value"
)

func main() {

	// Creates a table of values
	table := value.ValueTable()

	// Each slot in the table may contain values of different types
	table["health"] = value.Int(100)
	table["name"] = value.String("Rex")
	table["isRunning"] = value.Bool(false)
	table["friends"] = value.StringArray([]string{
		"Henry",
		"Jessie",
		"Stacy",
	})

	// Values are extracted with the appropriate 'AsTypename' function
	// All 'AsTypename' functions always return nil upon type mismatch
	if val := table["isRunning"].AsBool(); val != nil {
		// If the value was not nil, the type has been succesfully matched
		if *val {
			fmt.Println("It is running!")
		} else {
			fmt.Println("It is not running...")
		}
	}

	show := func (value value.Value) {
		if val := value.AsBool(); val != nil {
			fmt.Print(*val)
		} else if val := value.AsInt(); val != nil {
			fmt.Print(*val)
		} else if val := value.AsString(); val != nil {
			fmt.Print(*val)
		} else if val := value.AsStringArray(); val != nil {
			fmt.Print(val)
		}
	}

	for key, value := range table {
		fmt.Printf("The key %s holds the value(s): ", key)
		show(value)
			fmt.Println()
	}

}

```

## Installation
```sh
go get github.com/atemmel/value
```

## Avilable types

* `int`
* `uint`
* `bool`
* `string`
* `float32`
* `float64`
* `byte`
* `[]int`
* `[]uint`
* `[]bool`
* `[]string`
* `[]float32`
* `[]float64`
* `[]byte`
* `[]Value` - Used to keep values of different types within the same container

## Ideal for

* Light scripting implementations
* Generic data manipulation

## Todo

* Serialization
* Array instantiations nil checks

## Design references

* [C++'s `dynamic_cast` conversion](https://en.cppreference.com/w/cpp/language/dynamic_cast)
* [Java's and JavaScript's `instanceof` operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/instanceof)
