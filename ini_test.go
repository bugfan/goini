package goini
import(
	"testing"	
	"log"
)
func TestIni(t *testing.T) {
	LoadConfig("conf/test.conf")
	log.Println("num:",Config.GetInt64("num"))
	log.Println("float:",Config.GetInt64("float"))
	log.Println("admin:",Config.GetString("admin"))
	log.Println("tiger:",Config.GetString("tiger"))
	log.Println("tiger_int64:",Config.GetInt64("tiger"))

	log.Println("zhang_default:",Config.GetString("zhang"))	
	log.Println("zhang_section:",Config.GetSectionString(":","zhang"))
	
	
}
