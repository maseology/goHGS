package structure

import "github.com/maseology/goHGS/readers"

func Read(directoryPrefix string) *HGS {
	coords, nps, epl, nsl := readers.ReadHGSstructure(directoryPrefix + "o.coordinates_pm")
	return &HGS{
		Nn:    nps * nsl,
		Ne:    epl * (nsl - 1),
		Nps:   nps,
		Epl:   epl,
		Nsl:   nsl,
		Nly:   nsl - 1,
		Coord: coords,
	}
}
