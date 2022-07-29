package types

type EndPoint interface {
	RegisterEndPoint()
}

type EndPointManager struct {
	EndPoints []EndPoint
}

func NewEndPointManager(v ...EndPoint) *EndPointManager {
	em := &EndPointManager{EndPoints: v}
	return em
}

func(e *EndPointManager) RegisterEndPoint() {
	for _, endPoint := range e.EndPoints{
		endPoint.RegisterEndPoint()
	}
}
