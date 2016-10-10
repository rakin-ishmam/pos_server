package apperr

// Format format the error with keys and return new error
// func Format(err error, keys ...string) error {
// 	str := "#"
// 	for i := 0; i < len(keys); i++ {
// 		if i == 0 {
// 			str += keys[i]
// 		} else {
// 			str += fmt.Sprintf("->%s", keys[i])
// 		}
// 	}
//
// 	return fmt.Errorf("%s:%s", str, err.Error())
// }
