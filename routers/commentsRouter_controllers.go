package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bugu/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bugu/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bugu/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bugu/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bugu/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bugu/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:UserController"] = append(beego.GlobalControllerRouter["bugu/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bugu/controllers:UserController"] = append(beego.GlobalControllerRouter["bugu/controllers:UserController"],
		beego.ControllerComments{
				Method: "Login",
				Router: `/login`,
				AllowHTTPMethods: []string{"get"},
				MethodParams: param.Make(),
				Params: nil})

}
