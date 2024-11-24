package ttkdescriber

import "testing"

func TestAdd(t *testing.T) {
	damage := NewDamage()
	CalCombinationTTK(damage)
}

func NewDamage() Damage {
	damage := Damage{}
	// damage.firerate = 667
	// damage.health = 300

	// damage.headRate = 15
	// damage.neckRate = 5
	// damage.upperTorsoRate = 30
	// damage.lowerTorsoRate = 20
	// damage.upperArmRate = 10
	// damage.lowerArmRate = 5
	// damage.upperLegRate = 10
	// damage.lowerLegRate = 5

	// damage.head = 43
	// damage.neck = 41
	// damage.upperTorso = 41
	// damage.lowerTorso = 36
	// damage.upperArm = 41
	// damage.lowerArm = 41
	// damage.upperLeg = 34
	// damage.lowerLeg = 34

	damage.firerate = 800
	damage.health = 300

	damage.headRate = 50
	damage.neckRate = 50
	// damage.upperTorsoRate = 25
	// damage.lowerTorsoRate = 25

	damage.head = 50
	damage.neck = 40
	// damage.upperTorso = 100
	// damage.lowerTorso = 50

	return damage
}
