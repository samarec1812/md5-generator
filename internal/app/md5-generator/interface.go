package md5_generator


type Service interface {
	Run() error
}

type Services struct {
	PORT string
}
