package kata

func Digitize(n int) []int {
  if n == 0 {return []int{0}};

  result := []int{}
  
  for n > 0 {
    result = append(result, n%10)
    n /= 10
  }
  
  return result
}
