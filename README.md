# buts (bounded, unique, timeout stack)

This is an implementation of a bounded, unique, timeout stack, which means that this is a simple Stack with the following properties:
- Limited Bounds (e.g., it has a capacity which will not be exceeded, overflowing items will be discarded from the bottom of the stack)
- Unique, items can of any kind but can't be contained n>1 times
- Timeout, items have a limited lifetime in the stack. Items timed out will be removed from the stack

## Installation

Simply add this repository by executing `go get github.com/lakrizz/buts` in your shell. It should automatically appear in your `go.mod`-file.

## Example Code

```golang
import (
	"log"
	"time"

	"github.com/lakrizz/buts"
)

func main() {
	b, err := buts.NewBoundedTimeoutStack(1*time.Minute, 5)
	if err != nil {
		log.Fatal(err)
	}
	b.Push(5)
	log.Println(b.GetItemsSlice())
}

```

## Future Ideas
- Ensure thread safety
- Make use of generics to keep users from typecasting all the time
- Precision is currently off by about +-1.1ms (due to computation time; i'm not sure how to keep track of it right now)