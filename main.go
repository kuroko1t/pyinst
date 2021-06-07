package main

import (
    "os/exec"
    "fmt"
    "flag"
    "regexp"
    "log"
)

func cmdRun(cmd []string) bool {
    out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
    if err != nil {
        re := regexp.MustCompile(`ModuleNotFoundError: No module named (.*)`)
        findresult := re.FindAllStringSubmatch(string(out), -1)
        if len(findresult) == 0 {
            return false
        }
        result := findresult[0][1]
        module_name := result[1:len(result)-1]
        module_name = checkPackage(module_name)
        fmt.Println(module_name)
        out, err := exec.Command("pip", "install", module_name).CombinedOutput()
        if err != nil {
            log.Fatal(string(out))
        }
        return true
    }
    return false
}

func checkPackage(module_name string) string {
    switch module_name {
    case "PIL":
        return "Pillow"
    case "cv2":
        return "opencv-python"
    case "yaml":
        return "pyyaml"
    default:
        return module_name
    }
}


func main() {
    flag.Parse()
    cmd := flag.Args()
    fmt.Println(cmd)
    fmt.Println("Install Package")
    fmt.Println("---------------")
    for {
        installFlag := cmdRun(cmd)
        if !installFlag {
            break
        }
    }
} 