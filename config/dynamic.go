package config

type DatabaseLocalType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

type DatabaseDevType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

type DatabaseProdType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

const MaxRowsTable = 30

var MaxRowList = []int{10, 20, 30, 50, 100}

// For Dynamic DB
type ConfigDynamicType struct {
	DatabaseLocal DatabaseLocalType
	DatabaseDev   DatabaseDevType
	DatabaseProd  DatabaseProdType
}

var ConfigDynamic ConfigDynamicType

func LoadConfigDynamic() {
	ConfigDynamic = ConfigDynamicType{
		DatabaseLocal: DatabaseLocalType{
			Host:         "localhost",
			Username:     "root",
			Password:     "",
			DatabaseName: "grelegant",
			Port:         "8080",
		},
		DatabaseDev: DatabaseDevType{
			Host:         "ep-empty-grass-a2xemsef.eu-central-1.pg.koyeb.app", //deploy dev host
			Username:     "koyeb-adm",
			Password:     "qxALFflgc3N7	",
			DatabaseName: "planting-db",
			Port:         "5432",
		},
		DatabaseProd: DatabaseProdType{
			Host:         "ep-empty-grass-a2xemsef.eu-central-1.pg.koyeb.app", //deploy prod host
			Username:     "koyeb-adm",
			Password:     "qxALFflgc3N7	",
			DatabaseName: "planting-db",
			Port:         "5432",
		},
	}
}
