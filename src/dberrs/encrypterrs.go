package dberrs

func DB_E01() AxaErr{
	return AxaErr{`(axa err : db-e01) encryption process: encryption failed
	? could not encrypt data
	! axa failed...`}
}