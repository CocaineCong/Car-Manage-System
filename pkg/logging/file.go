package logging

//var (
//	LogSavePath = "runtime/logs/go.log"
//	LogSaveName = "log"
//	LogFileExt  = "log"
//	TimeFormat  = "20060102"
//)

//func getLogFilePath() string {
//	return fmt.Sprintf("%s", LogSavePath)
//}

//func openLogFile(filePath string) *os.File {
//	_, err := os.Stat(filePath)
//	switch {
//	case os.IsNotExist(err):
//		mkDir()
//	case os.IsPermission(err):
//		log.Fatalf("Permission:%v", err)
//	}
//	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0644)
//	if err != nil {
//		log.Fatalf("Fail to OpenFile:%v", err)
//	}
//	return handle
//}

//func mkDir() {
//	dir, _ := os.Getwd()
//	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
