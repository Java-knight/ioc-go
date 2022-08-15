package autowire

import (
	"fmt"
	"ioc-go/logger"
)

// 打印注册sd
func printAutowireRegisteredStructDescriptor() {
	for autowireType, autowire := range GetAllWrapperAutowires() {
		logger.Blue("[Autowire Type] Found registered autowire type %s", autowireType)
		for sdID := range autowire.GetAllStructDescriptors() {
			logger.Blue("[Autowire Struct Descriptor] Found type %s registered SD %s", autowireType, sdID)
		}
	}
}

// 加载
func Load() error {
	printAutowireRegisteredStructDescriptor()

	// 自动装配所有有入口的struct
	for _, autowire := range GetAllWrapperAutowires() {
		for sdID := range autowire.GetAllStructDescriptors() {
			if autowire.CanBeEntrance() {
				sd := GetStructDescriptor(sdID)
				if sd == nil {
					continue
				}
				// 无参数实现
				_, err := autowire.ImplWithoutParam(sdID, !sd.DisableProxy)
				if err != nil {
					fmt.Errorf("[Autowire] Impl sd %s failed, readson is %s", sdID, err)
				}
			}
		}
	}
	return nil
}

func Impl(autowireType, key string, param interface{}) (interface{}, error) {
	return impl(autowireType, key, param, false)
}

func ImplWithProxy(autowireType, key string, param interface{}) (interface{}, error) {
	return impl(autowireType, key, param, true)
}

// 根据autowireType和sdId、param 获取实现, expectWithProxy表示是否通过代理（AOP）
func impl(autowireType, key string, param interface{}, expectWithProxy bool) (interface{}, error) {
	targetSDID := GetSDIDByAliasIfNecessary(key)

	// check 代理标记
	sd := GetStructDescriptor(targetSDID)
	if sd != nil && sd.DisableProxy {
		// 如果sd禁用了代理，则将 expectWithProxy（代理标记）设置为 false
		expectWithProxy = false
	}

	// 遍历所有的wrapper后的autowire（可能经过proxy）
	for _, wrapperAutowire := range GetAllWrapperAutowires() {
		if wrapperAutowire.TagKey() == autowireType {
			return wrapperAutowire.ImplWithParam(targetSDID, param, expectWithProxy)
		}
	}
	logger.Red("[Autowire] SDID %s with autowire type %s not found in all autowire", key, autowireType)
	return nil, fmt.Errorf("[Autowire] SDID %s with autowire type %s not found in all autowire", key, autowireType)
}
