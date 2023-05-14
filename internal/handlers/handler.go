package handlers

type handler struct{}

func NewHandler() (string, *handler) {
	return schema.String(), &handler{}
}
