print("\n--- Quaternion 1 ---")
local q1 = Quat(0, 0, 0, 1)
print("String representation:", q1)
print("Table:")
for key, value in pairs(q1.table) do
  print(key, value)
end
local q2 = Quat(1, 0, 0, 0)
print("\n--- Quaternion 2 ---")
print("String representation:", q2)
print("Fields:")
print("X:", q2.x)
print("Y:", q2.y)
print("Z:", q2.z)
print("W:", q2.w)
print("\n--- Operations ---")
print("Multiplication:", q1 * q2)
print("Addition:", q1 + q2)
print()
print(string.format("Angle between quaternions: %.3f radians", q1:rad(q2)))
print(string.format("Angle between quaternions: %.3f degrees", q1:deg(q2)))
