package services

import (
	"errors"
	//"fmt"

	//"fmt"
	//"log"
	//"meliQuasar/model"
	"meliQuasar/repository"
	"meliQuasar/util"
)

const MESSAGE_ERROR_MAX_LEN string = "the number of messages does not match the number of satellites"

func getMaxLength(totalSatellite int, ar ...[]string)(int, error){

	maxSize := 0
	sizeColl := len(ar)

	if sizeColl == 0 || sizeColl < totalSatellite{
		return  0, &util.Exception{
			StatusCode: 502,
			Err : errors.New(MESSAGE_ERROR_MAX_LEN),
		}
	}

	for m:=0; m <len(ar); m++{
		if (len(ar[m])>maxSize){
			maxSize = len(ar[m])
		}
	}
	return maxSize, nil
}

func GetMessage(messages ...[]string) ([]string, error){

	lstSat := repository.GetSatellites()

	maxSize, err := getMaxLength(len(lstSat), messages...)
	if err != nil{
		return []string{""}, err
	}

	messages = matchLegth(maxSize, messages...)

	var t int = 0
	for l:=0; l< len(messages);l++{
		t++
		if (l < len(messages)-1){
			messages[l+1] = combinedCollection(messages[l], messages[l+1])
			l++
		}else{
			messages[t] = combinedCollection(messages[l-1], messages[t])
		}	
	}
	return getUniqueValues(messages...),nil
}

func combinedCollection(col1, col2 []string) []string{
	var j int = 0

	for i:=0; i < len(col1); i++{
		if (j>= len(col2) || col1[i] != col2[i]){
			//missing = append(missing, res1[i])
			col2[i] = col1[i]
		}else{
			j=j+1
		}
	}
	return col2
}

func matchLegth(newSize int, ar ...[]string) [][]string{
	for m:=0; m <len(ar); m++{
		if (len(ar[m])< newSize){
			ar[m] = append(ar[m], make([]string, newSize-len(ar[m]))... ) 
		}
	}
	return ar
}

func getUniqueValues(ar ...[]string)[]string{
	a:=make([]string, 1)

	for x:= 0; x < len(ar); x++{
		a = append(a, ar[x]...)
	}

	uniqueValues := make(map[string]bool)

	for _, value := range a {
		if len(value)>0{
			uniqueValues[value] = true
		}
        
    }

	keys := make([]string, 0, len(uniqueValues))
	for key := range uniqueValues {
		keys = append(keys, key)
	}
    return keys
}