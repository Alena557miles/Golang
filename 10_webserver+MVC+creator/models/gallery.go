package models

type Gallery struct {
	Name string
	//Location string
	Artists []*Artist
}

func (g *Gallery) DeleteArtist(artist *Artist) []*Artist {
	for i, a := range g.Artists {
		if artist == a {
			g.Artists = append(g.Artists[:i], g.Artists[i+1:]...)
			break
		}
	}
	return g.Artists
}
