package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//1. 主键的约定：GORM 默认会使用名为ID的字段作为表的主键。
//2. 表名默认就是结构体名称的复数
//3. 列名由字段名称进行下划线分割来生成,可以使用结构体tag指定列名：
//4. 如果模型有 CreatedAt字段，该字段的值将会是初次创建记录的时间。
//5. 如果模型有UpdatedAt字段，该字段的值将会是每次更新记录的时间。
//6. 如果模型有DeletedAt字段，调用Delete删除该记录时，将会设置DeletedAt字段为当前时间，而不是直接将记录从数据库中删除。

type User struct {
	gorm.Model   //  Model包含了一些必需的字段，如主键之类的
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` //下面分别是结构体标记
	Role         string  `gorm:"size:255"`                       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段
}

// gorm 连接数据库
func main() {

	db, err := gorm.Open("mysql", "root:bin123456@(127.0.0.1:3306)/binshow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//1. 自动创建表，自动迁移，把数据库和数据表进行对应
	db.AutoMigrate(&User{})

}
