package ndfc

import (
	"log"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"
)

type NDFC struct {
	url       string
	apiClient nd.Client
	rscMutex  map[string]*sync.Mutex
}

var instance *NDFC

func NewNDFCClient(host string, user string, pass string, domain string, insecure bool) (*NDFC, error) {
	log.Printf("New NDFC client")
	ndfc := new(NDFC)
	ndfc.url = "/appcenter/cisco/ndfc/api/v1"
	var err error
	ndfc.apiClient, err = nd.NewClient(host, ndfc.url, user, pass, domain, insecure, nd.MaxRetries(500))
	if err != nil {
		return nil, err
	}
	ndfc.rscMutex = make(map[string]*sync.Mutex)
	instance = ndfc
	return ndfc, nil
}

func NewResource(rscName string) {
	log.Printf("New Resource %s\n", rscName)
	if _, ok := instance.rscMutex[rscName]; !ok {

		instance.rscMutex[rscName] = new(sync.Mutex)
	}
}

func (c NDFC) GetLock(rscName string) *sync.Mutex {
	return (instance.rscMutex[rscName])
}
