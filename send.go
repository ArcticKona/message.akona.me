package message

func ( self Message )Send( ) [ ]error {
	var err [ 16 ]error
	err[ 0 ] = self.Text( )
	err[ 1 ] = self.Email( )
	for _ , err := range err {
		if err == nil {
			return [ ]error{ } } }
	return err
}


