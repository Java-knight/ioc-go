package autowire

// 结构性描述符(SD)Map
var structDescriptorsMap = make(map[string]*StructDescriptor)

// 注册 SD
func RegisterStructDescriptor(sdID string, descriptor *StructDescriptor) {
	structDescriptorsMap[sdID] = descriptor
}

// 根据 sdID 获取 SD
func GetStructDescriptor(sdID string) *StructDescriptor {
	return structDescriptorsMap[sdID]
}

// 获取 SDMap
func GetStructDescriptorsMap() map[string]*StructDescriptor {
	return structDescriptorsMap
}
