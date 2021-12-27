package data

import (
	"docker-manager/data/base"
	"docker-manager/data/table"
	"log"
)

func AddWhiteIp(ip string) (err error) {
	record := table.WhiteIp{
		IP: ip,
	}
	_, err = base.DBEngine.Table("white_ip").Insert(&record)
	return err
}

func DelWhiteIp(ip string) (err error) {
	var record table.WhiteIp
	_, err = base.DBEngine.Table("white_ip").Where("ip=?", ip).Get(&record)
	if err != nil {
		log.Println("DelWhiteIp.find.err:", err)
		return err
	}
	affected, err := base.DBEngine.Table("white_ip").Delete(&record)
	if err != nil {
		log.Println("DelWhiteIp.err:", err)
		return err
	}
	if affected == 0 {
		log.Println("DelWhiteIp.affected is 0")
	}
	return
}

func GetWhiteIpList() (record []table.WhiteIp, err error) {
	err = base.DBEngine.Table("white_ip").Find(&record)
	if err != nil {
		log.Println("GetWhiteIpList:", err)
	}
	return
}
