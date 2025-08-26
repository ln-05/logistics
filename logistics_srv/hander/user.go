package hander

import (
	"context"
	"fmt"
	"logistics_srv/basic/global"
	"logistics_srv/model"
	__ "logistics_srv/proto"
	"math/rand"
	"regexp"
	"time"
)

type UserServer struct {
	__.UnimplementedUserServer
}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

// 发送验证码
func (c *UserServer) SendSms(_ context.Context, in *__.SendSmsRequest) (*__.SendSmsResponse, error) {
	//判断手机号是否正确
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}
	// 用redis分布式锁实现60秒内只能获取一次
	set, _ := global.Rdb.SetNX(context.Background(), "SendSmsLimit:"+in.Mobile, 1, time.Second*60).Result()
	if !set {
		return &__.SendSmsResponse{
			Code:    200,
			Message: "短信发送过于频繁，请60秒后再试",
		}, nil
	}

	//生成验证码
	code := rand.Intn(9000) + 1000

	//将生成的验证码存储到redis缓存中，并设置过期时间5分钟
	global.Rdb.Set(context.Background(), "SendSms"+in.Mobile, code, time.Minute*5)

	return &__.SendSmsResponse{
		Code:    200,
		Message: "发送验证码成功",
	}, nil
}

// 登录注册
func (c *UserServer) Login(_ context.Context, in *__.LoginRequest) (*__.LoginResponse, error) {
	//判断手机号
	if !isValidMobile(in.Mobile) {
		return nil, fmt.Errorf("手机号格式不正确")
	}

	var user model.User
	global.DB.Where("mobile=?", in.Mobile).Find(&user)
	// 判断用户是否存在，如果不存在直接注册
	if user.Id == 0 {
		newUser := model.User{
			Mobile: in.Mobile,
		}

		if err := global.DB.Create(&newUser).Error; err != nil {
			return nil, fmt.Errorf("注册失败")
		}
	}
	// 接受验证码，判断验证码是否过期和验证码是否错误
	get := global.Rdb.Get(context.Background(), "SendSms"+in.Mobile)
	if get.Err() != nil {
		return nil, fmt.Errorf("验证码已过期")
	}

	if get.Val() != in.SendSmsCode {
		return nil, fmt.Errorf("验证码错误")
	}
	// 5.验证成功后，删除验证码
	global.Rdb.Del(context.Background(), "SendSms"+in.Mobile)

	return &__.LoginResponse{
		Id: int64(user.Id),
	}, nil
}
