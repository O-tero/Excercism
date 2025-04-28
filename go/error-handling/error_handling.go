package erratum

import (
	"fmt"
)

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	// Retry logic for TransientError
	for {
		res, err = opener()
		if err == nil {
			break
		}

		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	// Ensure we close the resource even if there are panics or errors
	defer func() {
		if res != nil {
			res.Close()
		}
	}()

	// Recover from panic if it is a FrobError
	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				res.Defrob(frobErr.defrobTag)
				err = frobErr
			} else {
				if rErr, ok := r.(error); ok {
					err = rErr
				} else {
					err = fmt.Errorf("unexpected panic: %v", r)
				}
			}
		}
	}()

	// Call the Frob function
	res.Frob(input)

	return nil
}
