package main

import "fmt"

type filterService interface {
	run()
}

type MarketplaceFilterImpl struct {}

type AmountFilterImpl struct {}

func (service MarketplaceFilterImpl) run(){
	fmt.Println("Run MarketplaceFilter...")
}

func (service AmountFilterImpl) run(){
	fmt.Println("Run AmountFilter...")
}

func main(){

	var anything []filterService

	anything = append(anything, MarketplaceFilterImpl{})
	anything = append(anything, AmountFilterImpl{})

	for _, val := range anything{
		val.run()
	}

}
