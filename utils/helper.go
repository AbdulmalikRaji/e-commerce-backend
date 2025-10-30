package utils

import (
	"bytes"
	"encoding/json"
	"math"
	"regexp"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConvertStruct(sourceItem any, targetItem any) error {
	jsonItem, err := json.Marshal(sourceItem)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonItem, targetItem)
	if err != nil {
		return err
	}

	return nil
}

func GetLanguage(c *fiber.Ctx) string {
	lang := c.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}
	return lang
}

const (
	PgDuplicateErrorCode = "23505"
)

func EmailRegex(email string) bool {
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return regexpEmail.MatchString(email)
}

func PhoneNumberRegex(phoneNumber string) bool {
	regexpPhoneNumber := regexp.MustCompile(`^\+[0-9]{1,3}[\s.-]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}([-\s\.]?[0-9]){4,6}$`)
	return regexpPhoneNumber.MatchString(phoneNumber)
}

func CalculatePercentage(currentValue float64, previousValue float64) (float64, error) {
	percentage := (currentValue - previousValue) / previousValue * 100
	return percentage, nil
}

func ConvertStringToFloat(value string) (float64, error) {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return convertedValue, nil
}

type File struct {
	File     *bytes.Buffer
	FileName string
	FileSize int
	FileType string
}

func StructToArrayMap(data any) (map[string][]string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result map[string][]string
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ContainsInt(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsString(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func Round(number float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(number*output) / output
}

func IsFloat64(value float64) bool {
	return value != float64(int(value))
}

func IsInt(value float64) bool {
	return value == float64(int(value))
}

func MapToJSONString(m map[string]bool) (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func JsonStringToMap(s string) (map[string]bool, error) {
	var m map[string]bool
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func StringContains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
