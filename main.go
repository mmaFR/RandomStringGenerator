package RandomStringGenerator

import (
	"math/rand"
	"time"
)

const LETTERS string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*/.=}])@^\_|([{#~&'`

type generator struct {
	initialized bool
	letters     string
	lenLetters  int
}

func (g *generator) init() {
	if !g.initialized {
		rand.Seed(time.Now().UnixNano())
		g.initialized = true
	}
}
func (g *generator) setLetters(Letters string) {
	g.letters = Letters
	g.lenLetters = len(Letters)
}
func (g *generator) Generate(length int) string {
	g.init()
	var buf []byte = make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = g.letters[rand.Intn(g.lenLetters)]
	}
	return string(buf)
}
func (g *generator) Read(p []byte) (n int, err error) {
	g.init()
	for i := 0; i < len(p); i++ {
		p[i] = g.letters[rand.Intn(g.lenLetters)]
	}
	return len(p), nil
}

func NewGenerator(letters ...string) *generator {
	var generator *generator = new(generator)
	if len(letters) > 0 {
		generator.setLetters(letters[0])
	} else {
		generator.setLetters(LETTERS)
	}
	return generator
}
