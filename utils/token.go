package utils

func GenerateToken(pepper string, cost int) (string, error) {
	uuid := GenerateRandomID().String()
	return GenerateHash(uuid, pepper, cost)
}
