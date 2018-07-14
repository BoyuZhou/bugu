package controllers

// import (
// 	"github.com/astaxie/beego"

// 	"bugu/models"
// )

// type BaseController struct {
// 	beego.Controller
// 	controllerName string             //当前控制名称
// 	actionName     string             //当前action名称
// 	curUser        models.User //当前用户信息
// }

// func (c *BaseController) Prepare() {
// 	//附值
// 	c.controllerName, c.actionName = c.GetControllerAndAction()
// 	//从Session里获取数据 设置用户信息
// 	c.adapterUserInfo()
// }

// //从session获取user信息
// func (c *BaseController) adapterUserInfo() {
// 	a := c.GetSession("user")
// 	if a != nil {
// 		c.curUser = a.(models.User)
// 		c.Data["user"] = a
// 	}
// }

// //SetUser2Session 获取用户信息（包括资源UrlFor）保存至Session
// func (c *BaseController) setUser2Session(userId int) error {
// 	m, err := models.GetUserById(userId)
// 	if err != nil {
// 		return err
// 	}
// 	//获取这个用户能获取到的所有资源列表
// 	// resourceList := models.ResourceTreeGridByUserId(userId, 1000)
// 	// for _, item := range resourceList {
// 	// 	m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
// 	// }
// 	c.SetSession("user", *m)
// 	return nil
// }

// // func (c *BaseController) jsonResult(code utils.JsonResultCode, msg string, obj interface{}) {
// // 	r := &models.JsonResult{code, msg, obj}
// // 	c.Data["json"] = r
// // 	c.ServeJSON()
// // 	c.StopRun()
// // }

// // 重定向
// func (c *BaseController) redirect(url string) {
// 	c.Redirect(url, 302)
// 	c.StopRun()
// }

// // 重定向 去错误页
// func (c *BaseController) pageError(msg string) {
// 	errorurl := c.URLFor("HomeController.Error") + "/" + msg
// 	c.Redirect(errorurl, 302)
// 	c.StopRun()
// }

// // 重定向 去登录页
// func (c *BaseController) pageLogin() {
// 	url := c.URLFor("HomeController.Login")
// 	c.Redirect(url, 302)
// 	c.StopRun()
// }
