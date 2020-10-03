package strutils

var favorites []string

func init() {
	favorites = make([]string, 3)
	favorites[0] = "got.com"
	favorites[1] = "aaa.com"

}

// Add sipal
func Add(favorite string) {
	favorites = append(favorites, favorite)
}

// GetAll gela
func GetAll() []string {
	return favorites
}
