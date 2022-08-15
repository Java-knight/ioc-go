package autowire

// proxy function

var pf func(interface{}) interface{}

// 注册代理函数
func RegisterProxyFunction(f func(interface{}) interface{}) {
	pf = f
}

// 获取代理函数
func GetProxyFunction() func(interface{}) interface{} {
	if pf == nil {
		return func(i interface{}) interface{} {
			return i
		}
	}
	return pf
}
