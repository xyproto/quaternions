// Example code for getting closer to operator overloading in Go through employing Lua
package main

import (
	"fmt"

	"github.com/go-utils/unum"
	"github.com/yuin/gopher-lua"
)

const (
	// Lua file that can use the Quat class that is defined below
	luaFilename = "main.lua"

	// Internal ID for the Quat class in Lua
	lQuatClass = "Quat"
)

// List of methods to register for the Quat class in Lua
// For an overview of available metamethods, see:
// http://www.lua.org/manual/5.1/manual.html#2.8
var quatMethods = map[string]lua.LGFunction{
	"__tostring": quatString,
	"__mul":      quatMul,
	"__add":      quatAdd,
	"rad":        quatAngleRad,
	"deg":        quatAngleDeg,
}

// Method for multiplying quaternions in Lua
func quatMul(L *lua.LState) int {
	self := checkQuat(L, 1)  // arg 1 (the object)
	other := checkQuat(L, 2) // arg 2 (first method argument)
	// Multiply self with other and return the result
	result := self.Mul(other)
	userdata, err := constructQuat(L, result)
	if err != nil {
		fmt.Println("ERROR", err)
		L.Push(lua.LString(err.Error()))
		return 1 // Number of returned values
	}
	L.Push(userdata)
	return 1 // Number of returned values
}

// Method for adding quaternions in Lua
func quatAdd(L *lua.LState) int {
	self := checkQuat(L, 1)  // arg 1 (the object)
	other := checkQuat(L, 2) // arg 2 (first method argument)
	// Add self with other and return the result
	result := unum.NewQuat(self.X+other.X, self.Y+other.Y, self.Z+other.Z, self.W+other.W)
	userdata, err := constructQuat(L, result)
	if err != nil {
		fmt.Println(err)
		L.Push(lua.LString(err.Error()))
		return 1 // Number of returned values
	}
	L.Push(userdata)
	return 1 // Number of returned values
}

// Method for getting the angle between two quaternions in Lua, in radians
func quatAngleRad(L *lua.LState) int {
	self := checkQuat(L, 1)  // arg 1 (the object)
	other := checkQuat(L, 2) // arg 2 (first method argument)
	// Get the angle in float64
	result := self.AngleRad(other)
	L.Push(lua.LNumber(result))
	return 1 // Number of returned values
}

// Method for getting the angle between two quaternions in Lua, in degrees
func quatAngleDeg(L *lua.LState) int {
	self := checkQuat(L, 1)  // arg 1 (the object)
	other := checkQuat(L, 2) // arg 2 (first method argument)
	// Get the angle in float64
	result := self.AngleDeg(other)
	L.Push(lua.LNumber(result))
	return 1 // Number of returned values
}

// For converting from a *unum.Quat to a userdata Quat in Lua
func constructQuat(L *lua.LState, q *unum.Quat) (*lua.LUserData, error) {
	ud := L.NewUserData()
	ud.Value = q
	L.SetMetatable(ud, L.GetTypeMetatable(lQuatClass))
	return ud, nil
}

// Check that the given argument number is a userdata Quat, and return it
func checkQuat(L *lua.LState, argnr int) *unum.Quat {
	ud := L.CheckUserData(argnr)
	if quat, ok := ud.Value.(*unum.Quat); ok {
		return quat
	}
	L.ArgError(argnr, "Quat expected")
	return nil
}

// Represent a userdata Quat as a Lua string
func quatString(L *lua.LState) int {
	self := checkQuat(L, 1) // arg 1 (before the ":")
	// Create a string representation of the Quaternion
	repr := fmt.Sprintf("[%.3f %.3f %.3f %.3f]", self.X, self.Y, self.Z, self.W)
	L.Push(lua.LString(repr))
	return 1 // number of results
}

func main() {
	// Create a new Lua VM
	L := lua.NewState()
	defer L.Close()

	// Register the Quat class and the methods that belongs with it.
	mt := L.NewTypeMetatable(lQuatClass)
	mt.RawSetH(lua.LString("__index"), mt)
	L.SetFuncs(mt, quatMethods)

	// The Lua constructor for new Quat objects. Takes four numbers.
	L.SetGlobal("Quat", L.NewFunction(func(L *lua.LState) int {

		// Get the four numeric values
		x := float64(L.ToNumber(1)) // argument 1
		y := float64(L.ToNumber(2)) // argument 2
		z := float64(L.ToNumber(3)) // argument 3
		w := float64(L.ToNumber(4)) // argument 4

		// Construct a new Quat
		userdata, err := constructQuat(L, unum.NewQuat(x, y, z, w))
		if err != nil {
			fmt.Println("ERROR", err)
			L.Push(lua.LString(err.Error()))
			return 1 // Number of returned values
		}

		// Construct a Lua table with the four values
		table := L.NewTable()
		table.Append(lua.LNumber(x))
		table.Append(lua.LNumber(y))
		table.Append(lua.LNumber(z))
		table.Append(lua.LNumber(w))

		// Set object fields that are not methods, but values
		mt := L.GetMetatable(userdata)
		L.SetField(mt, "x", lua.LNumber(x))
		L.SetField(mt, "y", lua.LNumber(y))
		L.SetField(mt, "z", lua.LNumber(z))
		L.SetField(mt, "w", lua.LNumber(w))
		L.SetField(mt, "table", table)
		L.SetMetatable(userdata, mt)

		// Return the Lua Quat object
		L.Push(userdata)
		return 1 // number of results
	}))

	// Run the Lua script
	if err := L.DoFile(luaFilename); err != nil {
		panic(err)
	}
}
