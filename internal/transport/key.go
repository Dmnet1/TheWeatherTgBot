package transport

var (
	APIKey      KeyAPIGetter
	Key, Answer string
)

type (
	API struct {
		Key string
	}
	KeyAPIGetter interface {
		getKey() string
	}
)

func (a *API) getKey() string {
	return a.Key
}

func GetAPIKey(key KeyAPIGetter) string {
	return key.getKey()
}
