/****************************************
http://go-tour-kr.appspot.com/#23

함수와 루프의 사용법을 익히는 간단한 연습으로,
제곱근 함수를 뉴턴의 방법(Newton's method)을 이용하여 구현합니다.
여기서 뉴턴의 방법이란 초기값 z를 선택한 후에 다음의 공식을 이용하여 반복적으로 Sqrt(x)
함수의 근사값을 찾아가는 방법을 말합니다:
z = z - (z * z - x) / (2 * z)
*/
package main
// math.Sqrt 사용용도
import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  for {
    rtn := z - (z * z - x) / (2 * z)
    if float32(rtn) == float32(z) {
      return rtn
    } else {
      z = rtn
    }
  }
}

func main() {
  fmt.Println(Sqrt(35))
  fmt.Println(math.Sqrt(35))
}
