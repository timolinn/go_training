package facebook

type Facebook struct {
	URL     string
	Name    string
	Friends int
}

// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

// Fame tell how famous a user is. The higher
// the number of friends the more famous
func (f *Facebook) Fame() int {
	return f.Friends
}
