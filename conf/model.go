package conf

type ConfigDB struct {
	DriverName    string `yaml:"driverName"`
	DataSourceUrl string `yaml:"dataSourceUrl"`
	UseCache      string `yaml:"useCache"`
}

type ConfigServer struct {
	ApplicationName string `yaml:"applicationName"`
	Port            string `yaml:"port"`
	ContextPath     string `yaml:"contextPath"`
	ConsoleCode     string `yaml:"consoleCode"`
}

type Config struct {
	Server ConfigServer `yaml:"server"`
	DB     ConfigDB     `yaml:"db"`
}
