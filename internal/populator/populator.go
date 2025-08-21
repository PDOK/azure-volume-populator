package populator

import (
	"os"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

func Populate(fileName, fileContents string) {
	if "" == fileName || "" == fileContents {
		klog.Fatalf("Missing required arg")
	}
	f, err := os.Create(fileName)
	if nil != err {
		klog.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	if !strings.HasSuffix(fileContents, "\n") {
		fileContents += "\n"
	}

	_, err = f.WriteString(fileContents)
	if nil != err {
		klog.Fatalf("Failed to write to file: %v", err)
	}

	time.Sleep(1 * time.Minute)
}
