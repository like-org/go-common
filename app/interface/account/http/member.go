package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonCreateMember struct {
	Mobile string `json:"mobile"`
	Code   int    `json:"code"`
}

func (srv *Server) CreateMember(ctx *gin.Context) {
	var data jsonCreateMember
	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Printf("mobile get error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad param"})
		return
	}

	member, err := srv.accSrv.RegisterByMobile(ctx, data.Mobile, data.Code)
	if err != nil {
		log.Printf("RegisterByMobile error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RegisterByMobile Error"})
		return
	}

	mem := gin.H{
		"id":         member.ID,
		"mobile":     member.Mobile,
		"created_at": member.CreatedAt,
	}
	ctx.JSON(http.StatusOK, mem)
}

func (srv *Server) GetVerfiyCode(ctx *gin.Context) {
	key := ctx.DefaultQuery("key", "")
	if len(key) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad param"})
		return
	}
	err := srv.accSrv.SendVerifyCode(ctx, key)
	if err != nil {
		log.Printf("Err SendVerifyCode: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "send verify code failure"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
