package check

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheck_Parse(t *testing.T) {
	t.Run("singleDocument", func(t *testing.T) {
		data := strings.TrimSpace(`
---
checks:
  - version: 1
    type: MYSQL_SHOW
    query: VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');
    script: |
        def function1(args):
            pass

  - version: 2
    type: POSTGRESQL_SELECT
    query: id, name FROM table WHERE id=123;
    script: |
        def function2(args):
            pass
`)

		cs, err := Parse(bytes.NewReader([]byte(data)))
		require.NoError(t, err)

		assert.Len(t, cs, 2)

		assert.Equal(t, uint32(1), cs[0].Version)
		assert.Equal(t, MySQLShow, cs[0].Type)
		assert.Equal(t, "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", cs[0].Query)
		assert.Equal(t, cs[0].Script, "def function1(args):\n    pass\n")

		assert.Equal(t, uint32(2), cs[1].Version)
		assert.Equal(t, PostgreSQLSelect, cs[1].Type)
		assert.Equal(t, "id, name FROM table WHERE id=123;", cs[1].Query)
		assert.Equal(t, cs[1].Script, "def function2(args):\n    pass")
	})

	t.Run("multipleDocuments", func(t *testing.T) {
		data := strings.TrimSpace(`
---
checks:
  - version: 1
    type: MYSQL_SHOW
    query: VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');
    script: |
        def function1(args):
            pass
---
checks:
  - version: 2
    type: POSTGRESQL_SELECT
    query: id, name FROM table WHERE id=123;
    script: |
        def function2(args):
            pass
`)
		cs, err := Parse(bytes.NewReader([]byte(data)))
		require.NoError(t, err)

		assert.Len(t, cs, 2)

		assert.Equal(t, uint32(1), cs[0].Version)
		assert.Equal(t, MySQLShow, cs[0].Type)
		assert.Equal(t, "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", cs[0].Query)
		assert.Equal(t, cs[0].Script, "def function1(args):\n    pass\n")

		assert.Equal(t, uint32(2), cs[1].Version)
		assert.Equal(t, PostgreSQLSelect, cs[1].Type)
		assert.Equal(t, "id, name FROM table WHERE id=123;", cs[1].Query)
		assert.Equal(t, cs[1].Script, "def function2(args):\n    pass")
	})
}

func TestCheck_CheckValidate(t *testing.T) {
	tests := []struct {
		name   string
		check  *Check
		errStr string
	}{
		{
			name:   "mysql_show",
			check:  &Check{Type: MySQLShow, Query: "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", Script: "def func(args): pass"},
			errStr: "",
		},
		{
			name:   "mysql_select",
			check:  &Check{Type: MySQLSelect, Query: "id, name FROM table WHERE id=123;", Script: "def func(args): pass"},
			errStr: "",
		},
		{
			name:   "postgresql_show",
			check:  &Check{Type: PostgreSQLShow, Query: "", Script: "def func(args): pass"},
			errStr: "",
		},
		{
			name:   "postgresql_select",
			check:  &Check{Type: PostgreSQLSelect, Query: "id, name FROM table WHERE id=123;", Script: "def func(args): pass"},
			errStr: "",
		},
		{
			name:   "mongodb_get_parameter",
			check:  &Check{Type: MongoDBGetParameter, Query: "\"saslHostName\" : 1", Script: "def func(args): pass"},
			errStr: "",
		},
		{
			name:   "clickhouse_show",
			check:  &Check{Type: "CLICKHOUSE_SHOW", Query: "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", Script: "def func(args): pass"},
			errStr: "unknown check type: CLICKHOUSE_SHOW",
		},
		{
			name:   "empty_type",
			check:  &Check{Type: "", Query: "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", Script: "def func(args): pass"},
			errStr: "check type is empty",
		},
		{
			name:   "empty_query",
			check:  &Check{Type: MySQLShow, Query: "", Script: "def func(args): pass"},
			errStr: "check query is empty",
		},
		{
			name:   "non_empty_query_for_postgresql_show",
			check:  &Check{Type: PostgreSQLShow, Query: "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", Script: "def func(args): pass"},
			errStr: "POSTGRESQL_SHOW check type should have empty query",
		},
		{
			name:   "empty_script",
			check:  &Check{Type: MySQLShow, Query: "VARIABLES WHERE Variable_name IN ('have_ssl', 'have_openssl');", Script: ""},
			errStr: "check script is empty",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.check.validate()

			if tt.errStr != "" {
				assert.EqualError(t, err, tt.errStr)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestCheck_ResultValidate(t *testing.T) {
	tests := []struct {
		name   string
		result *Result
		errStr string
	}{
		{
			name:   "success_result_without_message",
			result: &Result{Status: Success, Message: ""},
			errStr: "",
		},
		{
			name:   "success_result_with_message",
			result: &Result{Status: Success, Message: "everything is fine!"},
			errStr: "",
		},
		{
			name:   "failed_result_with_message",
			result: &Result{Status: Fail, Message: "something bad happened!"},
			errStr: "",
		},
		{
			name:   "failed_result_without_message",
			result: &Result{Status: Fail, Message: ""},
			errStr: "failed check result should have message",
		},
		{
			name:   "empty_status",
			result: &Result{Status: "", Message: ""},
			errStr: "result status is empty",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.result.Validate()

			if tt.errStr != "" {
				assert.EqualError(t, err, tt.errStr)
				return
			}

			assert.NoError(t, err)
		})
	}
}

const data = `random data`

const publicKey = `RWRQmBOLeYzAeuR2L6L1GJN9qTR8ceQrawtijPTQkVbf3LJsrLeUjQcL`

const signature = `untrusted comment: signature from minisign secret key
RWRQmBOLeYzAetS6fGVWAvzwCgDuo/zNlvdOrClAvjCUSMLnUimp6NQd1L+x77HZa0kEB7ei+K9lW+W4hIf1D8gRNm+cdQr7dgk=
trusted comment: timestamp:1586854934	file:data
WXAxVyC6G82QuXtGlJZzLWoVmw8QNWks2T6RfXo8F9oKjI+sPbBf0ZOBWD2hXKFBCo5pKPSJiaVeI4G36OlEAw==
`

func TestCheck_Verify(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		err := Verify([]byte(data), publicKey, signature)
		require.NoError(t, err)
	})

	t.Run("invalid signature", func(t *testing.T) {
		err := Verify([]byte(data), publicKey, strings.TrimSpace(`
untrusted comment: signature from minisign secret key
RWRQmBOLeYzAetS6fGVWAvzwCgDuo/zNlvdOrClAvjCUSMLnUimp6NQd1L+f3fHZa0kEB7ei+K9lW+W4hIf+INVALID+INVALID=
trusted comment: timestamp:1586854934	file:data
WXAxVyC6G82QuXtGlJZzLWoVmw8QNWks2T6RfXo8F9oKjI+sPbBf0ZOBWD2hXKFBCo5pKPSJiaVeI4G36OlEAw==`))

		assert.EqualError(t, err, "invalid signature")
	})

	t.Run("invalid global signature", func(t *testing.T) {
		err := Verify([]byte(data), publicKey, strings.TrimSpace(`
untrusted comment: signature from minisign secret key
RWRQmBOLeYzAetS6fGVWAvzwCgDuo/zNlvdOrClAvjCUSMLnUimp6NQd1L+x77HZa0kEB7ei+K9lW+W4hIf1D8gRNm+cdQr7dgk=
trusted comment: timestamp:1586854934	file:data
WXAxVyC6G82QuXtGlJZzLWoVmw8QNWks2veRfXo8F9oKjI+sPbBf0ZOBWD2hXKFBCo5pKP+INVALID+INVALID==`))
		assert.EqualError(t, err, "invalid global signature")
	})

	t.Run("invalid trusted comment", func(t *testing.T) {
		err := Verify([]byte(data), publicKey, strings.TrimSpace(`
untrusted comment: signature from minisign secret key
RWRQmBOLeYzAetS6fGVWAvzwCgDuo/zNlvdOrClAvjCUSMLnUimp6NQd1L+x77HZa0kEB7ei+K9lW+W4hIf1D8gRNm+cdQr7dgk=
trusted comment: timestamp:1586854934	file:INVALID COMMENT
WXAxVyC6G82QuXtGlJZzLWoVmw8QNWks2T6RfXo8F9oKjI+sPbBf0ZOBWD2hXKFBCo5pKPSJiaVeI4G36OlEAw==`))
		assert.EqualError(t, err, "invalid global signature")
	})

	t.Run("invalid public key", func(t *testing.T) {
		err := Verify([]byte("random data"), "RWRQmBOLeYzAeu5FL8f1JMN9qTR8CDfrabdtjPTQ+INVALID+INVALID", signature)
		assert.EqualError(t, err, "invalid signature")
	})

	t.Run("empty data", func(t *testing.T) {
		err := Verify(nil, publicKey, signature)
		assert.EqualError(t, err, "invalid signature")
	})

	t.Run("empty signature", func(t *testing.T) {
		err := Verify([]byte(data), publicKey, "")
		assert.EqualError(t, err, "incomplete signature")
	})

	t.Run("empty key", func(t *testing.T) {
		err := Verify([]byte(data), "", signature)
		assert.EqualError(t, err, "invalid public key")
	})
}
