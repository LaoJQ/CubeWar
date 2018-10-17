package main

import (
    "fmt"    
)

// 魔方整体数据结构
type Cube struct {
    face [][]Square
}

// 单个格子数据
type Square struct {
    color byte // just for test

    id int // 道具Id
    batteryOrien int // 炮台朝向, 当前炮台朝向[]RotateRule索引
    batteryHP int // 炮台生命值
}

// 道具接口
type Proper interface {
    Use(*Cube) error
}


var ColorQueue []byte = []byte{'Y','W','B','G','O','R'}

func NewCube() *Cube {
    cube := new(Cube)
    for i:=0; i<6; i++ {
        oneFace := make([]Square, 9)
        for j:=0; j<9; j++ {
            oneFace[j].color = ColorQueue[i]
        }
        oneFace[8].batteryOrien = Rand.Number(4)
        oneFace[8].batteryHP = 5
        cube.face = append(cube.face, oneFace)
    }
    return cube
}

func (cube *Cube) Print() {
    for i:=0; i<6; i++ {
        squares := make([]byte, 9)
        for _, square := range cube.face[i] {
            squares = append(squares, square.color)
        }
        fmt.Println(string(squares), RotateRules[i][cube.face[i][8].batteryOrien].faceIdx, cube.face[i][8])
    }
    fmt.Println("---------------")
}
