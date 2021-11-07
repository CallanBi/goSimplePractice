package returnfunc

func DeferReturnFunc() (ret int) {
  defer func() {
    ret++
  }()
  return 1
}

