package main

import (
	"fmt"
	"math/rand"

	"github.com/NotNoss/mail-alias/util"
)

func main() {
	util.CheckConfig()

	config := util.Config{}
	cfg := config.GetDomain()

	wordOne := util.GetWords()
	wordTwo := util.GetWords()
	numbers := rand.Intn(1000)

	alias := fmt.Sprintf("%s_%s%v@%s", wordOne, wordTwo, numbers, cfg)

	fmt.Println(alias)
}
