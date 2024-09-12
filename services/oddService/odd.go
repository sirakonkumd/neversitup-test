package oddService

import "fmt"

func (s *oddService) FindOddService(arr []int) (OddServiceResponse, error) {
	var resp OddServiceResponse
	res, err := s.conf.GetOdd()
	if err != nil {
		// handle return err
		fmt.Println("GetOddDao ", err.Error())
		return resp, err
	}
	resp.Res = res
	return resp, nil
}

func FindOdd(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	countMap := make(map[int]int)

	for _, num := range arr {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	return 0
}
