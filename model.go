package platform

type Theme struct{}

type TemplateModuleConfig struct {
	RenderAllSpaceCentered bool // ej para login
	Module                 module
	// ej:	`<div class="target-module">
	// 	<select name="select">
	// 		<option value="value1">Value 1</option>
	// 		<option value="value2" selected>Value 2</option>
	// 	</select>
	// </div>`
	HeaderInputTarget string

	Form        object
	FormButtons []*ButtonForm

	AsideList containerViewAdapter
}

type ButtonForm struct {
	ModuleName string // ej:c.Module
	ObjectName string // ej: c.Form.ObjectName
	ButtonName string // ej: btn_printer
	Title      string // ej: Imprimir
	IconID     string // ej: icon-printer
	OnclickFun string // ej: printForm(this)
	Disabled   bool   // activado por defecto
}

// todo el contenido html por defecto del objeto
type containerViewAdapter interface {
	BuildContainerView(id, field_name string, allow_skip_completed bool) string
}

type module interface {
	//nombre modulo ej: chat,patient,user
	ModuleName() string
	//Titulo que vera el usuario ej: "Modulo Fotografía"
	Title() string
	// id icono para utilizar en sprite svg ej: icon-info
	IconID() string
	//interfaz usuario modulo
	UserInterface(u *User) string
}

type User struct {
	Token          string // token de sesión solicitante
	Id             string // id usuario
	Ip             string
	Name           string
	Area           string // valor de carácter ej: a,s,p... key "" sin area
	AreaName       string // nombre del area
	AccessLevel    string // aquí valor numérico 0 a 255
	LastConnection string //time.Time

	R any //r *http.Request.
	W any //w http.ResponseWriter
}

type object interface {
	ObjectName() string
	PrimaryKeyName() string
	RenderFields() []field
}

type field interface {
	Name() string
	Legend() string
	SkipCompletionAllowed() bool //se permite que el campo no este completado. por defecto siempre sera requerido
	Input() input
}

type input interface {
	HtmlName() string
	BuildContainerView(id, field_name string, allow_skip_completed bool) string
}
