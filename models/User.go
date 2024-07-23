// @Author scy
// @Time 2024/7/23 13:38
// @File User.go

package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/chyshen/ginblog/utils/gcode"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	UserId    uint         `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username  string       `json:"username" gorm:"type:VARCHAR(100);not null;comment:'用户名'" binding:"required,min=4,max=100" label:"用户名" example:"test"`
	Password  string       `json:"password" gorm:"type:VARCHAR(64);not null;comment:'密码'" binding:"required,min=6,max=64" label:"密码" example:"123456"`
	Role      uint8        `json:"role" gorm:"type:TINYINT UNSIGNED;not null;default:1;comment:'角色'" binding:"required" label:"角色" example:"3"`
	CreatedAt time.Time    `json:"created_at" gorm:"comment:'创建时间'"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"comment:'更新时间'"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index;comment:'删除时间'"`
}

type UserAddModel struct {
	Username string `json:"username" gorm:"type:VARCHAR(100);not null;comment:'用户名'" binding:"required,min=4,max=100" label:"用户名" example:"test"`
	Password string `json:"password" gorm:"type:VARCHAR(64);not null;comment:'密码'" binding:"required,min=6,max=64" label:"密码" example:"123456"`
	Role     uint8  `json:"role" gorm:"type:TINYINT UNSIGNED;not null;default:1;comment:'角色'" binding:"required" label:"角色" example:"3"`
}

// Md5Pwd 生成加密密码
func Md5Pwd(password string) string {
	// 创建md5算法
	hash := md5.New()
	// 写入需要加密的数据, "X!d3a@" 加盐
	_, err := hash.Write([]byte(password + "X!d3a@"))
	if err != nil {
		panic(err)
	}
	// 获取hash字符切片，Sum函数接受一个字符切片，切片的内容会原样追加到password加密后的hash值的前面，不需要可直接传入nil
	sum := hash.Sum(nil)
	// byte转为string
	fmt.Println("hex.EncodeToString(sum): ", hex.EncodeToString(sum))
	return hex.EncodeToString(sum)
}

// ScryptPwd 使用scrypt算法生成加密密码
func ScryptPwd(password string) string {
	// 加盐
	salt := []byte("Aszx!@123")
	hashPwd, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(hashPwd)
}

// BeforeCreate gorm钩子函数，会在创建对象前执行
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	//u.Password = Md5Pwd(u.Password)
	u.Password = ScryptPwd(u.Password)
	//fmt.Printf("加密后的密码：{ %v }, 长度为：{ %v }\n", u.Password, len(u.Password))
	return
}

// CheckUser 验证用户是否存在
func (u *UserAddModel) CheckUser(username string) int {
	var user User
	tx := DB.Where("username = ?", username).First(&user)
	if tx.RowsAffected > 0 {
		return gcode.ErrorUsernameUsed
	}
	return gcode.Success
}

// UserAdd 添加用户
func UserAdd(useradd *UserAddModel) int {
	//  提交前将密码加密，也可使用BeforeCreate钩子函数
	//useradd.Password = Md5Pwd(useradd.Password)
	//useradd.Password = ScryptPwd(useradd.Password)
	user := User{
		Username: useradd.Username,
		Password: useradd.Password,
		Role:     useradd.Role,
	}
	err := DB.Create(&user).Error
	if err != nil {
		return gcode.Error
	}
	return gcode.Success
}

type UserQueryModel struct {
	Username string `json:"username" binding:"required,min=4,max=50" label:"用户名" example:"test"`
	Role     uint8  `json:"role" binding:"required" label:"角色" example:"3"`
}

// UserQuery 查询单个用户
func UserQuery(id int) (UserQueryModel, int) {
	var user User
	tx := DB.Where("user_id = ?", id).First(&user)
	result := UserQueryModel{
		Username: user.Username,
		Role:     user.Role,
	}
	fmt.Println(result)
	if tx.RowsAffected > 0 {
		return result, gcode.Success
	}
	return result, gcode.ErrorUserNotExist
}

// UserList 用户列表
// 分页
// username 用户名
// pageSize 每页显示的条目数
// pageNum 页码
func UserList(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	// username 不为空
	if username != "" {
		// Where("username LIKE ?", "%"+username+"%") 模糊查询，只要查询字段中包含username，即符合条件
		DB.Select("id", "username", "role", "created_at").Where("username LIKE ?", "%"+username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		DB.Model(&users).Where("username LIKE ?", "%"+username+"%").Count(&total)
		return users, total
	}
	// username 为空
	DB.Select("id", "username", "role", "created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	DB.Model(&users).Count(&total)
	return users, total
}

type UserUpdateModel struct {
	Username string `json:"username" binding:"required,min=4,max=50" label:"用户名" example:"test"`
	Role     uint8  `json:"role" binding:"required" label:"角色" example:"3"`
}

// CheckUpUser 更新用户时调用此方法验证用户是否存在
func (u *UserUpdateModel) CheckUpUser(id int, username string) int {
	var user User
	// 通过username查询出id和username字段
	DB.Select("user_id, username").Where("username = ?", username).First(&user)
	//fmt.Println("user.ID: ", user.ID, "uint(id): ", uint(id))
	// user.ID为username用户id，如：user.ID=1，uint(id)为需要修改的id，如：uint(id)=6
	if user.UserId == uint(id) {
		return gcode.Success
	}
	// 已存在，tx.RowsAffected > 0 通过First()后 tx.RowsAffected = 1
	if user.UserId > 0 {
		return gcode.ErrorUsernameUsed //1001
	}
	// 不存在
	return gcode.Success
}

// UserUpdate 更新用户
func UserUpdate(id int, data *UserUpdateModel) int {
	user := User{
		Username: data.Username,
		Role:     data.Role,
	}
	err := DB.Model(&user).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return gcode.Error
	}
	return gcode.Success
}

// BeforeUpdate gorm钩子函数，会在更新对象前执行	 修改密码前调用此钩子函数不生效
//func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
//	fmt.Println("@@@@@@@@", u.Password)
//	u.Password = BcryptPwd(u.Password)
//	return
//}

type UserPasswordModel struct {
	Password string `json:"password" binding:"required,min=6,max=18" label:"密码" example:"123456"`
}

// UserChangePassword 修改用户密码
func UserChangePassword(id int, data *UserPasswordModel) (int, string) {
	var user User
	data.Password = ScryptPwd(data.Password)
	tx := DB.Model(&user).Where("id = ?", id).Updates(&data)
	if tx.Error != nil {
		return gcode.Error, ""
	}
	return gcode.Success, data.Password
}

// UserDel 删除用户
func UserDel(id int) int {
	// 如果模型（User）中包含了gorm.CreatedAt字段，会自动获得软删除
	// 调用Delete()，gorm不会从数据库中删除数据，只是在DeleteAt中设置当前时间，后面一般的查询方法是无法查找到该记录
	err := DB.Where("user_id = ?", id).Delete(&User{}).Error
	if err != nil {
		return gcode.Error
	}
	return gcode.Success
}

type UserLoginModel struct {
	Username string `json:"username" binding:"required,min=4,max=50" label:"用户名" example:"admin"`
	Password string `json:"password" binding:"required,min=6,max=18" label:"密码" example:"123456"`
}

// UserLogin 用户登录
func UserLogin(data *UserLoginModel) (UserLoginModel, int) {
	var user User
	tx := DB.Where("username = ?", data.Username).First(&user)
	if tx.Error != nil {
		return *data, gcode.ErrorUserNotExist
	}
	// 用户无权限
	if user.Role != 1 {
		return *data, gcode.ErrorUserNoRight
	}
	//// md5验证用户密码是否正确
	//md5pwd := Md5Pwd(data.Password)
	//// 密码不正确
	//if md5pwd != user.Password {
	//	return *data, gcode.ERROR_PASSWORD_WRONG
	//}
	// md5验证用户密码是否正确
	scryptPwd := ScryptPwd(data.Password)
	// 密码不正确
	if scryptPwd != user.Password {
		return *data, gcode.ErrorPasswordWrong
	}
	//// bcrypt验证用户密码是否正确
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	//// 密码不正确
	//if err != nil {
	//	return *data, gcode.ErrorPasswordWrong
	//}
	return *data, gcode.Success
}
