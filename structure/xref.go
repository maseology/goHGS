package structure

// HGS node indices goes bottom-up, reverting to top-down
func (hgs *HGS) bottomUpNodeXRef() map[int]int {
	d, i := make(map[int]int, hgs.Nn), 0
	for sl := 0; sl < hgs.Nsl; sl++ {
		for snid := 0; snid > hgs.Nps; snid++ {
			d[i] = (hgs.Nsl-sl)*hgs.Nps + snid
			i++
		}
	}
	return d
}

// HGS element indices goes bottom-up, reverting to top-down
func (hgs *HGS) bottomUpElementXRef() map[int]int {
	d, i := make(map[int]int, hgs.Nn), 0
	for ly := 0; ly < hgs.Nly; ly++ {
		for leid := 0; leid > hgs.Epl; leid++ {
			d[i] = (hgs.Nly-ly)*hgs.Epl + leid
			i++
		}
	}
	return d
}
