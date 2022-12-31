package my_global_ip

import (
	"dns_updater/client/my_http_client"
	"dns_updater/logger"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
)

var globalIPURL = "https://domains.google.com/checkip"

type GlobalIPClient struct {
	client *my_http_client.MyHttpClient
	logger *logger.Logger
	req    *http.Request
}

func validateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func (g *GlobalIPClient) Get() (*string, error) {
	getGlobalIPLock.Lock()
	defer getGlobalIPLock.Unlock()

	if globalIP != nil {
		g.logger.Debug(fmt.Sprintf("global ip: %s was found", *globalIP))
		return globalIP, nil
	}

	g.logger.Debug("update called")
	g.logger.Info("get my global ip addr start")
	body, err := g.client.Get(g.req)
	if err != nil {
		g.logger.Error("get global ip addr failed")
		return nil, err
	}
	_body, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	ip := string(_body)
	fmt.Println(ip)
	if !validateIP(ip) {
		return nil, errors.New(fmt.Sprintf("global ip addr was invalid: %s", ip))
	}
	globalIP = &ip
	g.logger.Info(fmt.Sprintf("get my global ip was ended successfully: %s", *globalIP))
	return globalIP, nil
}

var (
	SharedGlobalIPClient  *GlobalIPClient
	newGlobalIPClientLock sync.Mutex
	getGlobalIPLock       sync.Mutex
	globalIP              *string //直接参照させず、Get経由で取得させる
)

//GlobalIPClientは共有で使用する
func NewGlobalIPClient(timeout int, logger *logger.Logger) *GlobalIPClient {
	newGlobalIPClientLock.Lock()
	defer newGlobalIPClientLock.Unlock()

	if SharedGlobalIPClient != nil {
		return SharedGlobalIPClient
	}

	globalIPLogger := logger.Child("GlobalIPClient")
	req, err := http.NewRequest("GET", globalIPURL, nil)
	if err != nil {
		logger.Error("create DDNSClient failed")
		logger.Error(err.Error())
		return nil
	}
	client := my_http_client.NewMyHttpClient(timeout, globalIPLogger)
	SharedGlobalIPClient := &GlobalIPClient{client: client, logger: globalIPLogger, req: req}
	return SharedGlobalIPClient
}
