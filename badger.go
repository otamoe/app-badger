package appbadger

import (
	"fmt"
	"os"
	"path"

	"github.com/dgraph-io/badger/v3"
	appconfig "github.com/otamoe/app-config"
	applogger "github.com/otamoe/app-logger"
	"github.com/shirou/gopsutil/v3/mem"
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	appconfig.SetDefault("badger.indexDir", path.Join(homeDir, "."+appconfig.GetName(), "badger", "index"), "Badger index dir")
	appconfig.SetDefault("badger.valueDir", path.Join(homeDir, "."+appconfig.GetName(), "badger", "value"), "Badger index dir")
}

var DB *badger.DB

func GetDB() *badger.DB {
	return DB
}

func SetDB(v *badger.DB) {
	DB = v
}

func Close() error {
	return DB.Close()
}

func DefaultOptions() badger.Options {
	memorySize := GetMemorySize()
	return badger.DefaultOptions(appconfig.GetString("badger.indexDir")).
		WithValueDir(appconfig.GetString("badger.valueDir")).
		WithBaseTableSize(1024 * 1024 * 8).
		WithMemTableSize(int64(memorySize / 32)).
		WithValueThreshold(1024 * 1).
		WithBlockCacheSize(int64(memorySize / 32)).
		WithIndexCacheSize(int64(memorySize / 32)).
		WithLogger(NewLogger(applogger.GetLogger()))
}

func GetMemorySize() uint64 {
	// 读取内存
	memStat, err := mem.VirtualMemory()
	if err != nil {
		panic(fmt.Errorf("get memory size", err))
	}
	return memStat.Total
}
