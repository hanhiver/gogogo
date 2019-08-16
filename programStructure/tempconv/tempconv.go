// tempconv负责华氏温度和摄氏温度之间的转化

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsojuteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
