package interfaces

type ControllerApi interface {
	CreatePr() (string, error)
}
