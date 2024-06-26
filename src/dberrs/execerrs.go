package dberrs

func DB_EX01() AxaErr{
	return AxaErr{`(axa err: db-ex01) executioner: execution command's format is faulty
  ? consult the documentation for execution commands: https://github.com/kickhead13/axaDB
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

func DB_EX05() AxaErr{
	return AxaErr{`(axa err: db-ex05) executioner: payload is not in json format
	? possible examples:
	?  * ... at ... {: "where is my key?"}
	?  * ... at ... { "where is" : "my right bracket?
	! axa exec failed, command unexecuted"`}
}
func DB_EX06() AxaErr{
	return AxaErr{`(axa err: db-ex06) executioner: bad feed request
	? go panicked, can't execute command
	! axa exec failed, command unexecuted`}
}
func DB_EX07() AxaErr{
	return AxaErr{`(axa err: db-ex07) executioner: cannot find data
	? data you are looking for doesn't seem to exist
	! axa exec failed`}
}
