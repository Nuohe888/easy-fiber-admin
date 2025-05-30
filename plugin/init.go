package plugin

import "easy-fiber-admin/plugin/internal"

func Init() {
	RegisterStorage(StorageTypeLocal, func() IStorage {
		return &internal.StorageLocalImpl{}
	})

	RegisterStorage(StorageTypeMinio, func() IStorage {
		return &internal.StorageMinioImpl{}
	})
}
