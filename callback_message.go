package gov

type (
	CallbackMessage struct {
		FromStatus int
		ToStatus   int
		Payload    KriMessage
	}
)
