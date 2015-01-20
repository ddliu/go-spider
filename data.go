// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package spider

type Data map[string]interface{}

func (this Data) Has(key string) bool {
    _, exists := this[key]

    return exists
}

func (this Data) GetString(key string) (string, bool) {
    if v, ok := this[key]; ok {
        if vv, ok := v.(string); ok {
            return vv, true
        }
    }

    return "", false
}

func (this Data) MustGetString(key string) string {
    if v, ok := this.GetString(key); ok {
        return v
    }

    panic("Task has no string data: " + key)
}

func (this Data) GetInt(key string) (int, bool) {
    if v, ok := this[key]; ok {
        if vv, ok := v.(int); ok {
            return vv, true
        }
    }

    return 0, false
}

func (this Data) MustGetInt(key string) int {
    if v, ok := this.GetInt(key); ok {
        return v
    }

    panic("Task has no int data: " + key)
}

func (this Data) GetBytes(key string) ([]byte, bool) {
    if v, ok := this[key]; ok {
        if vv, ok := v.([]byte); ok {
            return vv, true
        }
    }

    return nil, false
}

func (this Data) MustGetBytes(key string) ([]byte) {
    if v, ok := this.GetBytes(key); ok {
        return v
    }

    panic("Task has no bytes data: " + key)
}