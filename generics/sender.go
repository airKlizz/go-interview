package generics

type Sender struct{}

func (s *Sender) SendInt(msg int) int {
	return msg
}

func (s *Sender) SendString(msg string) string {
	return msg
}
