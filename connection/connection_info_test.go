package connection

import (
  "testing"
)

func TestReadFile(t *testing.T) {
  const s = `apiaddress: https://192.0.2.10:8006/api2/json
userid:
  username: pveexporter
  idrealm: pve
apitoken:
  id: prometheus
  token: aaaaa-bbbbbb-ccccc-dd
`
  var c Info
  err2 := c.parseYaml([]byte(s))
  if err2 != nil {
    t.Errorf("struct not able to parse YAML")
  }
}
