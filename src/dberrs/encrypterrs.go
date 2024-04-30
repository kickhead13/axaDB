package dberrs

func DB_E01() AxaErr{
	return AxaErr{`(axa err : db-e01) encryption process: encryption failed
	? could not encrypt data
	? check for the existence of your environment variables
	? for more info try executing "axa" in your terminal of choice 
	! axa failed...`}
}