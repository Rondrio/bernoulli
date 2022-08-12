package bernoulli

import (
	"fmt"
	"math/rand"
	"time"
)

type bernoulli struct {
	funcs map[string]func(...any)
}

func New() *bernoulli {
	return &bernoulli{
		funcs: make(map[string]func(...any)),
	}
}

func (b bernoulli) AddFunc(name string, f func(...any)) {
	b.funcs[name] = f
}

func (b bernoulli) Exec(name string, params ...any) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	if n > 50 {
		b.funcs[name](params)
		return
	}
	fmt.Println("bernoulli said no to executing:", name, "\nbetter luck next time")
}
