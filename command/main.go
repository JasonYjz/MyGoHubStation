package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//cmd := exec.CommandContext(ctx, a.script, "start")
	//if err := cmd.Run(); err != nil {
	//	logger.Errorf("failed to start. err:%s", err.Error())
	//	a.Status = FAILURE
	//	return err
	//}
	fmt.Println("execute command: /var/opt/aidlux/data/cpf/aid-scs/file/pkg/ceshi220/2.10.220/manager.sh start")
	out, err := exec.CommandContext(ctx, "/var/opt/aidlux/data/cpf/aid-scs/file/pkg/ceshi220/2.10.220/manager.sh", "start").Output()
	if err != nil {
		fmt.Printf("failed to start. err:%s\n", err.Error())
		return
	}
	fmt.Printf("the ai program is started, out:%s\n", string(out))
}
