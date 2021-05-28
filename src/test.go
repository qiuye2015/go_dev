package main

import (
	"encoding/json"
	"fmt"
)

type RedisRes struct {
	Result []struct {
		Uuid string  `json:"uuid"`
		Ctr  float32 `json:"click_rate"`
	} `json:"result"`
}

func main() {
	var redisRes RedisRes
	res := "{\"result\":[{\"uuid\":\"2c24f79ebe753b58bc51a46018640bbe\",\"click_rate\":0.0743},{\"uuid\":\"75a3a7ee598c32cf86f3741244acfe7c\",\"click_rate\":0.01},{\"uuid\":\"1c10b53284dc3ce8915a3f6f7e4238d8\",\"click_rate\":0.0097},{\"uuid\":\"e339bfb1650c31b1ba233581bb8f4e59\",\"click_rate\":0.0086},{\"uuid\":\"8a0e1b54cde23927b74c6d0375267b03\",\"click_rate\":0.0085},{\"uuid\":\"8a92609be8cf39318547d87321133cd6\",\"click_rate\":0.008},{\"uuid\":\"5319c80611f436b9a9847a4a15c70162\",\"click_rate\":0.0075},{\"uuid\":\"170a2c876f1b3765bb2db746b230ca7e\",\"click_rate\":0.0075},{\"uuid\":\"451b57246e5d32d7ae5020a773040cce\",\"click_rate\":0.0074},{\"uuid\":\"c43140ce8cf235b5987991775069d581\",\"click_rate\":0.0065},{\"uuid\":\"951fc0d90db63a67b6277f480dfd9b6a\",\"click_rate\":0.0057},{\"uuid\":\"fe3c8baf1c4f3218a700123ac2e7fb00\",\"click_rate\":0.0055},{\"uuid\":\"7dbfc0e76c633676ad24537415d058c9\",\"click_rate\":0.005},{\"uuid\":\"efb4aea9d01231caa46720cb3c869d10\",\"click_rate\":0.0048},{\"uuid\":\"b2375f641c1337ce8e12ca735cfe8275\",\"click_rate\":0.0047},{\"uuid\":\"7aa5fef9c04c3ff9a04c882834ea297c\",\"click_rate\":0.0047},{\"uuid\":\"cead818f092c33f78a7217f81a8b87a7\",\"click_rate\":0.0042},{\"uuid\":\"2a8d966cff923ba4b77728f696f25f62\",\"click_rate\":0.0042},{\"uuid\":\"3861db056a22338798bf3dff3a2bc598\",\"click_rate\":0.0042},{\"uuid\":\"cecd28a7c6ab33faafe25e1122fee33a\",\"click_rate\":0.0041},{\"uuid\":\"9735d0a6911a3f45bc1d257e3940da02\",\"click_rate\":0.0041},{\"uuid\":\"0b794d15a17336c690c86dddd57418a4\",\"click_rate\":0.004},{\"uuid\":\"c5978b1ae12c35599b1015f53f712f0b\",\"click_rate\":0.004},{\"uuid\":\"1028fdd1bf8334aaa860ac457fb8f780\",\"click_rate\":0.004},{\"uuid\":\"eba32be3275b346786b0c2252a825c74\",\"click_rate\":0.0038},{\"uuid\":\"e43b655f905038328d2abfb99ab1798b\",\"click_rate\":0.0038},{\"uuid\":\"3232ab496f9638b08d63a71bb8f0b3a6\",\"click_rate\":0.0036},{\"uuid\":\"16b1029218ac3f11bea15b6034e792bd\",\"click_rate\":0.0035},{\"uuid\":\"d63888b6662c35938e78e84720496a97\",\"click_rate\":0.0032},{\"uuid\":\"8d9b13fb062c3244b01868b8ef202f5d\",\"click_rate\":0.0031}]}"
	err := json.Unmarshal([]byte(res), &redisRes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(redisRes)
	for _, value := range redisRes.Result {
		fmt.Println(value.Uuid, value.Ctr)
	}

}
