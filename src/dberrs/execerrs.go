package dberrs

func DB_EX01() AxaErr{
	return AxaErr{`(axa err: db-ex01) executioner: execution command's format is faulty
	? execution command must start with one of these:
    	* "feed" (example: feed in users roma:{password:romaspassword,email:roma@email.com})
		* "fetch" (example: fetch from users roma)
    	* "delete" (example: delete from users roma)
	! axa exec failed, command unexecuted`}
}

func DB_EX02(statement string) AxaErr{
	return AxaErr{`(axa err: db-ex02) executioner: command contains multiple "` + statement + `" statements
	? ... ` + statement + ` ... ` + statement + ` ...
	?           ^^^^ x
	? ... ` + statement + ` ... at ...
	?           ^^^^ v
	! axa exec failed, command unexecuted`}
}

func DB_EX03() AxaErr{
	return AxaErr{`(axa err: db-ex03) executioner: could not find collection
	? command must be of the following format
	? cmd in <collection> at ...
	! axa exec failed, command unexecuted`}
}

func DB_EX04() AxaErr{
	return AxaErr{`(axa err: db-ex04) executioner: two "at" statements back-to-back
	? ... at at ...
	?    ^^^^^^^ x
	? ... at ... at ...
	?    ^^^^^^^^^^^ v
	! axa exec failed, command unexecuted`}
}