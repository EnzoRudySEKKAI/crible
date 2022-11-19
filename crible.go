package main

func filter(nfiltre int, nbToTest chan int, chanCrea chan int) {
	last := true
	nextnNbToTest := make(chan int)
	for {
		nb := <-nbToTest
		if nb%nfiltre != 0 && last {
			chanCrea <- nb
			last = false
			go filter(nb, nextnNbToTest, chanCrea)
		} else if nb%nfiltre != 0 && !last {
			nextnNbToTest <- nb
		}
	}
}

func main() {
	nbToTest := make(chan int)
	chanCrea := make(chan int)
	premier := make([]int, 1)
	go filter(2, chanCrea, nbToTest)
	for i := 3; i < 10; i++ {
		nbToTest <- i
	}
	for {
		premier = append(premier, <-chanCrea)
	}
}
