package main

import (
    
)

type Role struct {
    faceIdx int // 所在面
    squareIdx int // 所在面的格子

    propRotation *Rotation // 道具:旋转
    propMissile *Missile // 道具:导弹
    propDice *Dice // 道具:骰子
}

func NewRole(face int) *Role {
    return &Role{
        faceIdx : face,
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
