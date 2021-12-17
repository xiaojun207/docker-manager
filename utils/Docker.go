package utils

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func TrimContainerName(s interface{}) string {
	names := ""
	typeName := reflect.TypeOf(s).String()
	if typeName == "[]interface {}" {
		names = s.([]interface{})[0].(string)
	} else if typeName == "string" {
		names = s.(string)
	} else {
		log.Println("typeName:", typeName)
	}

	return strings.TrimLeft(names, "/")
}

/**
To calculate the values shown by the `stats` command of the docker cli tool
the following formulas can be used:
* used_memory = `memory_stats.usage - memory_stats.stats.cache`
* available_memory = `memory_stats.limit`
* Memory usage % = `(used_memory / available_memory) * 100.0`
* cpu_delta = `cpu_stats.cpu_usage.total_usage - precpu_stats.cpu_usage.total_usage`
* system_cpu_delta = `cpu_stats.system_cpu_usage - precpu_stats.system_cpu_usage`
* number_cpus = `lenght(cpu_stats.cpu_usage.percpu_usage)` or `cpu_stats.online_cpus`
* CPU usage % = `(cpu_delta / system_cpu_delta) * number_cpus * 100.0`
* @param c
* @returns {string}
*/
func FormatCpu(cpu_stats, precpu_stats map[string]interface{}) float64 {
	// https://my.oschina.net/jxcdwangtao/blog/828648
	// cpu_delta = cpu_total_usage - pre_cpu_total_usage;
	// system_delta = system_usage - pre_system_usage;
	// CPU % = ((cpu_delta / system_delta) * length(per_cpu_usage_array) ) * 100.0

	if cpu_stats == nil || cpu_stats["cpu_usage"] == nil ||
		precpu_stats == nil || precpu_stats["cpu_usage"] == nil ||
		cpu_stats["system_cpu_usage"] == nil || precpu_stats["system_cpu_usage"] == nil {
		return 0
	}

	cpu_delta := cpu_stats["cpu_usage"].(map[string]interface{})["total_usage"].(float64) - precpu_stats["cpu_usage"].(map[string]interface{})["total_usage"].(float64)
	system_cpu_delta := cpu_stats["system_cpu_usage"].(float64) - precpu_stats["system_cpu_usage"].(float64)
	number_cpus := cpu_stats["online_cpus"].(float64)
	cpu_usage_percent := (cpu_delta / system_cpu_delta) * number_cpus * 100.0
	return cpu_usage_percent
}

func FormatSize(s float64) string {
	res := ""
	if s < 1024 {
		res = strconv.FormatFloat(s, 'f', 2, 64) + "B"
	} else if s < 1024*1024 {
		res = strconv.FormatFloat(s/1024, 'f', 2, 64) + "KB"
	} else if s < 1024*1024*1024 {
		res = strconv.FormatFloat(s/(1024*1024), 'f', 2, 64) + "MB"
	} else {
		res = strconv.FormatFloat(s/(1024*1024*1024), 'f', 2, 64) + "GB"
	}
	return res
}

func FormatMemoryPercent(m map[string]interface{}) float64 {
	if m == nil || m["usage"] == nil {
		return 0
	}
	return m["usage"].(float64) * 100.0 / m["limit"].(float64)
}

func FormatMemory(m map[string]interface{}) string {
	if m == nil || m["usage"] == nil {
		return ""
	}
	return FormatSize(m["usage"].(float64)) + " / " + FormatSize(m["limit"].(float64))
}

func FormatNet(n map[string]interface{}) string {
	if n == nil || n["eth0"] == nil {
		return ""
	}
	/**
	 * { "eth0": { "rx_bytes": 2336, "rx_dropped": 0, "rx_errors": 0, "rx_packets": 32, "tx_bytes": 0, "tx_dropped": 0, "tx_errors": 0, "tx_packets": 0 } }
	 */
	e := n["eth0"].(map[string]interface{})
	return FormatSize(e["rx_bytes"].(float64)) + " / " + FormatSize(e["tx_bytes"].(float64))
}
