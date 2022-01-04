package admin

import entity "Drop/DropAdminSupport/entities"

type AdminService interface {
	AddAcct(acct *Acct) (string, error)
	UpdateAcctStatus(acct *Acct, id string) (string, error)
	GetAllAcct(limit, skip int) ([]entity.AdminSupportDB, error)
	Login(cred Credentials) (string, error)
	GetAllShopAcct(limit, skip int) ([]entity.AdminSupportDB, error)
}

type Acct struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Status   string `bson:"status,omitempty" json:"status,omitempty"`
	Name     string `bson:"name" json:"name,omitempty"`
	Role     string `bson:"role" json:"role,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
	Type     string `bson:"type" json:"type,omitempty"`
}
type Credentials struct {
	Email    string `bson:"email" json:"email,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
}
