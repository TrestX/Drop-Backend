package general

import entity "Drop/DropUserAccount/entities"

type SettingService interface {
	SetSetting(setting *Setting, userid string) (string, error)
	GetSetting(userID string) (entity.UserDB, error)
}

type Setting struct {
	Theme    string `bson:"theme" json:"theme,omitempty"`
	Language string `bson:"language" json:"language,omitempty"`
}
