package structure

import "sort"

func (h *HGS) BuildElementalConnectivity(cardinalOnly bool) map[int][]int {
	o := make(map[int][]int, h.Ne)

	unique := func(intSlice []int) []int { // https://www.golangprograms.com/remove-duplicate-values-from-slice.html
		keys := make(map[int]bool)
		list := []int{}
		for _, entry := range intSlice {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		return list
	}

	if !cardinalOnly {
		// version 1: all adjacent elements, no ordering
		for eid, nids := range h.Exr {
			var l []int
			for _, nid := range nids {
				l = append(l, h.Nxr[nid]...)
			}
			u := unique(l)
			o[eid] = make([]int, 0, len(u)-1)
			for _, uu := range u {
				if uu != eid {
					o[eid] = append(o[eid], uu)
				}
			}
		}
	} else {
		// version 2: cardinal elements, ordered: [laterals]-bottom-top
		for eid, nids := range h.Exr {
			d := make(map[int]int)
			for _, nid := range nids {
				for _, eid2 := range h.Nxr[nid] {
					d[eid2] += 1
				}
			}
			delete(d, eid)
			for k, v := range d {
				if v < 3 {
					delete(d, k)
				}
			}
			u := make([]int, 0, len(d))
			for k := range d {
				u = append(u, k)
			}

			// ordering [laterals]-bottom-top
			sort.Ints(u)
			hastop := eid+h.Epl == u[len(u)-1]
			hasbot := eid-h.Epl == u[0]
			if hasbot && hastop {
				u = append(u[1:len(u)-1], u[0], u[len(u)-1])
			} else if hasbot {
				u = append(u[1:], u[0], -1)
			} else if hastop { // top only
				u = append(u[:len(u)-1], -1, u[len(u)-1])
			} else {
				panic("shouldn't occur unless this is a 1-layered model")
			}
			o[eid] = u
		}
	}
	return o
}
