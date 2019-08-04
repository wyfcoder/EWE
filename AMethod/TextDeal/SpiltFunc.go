package TextDeal

//分割冒号
func SpiltSemicolon(rune rune) bool{
	if rune == ';'{
		return true
	}
	return false
}
//分隔逗号和空格
func SpiltSpace(rune rune) bool{
	if  rune == ','{
		return true
	}
	return false
}
