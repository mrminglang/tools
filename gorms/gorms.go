package gorms

import (
	"errors"
	"fmt"
	"strings"
	"time"

	strings2 "github.com/mrminglang/tools/strings"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/rogger"
	"gorm.io/gorm"
)

// CheckAndCreateDatabase 检查数据库是否存在，不存在则创建
func CheckAndCreateDatabase(db *gorm.DB, dbName string, logger *rogger.Logger) error {
	var exists int64
	// 查询数据库是否存在
	db.Raw("SELECT COUNT(*) FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&exists)

	// 数据库不存在，创建数据库
	if exists == 0 {
		createSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", dbName)
		if err := db.Exec(createSQL).Error; err != nil {
			logger.Errorf("{CheckAndCreateDatabase dbName::%s error::%s}", dbName, err.Error())
			return err
		}
		logger.Infof("{CheckAndCreateDatabase dbName::%s created successfully}", dbName)
	} else {
		logger.Infof("{CheckAndCreateDatabase dbName::%s already exists}", dbName)
	}

	return nil
}

// CreateBaseTable 创建基础表
func CreateBaseTable(db *gorm.DB, tableName string, dst interface{}, logger *rogger.Logger) {
	_ = MigratorCreateTable(db, tableName, dst, logger)
	logger.Infof("{CreateBaseTable tableName}|%s", tableName)
}

// CreateCurDayTable 创建当前分表
func CreateCurDayTable(db *gorm.DB, tableName, mod string, dst interface{}, logger *rogger.Logger) {
	tableName = tableName + strings2.GetTaskTableSuffix("", mod, false)
	_ = MigratorCreateTable(db, tableName, dst, logger)
	logger.Infof("{CreateCurDayTable tableName}|%s", tableName)
}

// CreateFrontDayTable 创建后期分表
func CreateFrontDayTable(db *gorm.DB, tableName, mod string, dst interface{}, logger *rogger.Logger) {
	tableName = tableName + strings2.GetTaskTableSuffix("", mod, true)
	_ = MigratorCreateTable(db, tableName, dst, logger)
	logger.Infof("{createFrontDayTable tableName}|%s", tableName)
}

// MigratorCreateTable 表迁移
func MigratorCreateTable(db *gorm.DB, tableName string, dst interface{}, logger *rogger.Logger) (err error) {
	if !db.Migrator().HasTable(tableName) {
		err = db.Table(tableName).Migrator().CreateTable(dst)
		if err != nil {
			logger.Errorf("{MigratorCreateTable tableName::%s error::%s}", tableName, err.Error())
			return err
		}
	}

	return
}

// SaveDataAfterCreateTable 保存表记录并尝试建表
func SaveDataAfterCreateTable[T any](tableName string, db *gorm.DB, data *T, dst interface{}, retry int32, logger *rogger.Logger) (err error) {
	if retry == 0 {
		retry = 5
	}

	// 使用循环替代递归
	for i := int32(0); i <= retry; i++ {
		// 递归重试最多5次
		if i == retry {
			logger.Errorf("{saveData2Db tableName::%s retry::%d data::%+v 超过重试次数}", tableName, i, data)
			return errors.New("超过重试次数")
		}

		//直接尝试写入数据
		if err = db.Table(tableName).Save(data).Error; err != nil {
			if strings.Contains(err.Error(), "for key") {
				logger.Errorf("{saveData2Db tableName::%s retry::%d data::%+v 保存表记录失败 error::%s}", tableName, i, data, err.Error())
				break
			}
			logger.Errorf("{saveData2Db tableName::%s retry::%d data::%+v 保存表记录失败 error::%s}", tableName, i, data, err.Error())
			// 表不存在的错误
			if strings.Contains(err.Error(), "doesn't exist") {
				if err = MigratorCreateTable(db, tableName, dst, logger); err != nil {
					logger.Errorf("{saveData2Db tableName::%s retry::%d 自动创建表失败 error::%s}", tableName, i, err.Error())
				}
			}

			// 指数退避策略
			time.Sleep(time.Duration(1<<i) * time.Second)
		} else {
			logger.Infof("{saveData2Db tableName::%s retry::%d data::%+v 保存表记录成功}", tableName, i, data)
			return
		}
	}

	return
}

// InsertDataAfterCreateTable 创建表记录并尝试建表
func InsertDataAfterCreateTable[T any](tableName string, db *gorm.DB, data *T, dst interface{}, retry int32, logger *rogger.Logger) (err error) {
	if retry == 0 {
		retry = 5
	}

	// 使用循环替代递归
	for i := int32(0); i <= retry; i++ {
		// 如果是最后一次重试，返回错误
		if i == retry {
			logger.Errorf("{insertData2Db tableName::%s retry::%d data::%+v 超过重试次数}", tableName, i, data)
			return errors.New("超过重试次数")
		}

		//直接尝试写入数据
		if err = db.Table(tableName).Create(data).Error; err != nil {
			if strings.Contains(err.Error(), "for key") {
				logger.Errorf("{insertData2Db tableName::%s retry::%d data::%+v 创建表记录失败 error::%s}", tableName, i, data, err.Error())
				break
			}
			logger.Errorf("{insertData2Db tableName::%s retry::%d data::%+v 创建表记录失败 error::%s}", tableName, i, data, err.Error())
			// 表不存在的错误
			if strings.Contains(err.Error(), "doesn't exist") {
				if err = MigratorCreateTable(db, tableName, dst, logger); err != nil {
					logger.Errorf("{insertData2Db tableName::%s retry::%d 自动创建表失败 error::%s}", tableName, i, err.Error())
				}
			}

			// 指数退避策略
			time.Sleep(time.Duration(1<<i) * time.Second)
		} else {
			logger.Infof("{insertData2Db tableName::%s retry::%d data::%+v 创建表记录成功}", tableName, i, data)
			return
		}

	}
	return
}

// InsertInBatchesAfterCreateTable 批量创建表记录并尝试建表
func InsertInBatchesAfterCreateTable[T any](tableName string, db *gorm.DB, data *[]T, dst interface{}, retry, batchSize int32, logger *rogger.Logger) (err error) {
	if batchSize == 0 {
		batchSize = 100
	}
	if retry == 0 {
		retry = 5
	}

	// 使用循环替代递归
	for i := int32(0); i <= retry; i++ {
		// 如果是最后一次重试，返回错误
		if i == retry {
			logger.Errorf("{insertInBatches2Db tableName::%s retry::%d data::%+v 超过重试次数}", tableName, i, data)
			return errors.New("超过重试次数")
		}

		// 直接尝试写入数据
		if err = db.Table(tableName).CreateInBatches(data, int(batchSize)).Error; err != nil {
			if strings.Contains(err.Error(), "for key") {
				logger.Errorf("{insertInBatches2Db tableName::%s retry::%d data::%+v 批量创建表记录失败 error::%s}", tableName, i, data, err.Error())
				break
			}
			logger.Errorf("{insertInBatches2Db tableName::%s retry::%d len(data)::%d 批量创建表记录失败 error::%s}", tableName, i, len(*data), err.Error())
			// 表不存在的错误
			if strings.Contains(err.Error(), "doesn't exist") {
				if err = MigratorCreateTable(db, tableName, dst, logger); err != nil {
					logger.Errorf("{insertInBatches2Db tableName::%s retry::%d 自动创建表失败 error::%s}", tableName, i, err.Error())
				}
			}

			// 指数退避策略
			time.Sleep(time.Duration(1<<i) * time.Second)
		} else {
			logger.Infof("{insertInBatches2Db tableName::%s retry::%d len(data)::%d 批量创建表记录成功}", tableName, i, len(*data))
			return
		}
	}

	return
}

// RetryExecSql 重试SQL语句
func RetryExecSql(db *gorm.DB, sql string, retry int32, logger *rogger.Logger) (err error) {
	if retry == 0 {
		retry = 5
	}
	// 使用循环替代递归
	for i := int32(0); i <= retry; i++ {
		// 如果是最后一次重试，返回错误
		if i == retry {
			logger.Errorf("{RetryExecSql sql::%s retry::%d 超过重试次数}", sql, retry)
			return errors.New("超过重试次数")
		}

		if err = db.Exec(sql).Error; err != nil {
			logger.Errorf("{RetryExecSql sql::%s retry::%d 执行语句失败 err::%s}", sql, retry, err.Error())
			time.Sleep(time.Duration(1<<i) * time.Second)
		} else {
			logger.Infof("{RetryExecSql sql::%s retry::%d 执行语句成功}", sql, retry)
			return
		}
	}

	return
}

// ExecSql 执行SQL
func ExecSql(db *gorm.DB, sql string, logger *rogger.Logger) bool {
	result := db.Exec(sql)
	if result.Error != nil {
		logger.Errorf("{ExecSql sql^error}|%s|%s", sql, result.Error.Error())
		return false
	}
	if result.RowsAffected == 0 {
		logger.Errorf("{ExecSql sql^RowsAffected}|%s|%d", sql, result.RowsAffected)
		return false
	}

	return true
}
