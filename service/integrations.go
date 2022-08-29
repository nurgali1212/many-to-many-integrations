package service

func (s *Service) SendDataService(err error) error {
	return s.Client.SendData(err)
}
