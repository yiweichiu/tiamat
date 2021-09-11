package models

type LineProfile struct {
	UserId        string
	DisplayName   string
	PictureUrl    string
	StatusMessage string
	Language      string
}

func (lp *LineProfile) Name() string {
	return lp.DisplayName
}
