package geodb

type Coordinates struct {
	Latitude  Degree `json:"lat"`
	Longitude Degree `json:"lon"`
}

func (c *Coordinates) Location() Location {
	return MakeLocation(c.Latitude, c.Longitude)
}
