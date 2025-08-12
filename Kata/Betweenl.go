package kata

func Between(a, b int) []int {
  if a >= b  {return []int{}};
  
  var ret []int
  for ; a <= b; a++ {
    ret = append(ret, a)
  }
  
  return ret
}
