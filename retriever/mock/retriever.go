package mock

type Retriever struct {
	Contents string
}

func (r *Retriever) Post(url string,
	forms map[string]string) string {
	r.Contents = forms["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

