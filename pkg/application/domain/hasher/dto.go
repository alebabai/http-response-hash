package hasher

type HashURLContentInput struct {
	URL string
}

type HashURLContentOutput struct {
	URL  string
	Sum  []byte
	Size int
}
