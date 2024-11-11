package models

type GlobalData struct {
	IsAuthenticated bool
	Data            any
	Categories      []Category
}
