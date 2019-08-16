package tempconv

//把摄氏温度转为华氏温度
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 把华氏温度转为摄氏温度
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
