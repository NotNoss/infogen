package app

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/NotNoss/infogen/util"
)

func Email() {
	config := util.Config{}
	cfg := config.GetDomain()

	wordOne := util.GetWords()
	wordTwo := util.GetWords()
	numbers := rand.Intn(1000)

	alias := fmt.Sprintf("%s_%s%v@%s", wordOne, wordTwo, numbers, cfg)

	fmt.Println(alias)
	os.Exit(0)
}
