package main

import (
	"flag"
	"fmt"
	_ "github.com/codyguo/godaemon"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var _SWG sync.WaitGroup

// key prefix array
var _keyArray = []string{"device.profile", "user.device", "device.status"}

func initializeLogger(logMode int) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   "./logs/ttl.log", // 日志文件路径
		MaxSize:    200,              // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 100,              // 日志文件最多保存多少个备份
		MaxAge:     7,                // 文件最多保存多少天
		Compress:   true,             // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "log",
		//CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                          // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.123"), // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,                         //
		//EncodeCaller:   zapcore.FullCallerEncoder,                                 // 全路径编码器
		EncodeName: zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	multiWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)) // 打印到控制台和文件
	if logMode == 0 {
		multiWriter = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)) // 打印到文件
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		multiWriter,
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields()
	// 构造日志
	log := zap.New(core, caller, development, filed)
	log.Info("Log initialization successful!")
	return log
}

func _handle(keys []string, client *redis.Client, log *zap.Logger, statisticalRecords *sync.Map) {
	defer _SWG.Done()
	for _, key := range keys {
		//value, err = client.Get(ctx, key).Result()
		t, _ := client.TTL(key).Result()
		if t == time.Duration(-1) {
			_calculateRecords(statisticalRecords, "total_ttl_records")
			matched := false
			for _, k := range _keyArray {
				if strings.HasPrefix(key, k) {
					_calculateRecords(statisticalRecords, k)
					matched = true
					break
				}
			}
			if !matched {
				log.Error("special key:" + key)
				_calculateRecords(statisticalRecords, "special_records")
			}
		}
	}
}

func _calculateRecords(statisticalRecords *sync.Map, key string) {
	tmp, ok := statisticalRecords.Load(key)
	v := 1
	if ok {
		v = tmp.(int) + 1
	}
	statisticalRecords.Store(key, v)
}
func _printSummary(statisticalRecords *sync.Map, log *zap.Logger) {
	statisticalRecords.Range(func(key, value interface{}) bool {
		msg := "key:" + key.(string) + ",count:" + strconv.Itoa(value.(int))
		log.Info(msg)
		return true
	})
}

func main() {
	fmt.Println("current cpu num:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	taskBegin := time.Now()
	// 定义变量接收控制台参数
	// 日志输出模式
	var logMode int
	// 每次从redis中取数量
	var fetchNum int64
	//var categoryMapMux sync.RWMutex

	// StringVar用指定的名称、控制台参数项目、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量
	flag.IntVar(&logMode, "m", 0, "日志输出模式,0为控制台不输出日志,1控制台输出日志")
	flag.Int64Var(&fetchNum, "n", 2000, "每次读取批量数,默认为5000")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	log := initializeLogger(logMode)

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var cursor uint64
	var idx int
	statisticalRecords := &sync.Map{}
	fmt.Println("start all tasks.....", time.Now())
	log.Info("fetch keys per time:" + strconv.Itoa(int(fetchNum)))
	for {
		var keys []string
		var err error
		//*扫描所有key，每次N条
		keys, cursor, err = client.Scan(cursor, "*", fetchNum).Result()
		if err != nil {
			panic(err)
		}
		idx += len(keys)
		_SWG.Add(1)
		go _handle(keys, client, log, statisticalRecords)
		// exit loop
		if cursor == 0 {
			break
		}
	}
	_SWG.Wait()

	taskEnd := time.Now()
	duration := taskEnd.Sub(taskBegin)
	fmt.Printf("duration: %v\n", duration)
	//cost := time.Since(taskBegin).Seconds()
	fmt.Println("all task finished...", taskEnd)
	fmt.Println("all task finished, cost:", duration)
	log.Info("===>total records:" + strconv.Itoa(idx))
	_printSummary(statisticalRecords, log)
	log.Info("===>all tasks completed, ",
		zap.String("task begin time:", taskBegin.String()),
		zap.String("end time", taskEnd.String()),
		zap.Duration("cost", duration))
}
