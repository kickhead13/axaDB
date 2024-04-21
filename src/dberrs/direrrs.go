package dberrs

func DB_D01() AxaErr{
	return AxaErr{`(axa err : db-d01) database directory: already exists 
	? database directory must be created at database creation time
	? it must not exist prior to the creation of the database
	! axa init failed...`}
}

func DB_D02() AxaErr{
	return AxaErr{`(axa err : db-d02) database diectory: couldn't be created
	? system didn't grant permission for database directory creation
	! axa init failed...`}
}

func DB_D03() AxaErr{
	return AxaErr{`(axa err : db-d03) database directory: AXA_USERS couldn't be created
	? database directory must contain AXA_USERS collection
	? the collection couldn't be created successfuly
	! database directory might still be partially created on disk
	! axa init failed...`}
}

func DB_D04() AxaErr{
	return AxaErr{`(axa err : db-d04) database directory: could not create init.cfg file
	? system didn't grant permission for database init.cfg file creation
	! axa init failed...`}
}

func DB_D05() AxaErr{
	return AxaErr{`(axa err : db-d05) database directory: could not create AXA_USERS df.db datafile
	? system didn't grant permission for datafile creation
	! axa init failed...`}
}

func DB_D06(collection string) AxaErr{
	return AxaErr{`(axa err : db-d06) database directory: could not create rules file for ` + collection + ` collection
	? system didn't grant permission to create a collections rules file
	! axa init failed...`}
}

func DB_NORM() AxaErr{
	return AxaErr{``}
}
