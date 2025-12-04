package utils

import "core:os"
import "core:strings"

read_input :: proc(filename: string) -> (string, bool) {
    data, ok := os.read_entire_file(filename)
    if !ok do return "", false
    return string(data), true
}

read_lines :: proc(filename: string) -> ([]string, bool) {
    data, ok := os.read_entire_file(filename)
    if !ok do return nil, false
    
    content := string(data)
    return strings.split(content, "\n"), true
}
