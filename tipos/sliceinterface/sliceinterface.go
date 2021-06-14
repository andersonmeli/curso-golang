package main

import "fmt"

type filterService interface {
	run()
}

type MarketplaceFilterImpl struct {
	name string
}

type AmountFilterImpl struct {
	name string
}

type PayFilterImpl struct {
	name string
}

func (service MarketplaceFilterImpl) run(){
	fmt.Println("Run MarketplaceFilter...")
}

func (service AmountFilterImpl) run(){
	fmt.Println("Run AmountFilter...")
}

func (service PayFilterImpl) run(){
	fmt.Println("Run AmountFilter...")
}

func RemoveIndex(s []filterService, filter filterService) []filterService {
	index := -1
	for k, v := range s {
		if filter == v {
			index = k
		}
	}

	if index > -1 {
		return append(s[:index], s[index+1:]...)
	}else {
		return s
	}
}

func main(){

	var anything []filterService

	filter01 := MarketplaceFilterImpl{name: "MarketplaceFilterImpl"}
	filter02 := AmountFilterImpl{name: "AmountFilterImpl"}
	//filter03 := PayFilterImpl{name: "PayFilterImpl"}
	anything = append(anything, filter01)
	anything = append(anything, filter02)

	fmt.Println(anything)

	fmt.Println(RemoveIndex(anything, filter01))

	//for _, val := range anything{
	//	val.run()
	//}

}
