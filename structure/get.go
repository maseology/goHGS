package structure

import "github.com/maseology/goHGS/readers"

func (hgs *HGS) ReadElementalVectors(fp string) (float64, [][3]float32) {
	return readers.ReadHGSvector(fp, hgs.Nps, hgs.Epl, hgs.Nsl)
}
