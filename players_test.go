package gogamescore

func ListenAndServe(addr string, handler Handler) error

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}