package main

import (
	"github.com/dtherhtun/Learning-go/design-pattern/threelayerapp/applayer"
	"github.com/dtherhtun/Learning-go/design-pattern/threelayerapp/httplayer"
	"github.com/dtherhtun/Learning-go/design-pattern/threelayerapp/storelayer"
)

func main() {
	storeLayer := storelayer.New()
	appLayer := applayer.New(storeLayer)
	api := httplayer.New(appLayer)
	api.Engage()
}
