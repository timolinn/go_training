package linkedin

type Linkedin struct {
	URL         string
	Name        string
	Connections int
}

// Feed returns the latest Linkedin posts
func (l *Linkedin) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, I just started a new position at Hotels.ng",
	}
}

// Fame tell how famous a user is. The higher
// the number of Connections the more famous
func (l *Linkedin) Fame() int {
	return l.Connections
}
