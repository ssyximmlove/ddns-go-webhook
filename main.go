package main

import (
	"github.com/bytedance/sonic"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type requestBody struct {
	Key      string `json:"key"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Group    string `json:"group"`
	Priority int    `json:"priority"`
}

func init() {
	//初始化日志和配置文件
	InitLogger()
	LoadConfig()
}

func main() {
	//创建HTTP服务器
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    viper.GetString("app.addr"),
		Handler: mux,
	}
	//注册路由
	mux.HandleFunc("POST /webhook", webhookHandler)
	//启动服务器
	logger.Info("Server started", zap.String("addr", viper.GetString("app.addr")))
	err := server.ListenAndServe()
	//处理错误
	if err != nil {
		logger.DPanic("Failed to start server", zap.Error(err))
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	//接受DDNS-Go的Webhook请求
	payload := &requestBody{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Failed to read request body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Info("Received webhook request", zap.String("body", string(body)))
	//解析请求体
	err = sonic.Unmarshal(body, payload)
	if err != nil {
		logger.Error("Failed to unmarshal request body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//构造请求体
	form := url.Values{}
	form.Set("key", payload.Key)
	form.Set("content", payload.Content)
	form.Set("type", payload.Type)
	form.Set("title", payload.Title)
	form.Set("group", payload.Group)
	form.Set("priority", strconv.Itoa(payload.Priority))
	//发送请求
	resp, err := http.Post(
		viper.GetString("app.endpoint"),
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		logger.Error("Failed to send request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	//读取响应体
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//返回响应
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		logger.Error("Failed to write response", zap.Error(err))
	}
}
