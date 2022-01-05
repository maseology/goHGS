package structure

type HGS struct {
	Nn, Ne, Nsl, Nly, Nps, Epl int          // number of nodes, number of elements, number of slices, number of layers, nodes per slice, elements per layer
	Coord                      [][3]float64 // node coordinates
}
