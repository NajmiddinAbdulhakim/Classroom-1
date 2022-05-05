var workArray [10]uint8
	var kaka [6]uint8
	var hg,gh uint8
	for i := 0; i < 10; i++{
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scan(&kaka[i])
	}
	for i := 0; i < 6; i++ {
		if i % 2 == 0 {
			hg = workArray[kaka[i]] 
			gh = workArray[kaka[i+1]]
		} else {
			workArray[kaka[i]] = hg
			workArray[kaka[i-1]] = gh	
		}
	}
	for _,val := range workArray {
		fmt.Print(val," ")
	}



