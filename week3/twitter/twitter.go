package twitter

type Twitter struct {
	URL       string
	Username  string
	Followers int
}

// Feed returns the latest Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying other people's work",
	}
}

// Fame tell how famous a user is. The higher
// the number of followers the more famous
func (t *Twitter) Fame() int {
	return t.Followers
}
