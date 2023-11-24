package rusprofileclient

type HttpClient interface {
	GetIdByInn(string) (string, error)
}
