package lean

type leanClient struct {
	appId, appKey, masterKey string
	useSign                  bool
	Installation, User, Role *Collection
	File                     *Collection
}

//create a new lean client.
//You only need one client in a go application
//application will not check your keys if you don't call API that needs it
func NewClient(appId, appKey, masterKey string) *leanClient {

	ret := &leanClient{
		appId:     appId,
		appKey:    appKey,
		masterKey: masterKey,
	}

	installation := ret.Collection("_Installation")
	ret.Installation = &installation
	ret.Installation.classSubfix = "/installaions"

	user := ret.Collection("_User")
	ret.User = &(user)
	ret.User.classSubfix = "/users"

	role := ret.Collection("_Role")
	ret.Role = &(role)
	ret.Role.classSubfix = "/roles"

	file := ret.Collection("_File")
	ret.File = &(role)
	ret.File.classSubfix = "/files"

	return ret
}

const (
	//v1.1 api classes url base
	//all url will not end with /,should append it by yourself
	UrlBase         = "https://api.leancloud.cn/1.1"
	ClasssesUrlBase = UrlBase + "/classes"
)
