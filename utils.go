// utils.go kee > 2021/09/21

package log

func configGet(config map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value := config[key]
	if value != nil {
		return value
	}
	return defaultValue
}
