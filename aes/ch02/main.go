package main

func converArr(array [20]byte) []byte {
	s := []byte{}
	for _, elem := range array {
		s = append(s, elem)
	}
	return s
}

func main() {

	//md5 := md5.Sum([]byte("5d4b7a49339eb3191688116224871548Mysoft!@#$%^&"))
	//
	//sha1 := sha1.Sum([]byte("5d4b7a49339eb3191688116224871548Mysoft!@#$%^&"))
	//
	//key := GetByteArray(converArr(sha1), 32)
	//
	//iv := converArr(md5)

	//fmt.Println(realKeys)
	//
	//realKey, _ := strconv.ParseInt(hex.EncodeToString(GetByteArray([]byte{}, 32)), 16, 0)
	//
	//iv, _ := strconv.ParseInt(hex.EncodeToString(), 16, 0)
	//
	//fmt.Println(realKey)
	//
	//fmt.Println(iv)
}

func GetByteArray(src []byte, destLen int) []byte {
	dest := make([]byte, destLen)
	p := 0
	for p < destLen {
		for _, b := range src {
			if p >= destLen {
				return dest
			} else {
				dest[p] = b
				p++
			}
		}
	}
	return dest
}
