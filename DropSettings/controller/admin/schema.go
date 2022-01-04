package admin

import entity "Drop/DropSettings/entities"

type AdminSettingService interface {
	AddSettings(setting Settings, token string) (string, error)
	GetSettingsHistory(token string, limit, skip int) ([]entity.SettingDB, error)
	GetCurrentSettings(token string, limit, skip int) ([]entity.SettingDB, error)
}

type Settings struct {
	Drop                     string        `bson:"drop" json:"drop"`
	CutType                  string        `bson:"cut_type" json:"cut_type"`
	DeliveryCharge           []interface{} `bson:"delivery_charge" json:"delivery_charge"`
	DeliveryPersonPercentage string        `bson:"delivery_person_percentage" json:"delivery_person_percentage"`
	UpdatedBy                string        `bson:"updated_by" json:"updated_by"`
}
