package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

var gCube *Cube = NewCube()

type http2room struct {
    p Proper
    c *gin.Context
    retCh chan struct{}
}

var httpDataCh chan *http2room = make(chan *http2room, 1)
var closeCh chan struct{} = make(chan struct{}, 1)
var printCh chan struct{} = make(chan struct{}, 1)

const RESP = "respone"

func main() {
    gCube.Print()
    go StartHttp()

    for {
        select {
        case h2r :=<- httpDataCh:
            retStr, err := h2r.p.Use(gCube)
            if err != nil {
                h2r.c.Set(RESP, err.Error())
                h2r.retCh <- struct{}{}
                continue
            }
            str := gCube.HttpPrint()
            h2r.c.Set(RESP, retStr+str)
            h2r.retCh <- struct{}{}
        case <- printCh:
            gCube.Print()
        case <- closeCh:
            return
        }
    }
}

func StartHttp() {
    router := gin.Default()
    gin.SetMode(gin.ReleaseMode)

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"ping": "pong"})
    })

    router.GET("/close", func(c *gin.Context) {
        closeCh <- struct{}{}
        c.JSON(http.StatusOK, gin.H{"msg": "closed"})
    })

    router.GET("/print", func(c *gin.Context) {
        printCh <- struct{}{}
        c.JSON(http.StatusOK, gin.H{"msg": "print"})
    })

    router.GET("/show", func(c *gin.Context) {
        c.String(http.StatusOK, gCube.HttpPrint())
    })

    faceGroup := router.Group("/act")
    faceGroup.Use(getFace)
    
    faceGroup.GET("/rotation", actRotation) // /act/rotation?face=0&clockWise=true
    faceGroup.GET("/missile", actMissile) // /act/missile?face=0
    faceGroup.GET("/dice", actDice)  // /act/dice?face=0

    srv := &http.Server{
        Addr:    "127.0.0.1:8090",
        Handler: router,
    }
    srv.SetKeepAlivesEnabled(false)
    srv.ListenAndServe()
}

const PARAM_FACE = "face"
const PARAM_CLOCKWISE = "clockWise"

func getFace(c *gin.Context) {
    faceStr, ok := c.GetQuery(PARAM_FACE)
    if ok {
        face, err := strconv.Atoi(faceStr)
        if err == nil {
            c.Set(PARAM_FACE, face)
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{"err": "face参数不存在"})
    c.Abort()
}

func actRotation(c *gin.Context) {
    clockWiseStr, ok := c.GetQuery(PARAM_CLOCKWISE)
    if ok {
        var clockWise bool
        if clockWiseStr == "true" {
            clockWise = true
        }
        proper := gCube.roles[c.GetInt(PARAM_FACE)].propRotation
        proper.clockWise = clockWise
        h2r := &http2room{
            p : proper,
            c : c,
            retCh : make(chan struct{}, 1),
        }
        httpDataCh <- h2r
        <-h2r.retCh
        retStr, _ := c.Get(RESP)
        c.String(http.StatusOK, "%s", retStr)
        return
    }
    c.String(http.StatusOK, "clockWise参数不存在")
}

func actMissile(c *gin.Context) {
    h2r := &http2room{
        p : gCube.roles[c.GetInt(PARAM_FACE)].propMissile,
        c : c,
        retCh : make(chan struct{}, 1),
    }
    httpDataCh <- h2r
    <-h2r.retCh
    retStr, _ := c.Get(RESP)
    c.String(http.StatusOK, "%s", retStr)
}

func actDice(c *gin.Context) {
    h2r := &http2room{
        p : gCube.roles[c.GetInt(PARAM_FACE)].propDice,
        c : c,
        retCh : make(chan struct{}, 1),
    }
    httpDataCh <- h2r
    <-h2r.retCh
    retStr, _ := c.Get(RESP)
    c.String(http.StatusOK, "%s", retStr)
}
