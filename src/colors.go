package src

import (
	"math/rand"
	"time"
)

const (
	RED = 0xff1f1f

	ORANGE = 0xffa21f

	YELLOW = 0xffff1f

	GREEN = 0x41ff1f

	BLUE = 0x1fdaff

	INDIGO = 0x391fff

	PURPLE = 0xce1fff

	BLACK = 0x000000

	WHITE = 0xffffff
)

func RandomColors() int {
	list := []int{RED, ORANGE, YELLOW, GREEN, BLUE, INDIGO, PURPLE}
	return list[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(list))]
}
