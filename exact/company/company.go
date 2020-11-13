package company

// Company ...
type Company struct {
	CNPJ string `json:"cnpj" bson:"cnpj"`

	Info []CompanyInfo `json:"info" bson:"info"`
}

// CompanyInfo ...
type CompanyInfo struct {
	DataType string `json:"datatype" bson:"datatype"`
	Value    string `json:"value" bson:"value"`
	IsValid  bool   `json:"isvalid" bson:"isvalid"`
}

// New ...
func New(cnpj string) *Company {
	return &Company{
		CNPJ: cnpj,
		Info: []CompanyInfo{},
	}
}
