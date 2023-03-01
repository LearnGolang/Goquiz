func main() {
  var m sync.Mutex
  fmt.Print("A, ")
  m.Lock()

  go func() {
    time.Sleep(200 * time.Millisecond)
    m.Unlock()
  }()

  m.Lock()
  fmt.Print("B ")
}