package bonification

func CalculateBonification(coins int, exchange float32) int {
	return int(float32(coins) * exchange)
}
