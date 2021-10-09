package email

// Service ...
type Service interface {
	Send(sendRequest) (sendResponse, error)
}

// Email ...
type Email struct {
	To string `json:"to"`
}


// service
type service struct {
}

// NewService ...
func NewService() Service {
	return &service{}
}

// Send ...
func (svc *service) Send(request sendRequest) (sendResponse, error) {
	return sendResponse{Email: Email{To: ""}}, nil
}
