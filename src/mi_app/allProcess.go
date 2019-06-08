package main
import
(
	"os"
	"log"	
	"fmt"
	"strconv"
    
)

func existsFile(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}

func Pids() ([]int, error) {
    f, err := os.Open(`/proc`)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    names, err := f.Readdirnames(-1)
    if err != nil {
        return nil, err
    }
    pids := make([]int, 0, len(names))
    for _, name := range names {
		
        if pid, err := strconv.ParseInt(name, 10, 0); err == nil {
			pids = append(pids, int(pid))
		}
		
    }
    return pids, nil
}


func readAllProcess(){
	var Procesos_R  int
	Procesos_R = 0
	var Procesos_S  int
	Procesos_S = 0
	var Procesos_Z  int
	Procesos_Z = 0
	var Procesos_T  int
	Procesos_T = 0

	var Procesos_D  int
	Procesos_D = 0
	pids, err := Pids()
    if err != nil {
        fmt.Println("pids:", err)
        return
	}

	var index int
	var ruta string

	for i := range pids{

		index = pids[i]
		ruta= "/proc/"+strconv.Itoa(index)+"/status"
		if  existsFile(ruta){
			stat, err := ReadProcessStatus(ruta)
			if err != nil {
				log.Fatal("stat read fail:",err)
				continue;
			}
			if stat.State[0]=='R'{
				Procesos_R++
			}
			if stat.State[0]=='Z'{
				Procesos_Z++
			}
			if stat.State[0]=='S'{
				Procesos_S++
			}
			if stat.State[0]=='T'{
				Procesos_T++
			}
			if stat.State[0]=='D'{
				Procesos_D++
			}
		}
	}
	fmt.Println("Process Running:", Procesos_R)
	fmt.Println("Process Sleeping (uninterruptible sleep + interruptible sleep):", (Procesos_S+Procesos_D))
	fmt.Println("Process Zombies:", Procesos_Z)
    fmt.Println("Process Stopped:", Procesos_T)
    
}