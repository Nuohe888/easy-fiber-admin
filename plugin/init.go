package plugin

import "easy-fiber-admin/plugin/internal"

func Init() {
	registerStorage(StorageTypeLocal, func() IStorage {
		return &internal.StorageLocalImpl{}
	})

	registerStorage(StorageTypeMinio, func() IStorage {
		return &internal.StorageMinioImpl{}
	})

	// 注册密码加密插件
	registerCrypto(CryptoTypeMD5, func() ICrypto {
		return &internal.CryptoMD5Impl{}
	})

	registerCrypto(CryptoTypeSHA1, func() ICrypto {
		return &internal.CryptoSHA1Impl{}
	})

	registerCrypto(CryptoTypeSHA256, func() ICrypto {
		return &internal.CryptoSHA256Impl{}
	})

	registerCrypto(CryptoTypeSHA512, func() ICrypto {
		return &internal.CryptoSHA512Impl{}
	})
}
