package exception

func PanicIfErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
