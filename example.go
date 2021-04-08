package main

import (
	"fmt"

	"github.com/maseology/goHGS/readers"
	"github.com/maseology/mmio"
)

const (
	idir = "M:/CVC/HGS/IC_model_run4/"
	prfx = "IC_model_run4"
	sufx = "o.%s_pm.%04d"
)

func main() {

	nps, epl, nsl := readers.ReadHGSstructure(idir + prfx + "o.coordinates_pm")
	fmt.Println(nps, epl, nsl)

	ts := 0
	mvals := map[float64][]float64{}
	for {
		ts++
		fp := fmt.Sprintf(idir+prfx+sufx, "head", ts)
		if _, ok := mmio.FileExists(fp); !ok {
			break
		}
		t, v := readers.ReadHGSscalar(fp, nps, epl, nsl)
		fmt.Println(fp, "  t=", t)
		mvals[t] = v
	}

	fmt.Println(mvals[604800.][:10])
}
