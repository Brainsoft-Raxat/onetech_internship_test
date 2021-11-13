package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func() {
		for {
			num, ok := <-c.Input
			if !ok {
				break
			}
			c.Output <- num*num
		}
		close(c.Output)
	}()
}
