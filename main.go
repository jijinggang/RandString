package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	//"time"
)

var (
	_prefix  = flag.String("pre", "", "prefix")
	_postfix = flag.String("post", "", "postfix")
	_seed    = flag.Int64("seed", 0, "random seed")

	_begin   = flag.Int("begin", 0, "random series' begin position")
	_count   = flag.Int("count", 10000, "generated strings' total count")
	_len     = flag.Int("len", 10, "generated strings' length")
	_letters = flag.String("letter", "abcdefghijklmnopqrstuvwxyz1234567890", "letters of generated string")
)

func main() {
	flag.Parse()
	seed := *_seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	end := *_begin + *_count
	for i := 0; i < end; i++ {
		if str := NewRand(*_len, *_letters); !IsExists(str) {
			if i >= *_begin {
				fmt.Println(*_prefix + str + *_postfix)
			}
		}
	}

}

var _m map[string]bool

func IsExists(str string) bool {
	if _m == nil {
		_m = make(map[string]bool)
	}
	if _, ok := _m[str]; ok {
		return true
	}
	_m[str] = true
	return false
}
func NewRand(length int, letters string) string {

	buffer := make([]rune, length)
	size := len(letters)
	for i := range buffer {
		buffer[i] = rune(letters[rand.Intn(size)])
	}
	return string(buffer)
}
