package models

type GlobalData struct {
	IsAuthenticated bool
	Data            any
	UserName        string
	Categories      []Category
}
