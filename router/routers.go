/*
 * Demo App
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	api "tfnserver/go_api"
	log "tfnserver/log"

	"github.com/gorilla/mux"
)

const uploadPath = "./image"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		http.FileServer(http.Dir(uploadPath))))
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = log.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Xin chao!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/xuyendangvan/DemoAPI/1.0.0/",
		Index,
	},

	Route{
		"CalculateTuitionFee",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/admin/tuitionfee/{id}",
		api.CalculateTuitionFee,
	},

	Route{
		"GetFeedbacks",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/admin/feedback",
		api.GetFeedbacks,
	},

	Route{
		"AddClass",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/classes",
		api.AddClass,
	},

	Route{
		"DeleteClassByID",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/classes/{id}",
		api.DeleteClassByID,
	},

	Route{
		"FindScheduleByClassID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/classes/{id}/schedule",
		api.FindScheduleByClassID,
	},

	Route{
		"FindStudentByClassID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/classes/{id}/students",
		api.FindStudentByClassID,
	},

	Route{
		"FindClassByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/classes/{id}",
		api.FindClassByID,
	},

	Route{
		"FindClassListByIndex",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/classes",
		api.FindClassListByIndex,
	},

	Route{
		"UpdateClass",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/classes/{id}",
		api.UpdateClass,
	},

	Route{
		"CreateMenu",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/menu",
		api.CreateMenu,
	},

	Route{
		"DeleteMenuByID",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/menu/{id}",
		api.DeleteMenuByID,
	},

	Route{
		"GetMenuByDay",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/menu/day",
		api.GetMenuByDay,
	},

	Route{
		"GetMenuDetailForWeek",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/menu/detail",
		api.GetMenuDetailForWeek,
	},

	Route{
		"GetMenuByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/menu/{id}",
		api.GetMenuByID,
	},

	Route{
		"AddNotice",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/notices",
		api.AddNotice,
	},

	Route{
		"FindNoticeListByIndex",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/notices",
		api.FindNoticeListByIndex,
	},

	Route{
		"UpdateNotice",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/notices",
		api.UpdateNotice,
	},

	Route{
		"CreateNotification",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/notification",
		api.CreateNotification,
	},

	Route{
		"GetNotificationStatisticByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/notification/{id}/statistic",
		api.GetNotificationStatisticByID,
	},

	Route{
		"DeleteNotificationByID",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/notification/{id}",
		api.DeleteNotificationByID,
	},

	Route{
		"GetNotificationByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/notification/{id}",
		api.GetNotificationByID,
	},

	Route{
		"UpdateNotification",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/notification",
		api.UpdateNotification,
	},

	Route{
		"AddDeviceToken",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/device/subcribe",
		api.AddDeviceToken,
	},

	Route{
		"AddParent",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/parents",
		api.AddParent,
	},

	Route{
		"FindStudentByParentID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/parents/{id}/students",
		api.FindStudentByParentID,
	},

	Route{
		"FindStudentStatusByParentID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/parents/{id}/studentstatus",
		api.FindStudentStatusByParentID,
	},

	Route{
		"GetParentNotification",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/parent/{id}/notifications",
		api.GetParentNotification,
	},

	Route{
		"GetFormsByParentID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/parent/{id}/form",
		api.GetFormsByParentID,
	},

	Route{
		"UpdateParenNotificationStatus",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/parent/{id}/notifications",
		api.UpdateParenNotificationStatus,
	},

	Route{
		"CreateParentWithForm",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/parents/{id}",
		api.UpdateParent,
	},

	Route{
		"DeleteParent",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/parents/{id}",
		api.DeleteParent,
	},

	Route{
		"GetParentById",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/parents/{id}",
		api.GetParentById,
	},

	Route{
		"AddStudent",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/students",
		api.AddStudent,
	},

	Route{
		"FindStudentActivityByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/students/{id}/activity",
		api.FindStudentActivityByID,
	},

	Route{
		"FindStudentNoticeByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/students/{id}/notice",
		api.FindStudentNoticeByID,
	},

	Route{
		"FindStudentTuitionFeeyByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/students/tuitionfee",
		api.FindStudentTuitionFeeyByID,
	},

	Route{
		"DeleteStudentByID",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/students/{id}",
		api.DeleteStudentByID,
	},

	Route{
		"FindStudentByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/students/{id}",
		api.FindStudentByID,
	},

	Route{
		"FindStudentListByIndex",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/students",
		api.FindStudentListByIndex,
	},

	Route{
		"UpdateStudent",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/students/{id}",
		api.UpdateStudent,
	},

	Route{
		"UpdateStudentStatus",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/student/status",
		api.UpdateStudentStatus,
	},

	Route{
		"AddTeacher",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers",
		api.AddTeacher,
	},

	Route{
		"FindClassByTeacherID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}/class",
		api.FindClassByTeacherID,
	},

	Route{
		"FindNotificationByTeacherID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}/notification",
		api.FindNotificationByTeacherID,
	},

	Route{
		"FindActivityByTeacherID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}/activity",
		api.FindActivityByTeacherID,
	},

	Route{
		"DeleteTeacherByID",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}",
		api.DeleteTeacherByID,
	},

	Route{
		"FindTeacherByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}",
		api.FindTeacherByID,
	},

	Route{
		"FindTeacherListByIndex",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers",
		api.FindTeacherListByIndex,
	},

	Route{
		"UpdateTeacher",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/{id}",
		api.UpdateTeacher,
	},
	Route{
		"AddActivity",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/teachers/activity",
		api.AddActivity,
	},
	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/user",
		api.CreateUser,
	},

	Route{
		"CreateUsersWithListInput",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/users",
		api.CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/xuyendangvan/DemoAPI/1.0.0/user/{id}",
		api.DeleteUser,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/user/login",
		api.LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/user/logout",
		api.LogoutUser,
	},

	Route{
		"UpdatePassword",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/user/{id}/password",
		api.UpdatePassword,
	},

	Route{
		"GetUserByID",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/user/{id}",
		api.GetUserByID,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/user/{id}",
		api.UpdateUser,
	},

	Route{
		"CreateForm",
		strings.ToUpper("Post"),
		"/xuyendangvan/DemoAPI/1.0.0/form",
		api.CreateForm,
	},

	Route{
		"UpdateForm",
		strings.ToUpper("Put"),
		"/xuyendangvan/DemoAPI/1.0.0/form",
		api.UpdateForm,
	},

	Route{
		"GetStatisticForClass",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/statistic/classes/{id}",
		api.GetStatisticForClass,
	},

	Route{
		"GetStatisticForStudent",
		strings.ToUpper("Get"),
		"/xuyendangvan/DemoAPI/1.0.0/statistic/students",
		api.GetStatisticForStudent,
	},
}
