package outsideserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"songLibrary/internal/outsideserver/dto"
)

type OutsideServer struct {
	bingAddress string
}

func New(bingAddress string) *OutsideServer {
	return &OutsideServer{
		bingAddress: bingAddress,
	}
}
func (outsideserver *OutsideServer) GetInfo(req *dto.GetInfoReq) (*dto.GetInfoRes, error) {
	reqUrl, err := url.Parse(outsideserver.bingAddress + "/info")
	if err != nil {
		return nil, fmt.Errorf("outsideServer: %s", err)
	}
	params := url.Values{}
	params.Add("group", req.Group)
	params.Add("song", req.Song)
	reqUrl.RawQuery = params.Encode()
	resHttp, err := http.Get(reqUrl.String())
	if err != nil {
		return nil, fmt.Errorf("outsideServer: %s", err)
	}
	defer resHttp.Body.Close()

	if resHttp.StatusCode != 200 {
		return nil, fmt.Errorf("outsideServer: %s", resHttp.Status)
	}
	body, err := io.ReadAll(resHttp.Body)
	if err != nil {
		return nil, fmt.Errorf("outsideServer: %s", err)
	}
	res := dto.GetInfoRes{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("outsideServer: %s", err)
	}
	return &res, nil
}
