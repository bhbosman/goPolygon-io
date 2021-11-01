package internal

type IAppSettings interface {
	apply(setting *settings)
}
