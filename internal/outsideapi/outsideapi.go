package outsideapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"songLibrary/internal/outsideapi/dto"
	"songLibrary/internal/pkg/servererrors"

	log "github.com/sirupsen/logrus"
)

type OutsideApi struct {
	bingAddress string
}

func New(bingAddress string) *OutsideApi {
	return &OutsideApi{
		bingAddress: bingAddress,
	}
}
func (o *OutsideApi) GetInfo(req *dto.GetInfoReq) (*dto.GetInfoRes, error) {
	reqUrl, err := url.Parse(o.bingAddress + "/info")
	if err != nil {
		log.Errorf("outsideApi: getInfo: %s", err)
		return nil, err
	}
	params := url.Values{}
	params.Add("group", req.Group)
	params.Add("song", req.Song)
	reqUrl.RawQuery = params.Encode()
	resHttp, err := http.Get(reqUrl.String())
	if err != nil {
		log.Errorf("outsideApi: getInfo: %s", err)
		return nil, err
	}
	defer resHttp.Body.Close()

	if resHttp.StatusCode != 200 {
		err := servererrors.ErrorInvalidServerResponseStatus
		log.Errorf("outsideApi: getInfo: %s", err)
		return nil, err
	}
	body, err := io.ReadAll(resHttp.Body)
	if err != nil {
		log.Errorf("outsideApi: getInfo: %s", err)
		return nil, err
	}
	res := dto.GetInfoRes{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorf("outsideApi: getInfo: %s", err)
		return nil, err
	}
	return &res, nil
}
