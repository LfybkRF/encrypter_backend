package encodings

import (
	"math"
	"strconv"
	"strings"
)

// Функция для шифрования сообщения
func EncryptHill(message, key string) string {
	message = strings.ReplaceAll(message, " ", "")
	if !IsLetter(message) {
		return "Wrong input message"
	}
	if IsLetter(key) {
		return "Wrong input key"
	}
	keyslice := invert(key)
	if keyslice[0] * keyslice[2] - keyslice[1] * keyslice[3] == 0 {
		return "Wrong input key"
	}
	keymatrix := getMatrixKey(keyslice)

	//Добавляем недостающие символы
	message = addPadding(message)
	message = strings.ToUpper(message)

	base := byte('A')

	//Разбиваем сообщение на пары символов
	pairs := make([][]int, len(message)/2)
	for i := 0; i < len(message); i += 2 {
		pairs[i/2] = []int{int(message[i] - base), int(message[i+1] - base)}
	}

	//Умножаем каждую пару на ключ
	result := ""
	for _, pair := range pairs {
		encryptedPair := multiplyMatrix(keymatrix, [][]int{{pair[0]}, {pair[1]}})
		result += string(encryptedPair[0][0] + int(base))
		result += string(encryptedPair[1][0] + int(base))
	}

	return result
}

// Функция для дешифрования сообщения
func DecryptHill(message, key string) string {
	if strings.Contains(message, " ") {
		return "Wrong input message"
	}
	if !IsLetter(message) {
		return "Wrong input message"
	}
	if IsLetter(key) {
		return "Wrong input key"
	}

	keyslice := invert(key)
	keymatrix := getMatrixKey(keyslice)
	//Находим обратную матрицу
	inverseKey := inverseMatrix(keymatrix)
	message = strings.ReplaceAll(message, " ", "")
	base := byte('A')
	//Разбиваем сообщение на пары символов
	pairs := make([][]int, len(message)/2)
	for i := 0; i < len(message); i += 2 {
		pairs[i/2] = []int{int(message[i] - base), int(message[i+1] - base)}
	}

	//Умножаем каждую пару на обратный ключ
	result := ""
	for _, pair := range pairs {
		decryptedPair := multiplyMatrix(inverseKey, [][]int{{pair[0]}, {pair[1]}})
		result += string(decryptedPair[0][0] + int(base))
		result += string(decryptedPair[1][0] + int(base))
	}

	result = removePadding(result)
	return result
}

func invert(keystring string) []int {
	var err error
	keySliceStr := strings.Split(keystring, " ")
	keySliceInt := make([]int, len(keySliceStr))

	for i, x := range keySliceStr {
		keySliceInt[i], err = strconv.Atoi(x)
	}
	if err != nil {
		panic(err)
	}
	return keySliceInt
}

// Получения матричного ключа
func getMatrixKey(keyslice []int) [][]int {

	keymatrix := make([][]int, 2)
	for k := range keymatrix {
		keymatrix[k] = make([]int, 2)
	}

	k := 0
	for _, i := range keymatrix {
		for j := range i {
			i[j] = keyslice[k]
			k++
		}
	}

	return keymatrix
}

// Функция для нахождения обратной матрицы
func inverseMatrix(key [][]int) [][]int {
	determinant := key[0][0]*key[1][1] - key[0][1]*key[1][0]
	inverseDeterminant := int(math.Mod(float64(determinant), 26))

	//Обратный определитель
	for i := 1; i < 26; i++ {
		if (inverseDeterminant*i)%26 == 1 {
			inverseDeterminant = i
			break
		}
	}
	//Матрица алгебраических дополнений
	adjugateMatrix := [][]int{{key[1][1], -key[0][1]}, {-key[1][0], key[0][0]}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			adjugateMatrix[i][j] = int(math.Mod(float64(adjugateMatrix[i][j]*inverseDeterminant), 26))
		}
	}

	return adjugateMatrix
}

// Нахождение мода для отрицательных и положительных чисел
func mod(a, b int) int {
	return (a%b + b) % b
}

// Функция для умножения матриц
func multiplyMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for i := range result {
		result[i] = make([]int, len(matrix2[0]))
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[k][i] * matrix2[k][j]
			}
			result[i][j] = int(mod(int(result[i][j]), 26))
		}
	}

	return result
}
