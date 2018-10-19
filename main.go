package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

var gCube *Cube = NewCube()

var httpDataCh chan Proper = make(chan Proper, 1)
var closeCh chan struct{} = make(chan struct{}, 1)
var printCh chan struct{} = make(chan struct{}, 1)

func main() {
    gCube.Print()
    go StartHttp()

    for {
        select {
        case proper :=<- httpDataCh:
            proper.Use(gCube)
            gCube.Print()
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
        face := c.GetInt(PARAM_FACE)
        propRotation := gCube.roles[face].propRotation
        propRotation.clockWise = clockWise
        httpDataCh <- propRotation
        c.JSON(http.StatusOK, gin.H{"msg": "maybe ok"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"msg": "clockWise参数不存在"})
}

func actMissile(c *gin.Context) {
    face := c.GetInt(PARAM_FACE)
    propRotation := gCube.roles[face].propMissile
    httpDataCh <- propRotation
    c.JSON(http.StatusOK, gin.H{"msg": "maybe ok"})
}

func actDice(c *gin.Context) {
    face := c.GetInt(PARAM_FACE)
    propRotation := gCube.roles[face].propDice
    httpDataCh <- propRotation
    c.JSON(http.StatusOK, gin.H{"msg": "maybe ok"})
}
