package port

type ErrorsPort interface {
	OutputValidationError(error)
	Output400Error(error)
	Output403Error(error)
	Output404Error(error)
	Output500Error(error)
}
