package lean

type leanClient struct {
	appId, appKey, masterKey string
	useSign                  bool
}

//create a new lean client.
//You only need one client in a go application
//application will not check your keys if you don't call API that needs it
func NewClient(appId, appKey, masterKey string) *leanClient {
	return &leanClient{
		appId:     appId,
		appKey:    appKey,
		masterKey: masterKey,
	}
}

const (
	//v1.1 api classes url base
	//all url will not end with /,should append it by yourself
	UrlBase         = "https://api.leancloud.cn/1.1"
	ClasssesUrlBase = UrlBase + "/classes"
)
