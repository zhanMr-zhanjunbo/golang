package dao
var (
	// MgrObj 全局的Session管理对象（大仓库）
	MgrObj Manager
)
func InitMgr(name string,addr string,option...string){
	switch name{
	case "memory"://初始化一个内存版管理者
		//MgrObj=NewMemory()
	case "redis":
		MgrObj=NewRedisManager()
	}
	MgrObj.Init(addr,option...)//初始化mgr
}
