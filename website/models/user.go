package models

type Login struct {
	Username string `placeholder:"Usuario" title:"Ingrese su usuario"`
	Password string `placeholder:"Contraseña" title:"Ingrese su contraseña"`
}

type User struct {
	Id          string `html:"hidden"`
	Name        string `placeholder:"Nombre"`
	Area        string
	AreaName    string
	AccessLevel string
}
