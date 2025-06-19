package option

var (
	repo = make(map[*Option]string)
)

func AddToRepo(key *Option, value string) {
	repo[key] = value
}

func AddMapToRepo(data map[*Option]string) {
	for k, v := range data {
		repo[k] = v
	}
}

func GetRepo() map[*Option]string {
	return repo
}
