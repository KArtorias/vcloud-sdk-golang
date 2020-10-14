package main

import (
	"encoding/json"
	"fmt"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/vod"
)

func updateVideoInfo() {

}

func GetVideoInfos() {

}

func getRecommendedPosters() {

}

func updateVideoPlayStatus() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	req := vod.NewUpdateVideoPublishStatusRequest()
	req.SetVid("vidxxxx")
	req.SetStatus("Published")
	resp, code, _ := instance.UpdateVideoPublishStatus("vidxxxxx", "Published")
	fmt.Println(code)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func main() {
	updateVideoPlayStatus()
}
