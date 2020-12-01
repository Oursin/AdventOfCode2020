package lib

import "strconv"

func ConvToInt(content []string) ([]int, error) {
	var res []int
	for _, l := range content {
		i, err :=strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}