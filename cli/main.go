package main

import (
	"github.com/zmoog/lzss"
)

func main()  {
	
	lzss.Encode("lorem.txt", "lorem.lzss")
}