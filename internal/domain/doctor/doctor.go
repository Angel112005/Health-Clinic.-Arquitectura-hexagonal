package doctor

type Doctor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Specialty  string `json:"specialty"`
	Experience int    `json:"experience"`
}
