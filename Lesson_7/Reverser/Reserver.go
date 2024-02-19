package Reverser

func RVS(user_input []interface{}) []interface{} {
	switch user_input[0].(type) {
	case int:
		for i, j := 0, len(user_input)-1; i < j; i, j = i+1, j-1 {
			user_input[i], user_input[j] = user_input[j], user_input[i]
		}
		return user_input
	case string:
		for i, j := 0, len(user_input)-1; i < j; i, j = i+1, j-1 {
			user_input[i], user_input[j] = user_input[j], user_input[i]
		}
		return user_input
	case float64:
		for i, j := 0, len(user_input)-1; i < j; i, j = i+1, j-1 {
			user_input[i], user_input[j] = user_input[j], user_input[i]
		}
		return user_input
	default:
		return user_input
	}
}
