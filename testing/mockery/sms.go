package mockery

type SMS interface {
	Send(number int, text string) error
}

// Messager is a struct handling messaging of various types.
type Messager struct {
	sms SMS
}

// SendHelloWorld sends a Hello world SMS.
func (m *Messager) SendHelloWorld(number int) error {
	err := m.sms.Send(number, "Hello, world!")
	if err != nil {
		return err
	}
	return nil
}
