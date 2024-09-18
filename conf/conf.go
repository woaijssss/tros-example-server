package conf

/*
	除了可以使用系统配置外，还可以在此自定义配置文件内容，并通过初始化函数读取
*/

//func InitConfig() {
//	defer func() {
//		if r := recover(); r != nil {
//			log.Fatalf("init config down [stack=%+v] [recover=%+v]", string(debug.Stack()), r)
//		}
//	}()
//	initConfig()
//}
//
//func initConfig() {
//	config = viper.New()
//	config.AddConfigPath("./conf/")
//	config.SetConfigName("app")
//	config.SetConfigType("yaml")
//	err := config.ReadInConfig()
//	if err != nil {
//		log.Printf("%v", err)
//		os.Exit(1)
//	}
//
//	err = config.Unmarshal(&appConfig)
//	if err != nil {
//		fmt.Println(" Unmarshal error +v%", err)
//		os.Exit(1)
//	}
//}
