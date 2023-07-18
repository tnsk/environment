# environment

#### şimdilik sadece bu kütüphanede sürekli yazmaktan bıktığım environment kontrolcüsünü ekliyorum ileride eklemeler yapabilirim.

---

### Usage


```GO
package main

import (
    "log"
    ec "github/tnsk/environment"
)
func main(){
	if err := ec.Checks([]string{"PATH", "LOG_FILE"}); err != nil {
		log.Fatal(err)
	}
}
```



### Test Result

Valid ENV {"GOPATH", "SHELL"}

Invalid ENV {"GOPATH", "SHELL", "FISTIKCISAHAP"}

```text
=== RUN   TestChecks
=== RUN   TestChecks/Check_Valid
=== RUN   TestChecks/Check_Invalid
--- PASS: TestChecks (0.00s)
    --- PASS: TestChecks/Check_Valid (0.00s)
    --- PASS: TestChecks/Check_Invalid (0.00s)
PASS
ok      github.com/tnsk/environment     (cached)

```