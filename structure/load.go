package structure

import "github.com/maseology/goHGS/readers"

func Read(directoryPrefix string) *HGS {
	coords, nps, epl, nsl := readers.ReadHGSstructure(directoryPrefix + "o.coordinates_pm")
	exr, ne, epl2 := readers.GetElementsPM(directoryPrefix+"o.elements_pm", nsl-1)

	nxr := make([][]int, len(coords))
	for eid, nids := range exr {
		for _, nid := range nids {
			nxr[nid] = append(nxr[nid], eid)
		}
	}

	if ne != epl*(nsl-1) {
		panic("goHGS Read() 1")
	}
	if epl != epl2 {
		panic("goHGS Read() 2")
	}
	h := &HGS{
		Nn:   nps * nsl,
		Ne:   epl * (nsl - 1),
		Nps:  nps,
		Epl:  epl,
		Nsl:  nsl,
		Nly:  nsl - 1,
		Nxyz: coords,
		Exr:  exr,
		Nxr:  nxr,
	}
	return h
}
