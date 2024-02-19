package SREZ

func RNS(srez []int) ([]int, []int) {
	chetnie := make([]int, 0)
	ne_chetnie := make([]int, 0)
	for _, v := range srez {
		if v%2 == 0 {
			chetnie = append(chetnie, v)
		} else {
			ne_chetnie = append(ne_chetnie, v)
		}
	}
	return chetnie, ne_chetnie
}
