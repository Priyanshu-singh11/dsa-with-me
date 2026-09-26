
data := []int{6,8,18,527,69,8,1}
for i := 0; i < len(data); i++ {
  for j := i+1; j < len(data); j++ {
    if data[i] > data[j] {
      data[i], data[j] = data[j], data[i]
    }
  }
}
fmt.Println(data)
