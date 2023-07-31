package module

type AlamatKrim struct {
	Prov string
	Kota string
}

type User struct {
	Name   string
	Alamat AlamatKrim
}

type IfUser interface {
	GetKota() string
	GetProv() string
	GetName() string
}

func CreateUser(name, prv, kota string) *User {
	return &User{
		Name: name,
		Alamat: AlamatKrim{
			Prov: prv,
			Kota: kota,
		},
	}
}

func (u User) GetKota() string {
	return u.Alamat.Kota
}

func (u User) GetProv() string {
	return u.Alamat.Prov
}

func (u User) GetName() string {
	return u.Name
}
