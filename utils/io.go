package utils

import (
	"io"
	"log"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}
