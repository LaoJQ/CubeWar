package main

import (
    
)

type Role struct {
    faceIdx int // 所在面
    squareIdx int // 所在面的格子

    batteryOrien int // 炮台朝向, 当前炮台朝向[]RotateRule索引
    batteryHP int // 炮台生命值

    propRotation *Rotation // 道具:旋转
    propMissile *Missile // 道具:导弹
    propDice *Dice // 道具:骰子
}

func NewRole(face int) *Role {
    return &Role{
        faceIdx : face,
        squareIdx : 0,

        batteryOrien : Rand.Number(4),
        batteryHP : 5,
        
        propRotation : &Rotation{selfFace : face},
        propMissile : &Missile{selfFace : face},
        propDice : &Dice{selfFace : face},
    }
}

func (role *Role) Move(cube *Cube, point int) {
    newSquareIdx := (role.squareIdx+point)%8

    newSquarePorpId := cube.face[role.faceIdx][newSquareIdx].propId
    switch newSquarePorpId {
    case PROP_ROTATE:
        role.propRotation.num++
    case PROP_MISSILE:
        role.propMissile.num++
    }
    
    role.squareIdx = newSquareIdx
}
