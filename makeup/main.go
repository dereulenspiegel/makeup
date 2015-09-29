package makeup

import "github.com/dereulenspiegel/makeup/json"

type Cosmetics interface {
	Prettify(data []byte) ([]byte, error)
	PrettifyFile(path string) ([]byte, error)
}

func GetCosmetics(in string) Cosmetics {
	switch in {
	case "json":
		return &json.JsonMakeup{}
	default:
		return nil
	}
}
