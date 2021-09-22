package thirdparty

// 用户管理的能力
// 例如 同步过来的用户，我们可能需要能修改他的信息，例如企业微信中的通讯录，可以修改员工信息
// ** 目前还用不到该功能，所以暂时不实现 **
type UserManager interface {
	// 三方拥有管理用户的能力
	CanManager() bool
}
