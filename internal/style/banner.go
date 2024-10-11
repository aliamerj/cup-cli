package style

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBanner() {
	banner := `
  ____ _                 _    ____            
 / ___| | ___  _   _  __| |  / ___|   _ _ __  
| |   | |/ _ \| | | |/ _' | | |  | | | | '_ \ 
| |___| | (_) | |_| | (_| | | |__| |_| | |_) |
 \____|_|\___/ \__,_|\__,_|  \____\__,_| .__/ 
                                       |_|    
`
	fmt.Println(color.HiCyanString(banner))

	fmt.Println(color.HiYellowString(" Welcome to Cloud-Cup CLI - Your modern reverse proxy solution\n"))
}
