package gateway_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/RafaySystems/rcloud-base/components/common/pkg/gateway"
	"github.com/RafaySystems/rcloud-base/components/common/pkg/gateway/testdata"
)

func TestRafayJSONMarshaller(t *testing.T) {
	m := gateway.NewRafayJSON()

	now := time.Now()

	t1 := testdata.TestYAML{
		Name:   "test",
		Time:   &now,
		Labels: map[string]string{"l1": "l2"},
	}

	yb, err := m.Marshal(&t1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(yb))

	var t2 testdata.TestYAML

	err = m.Unmarshal(yb, &t2)
	if err != nil {
		t.Error(err)
	}

	t.Log(t2)

	bb1 := new(bytes.Buffer)

	bb1.Write(yb)

	dec := m.NewDecoder(bb1)
	var t3 testdata.TestYAML
	err = dec.Decode(&t3)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(t2)

	bb2 := new(bytes.Buffer)

	enc := m.NewEncoder(bb2)
	err = enc.Encode(&t1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(bb2.String())

}