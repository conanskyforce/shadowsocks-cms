package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/cheneylew/goutil/utils"
	"fmt"
	"time"
	"github.com/cheneylew/shadowsocks-cms/models"
	"github.com/cheneylew/goutil/utils/beego"
)

var O orm.Ormer
func init() {

	url := beego.DBUrl("cheneylew","12344321","47.91.151.207","3308","shadowsocks-servers")
	utils.JJKPrintln(url)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", url)
	O = orm.NewOrm()

	if err != nil {
		utils.JJKPrintln("========database can't connect! error:" + err.Error()+"========")
	} else {
		utils.JJKPrintln("========database connected success！========")
	}

}

func DBQueryServers(ip string) []*models.Server {
	var objects []*models.Server

	qs := O.QueryTable("server")
	_, err := qs.Filter("ip", ip).All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryServersWithSid(sid int) *models.Server {
	var objects []*models.Server

	qs := O.QueryTable("server")
	_, err := qs.Filter("server_id", sid).All(&objects)

	if err == nil && len(objects) > 0 {
		return objects[0]
	}

	return nil
}

func DBQueryServersAll() []*models.Server {
	var objects []*models.Server

	qs := O.QueryTable("server")
	_, err := qs.OrderBy("-server_id").RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithSid(sid int64) []*models.Port {
	var objects []*models.Port

	qs := O.QueryTable("port")
	_, err := qs.Filter("server__server_id", sid).OrderBy("-port_id").RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithUserId(userId int64) []*models.Port {
	var objects []*models.Port

	qs := O.QueryTable("port")
	_, err := qs.Filter("user__user_id", userId).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryPortsWithIP(ip string) []*models.Port {
	var objects []*models.Port

	qs := O.QueryTable("port")
	_, err := qs.Filter("server__ip", ip).RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryMaxPortWithIP(ip string) *models.Port {
	var objects []*models.Port

	qs := O.QueryTable("port")
	_, err := qs.Filter("server__ip", ip).OrderBy("-port").All(&objects)

	if err == nil && len(objects) > 0 {
		return objects[0]
	}

	return nil
}

func DBQueryPortWithPid(pid int) *models.Port {
	var objects []*models.Port

	qs := O.QueryTable("port")
	_, err := qs.Filter("port_id", pid).RelatedSel().All(&objects)

	if err == nil && len(objects) > 0 {
		return objects[0]
	}

	return nil
}

func DBQueryUserWithUid(uid int64) *models.User {
	var objects []*models.User
	qs := O.QueryTable("user")
	_, err := qs.Filter("user_id", uid).RelatedSel().All(&objects)

	if err != nil || len(objects) == 0 {
		return nil
	}

	return objects[0]
}

func DBQueryUsersAll() []*models.User {
	var objects []*models.User
	qs := O.QueryTable("user")
	_, err := qs.OrderBy("-user_id").RelatedSel().All(&objects)

	if err != nil {
		return objects
	}

	return objects
}

func DBQueryUserWithEmailOrMobile(emailOrMobile string) []models.User {
	var objects []models.User
	_, err := O.Raw(fmt.Sprintf("select * from user where email =? or mobile=?"),emailOrMobile, emailOrMobile).QueryRows(&objects)
	if err != nil {
		utils.JJKPrintln(err)
		return objects
	}

	return objects
}

func DBQueryMyListenPorts() []*models.Port {
	curIp, _ := utils.ExtranetIP()
	ports := DBQueryPortsWithIP(curIp)

	var filterPorts []*models.Port
	for _, port := range ports {
		if port.Ptype == 0 {
			//包年包月
			//判断截止时间
			if time.Now().Before(port.End_time) {
				filterPorts = append(filterPorts, port)
			}
		} else if port.Ptype == 1 {
			//包流量
			//流量是否超限
			if port.Flow_total < port.Flow_in_max {
				filterPorts = append(filterPorts, port)
			}
		}
	}


	return filterPorts
}

func dbStart()  {
	O = orm.NewOrm()
	O.Using("default") // 默认使用 default，你可以指定为其他数据库


}