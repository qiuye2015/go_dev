package model

import "encoding/json"

//type Profile struct {
//	Name     string
//	Age      int
//	Marriage string
//}

//安康 | 31岁 | 大专 | 离异 | 162cm | 3001-5000元
type Profile struct {
	Name     string
	Age      int
	Sex      string
	Marriage string
	Url      string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
