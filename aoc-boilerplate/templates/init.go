package templates

func Initialize() {
    templates = make(map[string]string)
    fileExtensions = make(map[string]string)

    templates["python"] = "CiMgUmVhZCBpbiB0aGUgaW5wdXQgZmlsZS4Kd2l0aCBvcGVuKCJpbnB1dC50eHQiLCAiciIpIGFzIG15ZmlsZToKICAgIGlucHV0X2ZpbGUgPSBbaS5yc3RyaXAoIlxuIikgZm9yIGkgaW4gbXlmaWxlLnJlYWRsaW5lcygpXQoK"
    fileExtensions["python"] = "py"
    templates["go"] = "cGFja2FnZSBtYWluCgppbXBvcnQgKAoJImJ1ZmlvIgoJImZsYWciCgkib3MiCikKCmZ1bmMgbWFpbigpIHsKCWJvb2xQdHIgOj0gZmxhZy5Cb29sKCJ0ZXN0IiwgZmFsc2UsICJ0ZXN0IG1vZGUiKQoJZmxhZy5QYXJzZSgpCgoJdmFyIGZpbGVuYW1lIHN0cmluZwoJaWYgKmJvb2xQdHIgewoJCWZpbGVuYW1lID0gImlucHV0LXRlc3QudHh0IgoJfSBlbHNlIHsKCQlmaWxlbmFtZSA9ICJpbnB1dC50eHQiCgl9CglpbnB1dF9maWxlLCBlcnIgOj0gcmVhZExpbmVzKGZpbGVuYW1lKQoJY2hlY2soZXJyKQoJLy8gWW91ciBDb2RlIGdvZXMgYmVsb3chCgp9CgpmdW5jIGNoZWNrKGUgZXJyb3IpIHsKCWlmIGUgIT0gbmlsIHsKCQlwYW5pYyhlKQoJfQp9CgpmdW5jIHJlYWRMaW5lcyhwYXRoIHN0cmluZykgKFtdc3RyaW5nLCBlcnJvcikgewoJZmlsZSwgZXJyIDo9IG9zLk9wZW4ocGF0aCkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBuaWwsIGVycgoJfQoJZGVmZXIgZmlsZS5DbG9zZSgpCgoJdmFyIGxpbmVzIFtdc3RyaW5nCglzY2FubmVyIDo9IGJ1ZmlvLk5ld1NjYW5uZXIoZmlsZSkKCWZvciBzY2FubmVyLlNjYW4oKSB7CgkJbGluZXMgPSBhcHBlbmQobGluZXMsIHNjYW5uZXIuVGV4dCgpKQoJfQoJcmV0dXJuIGxpbmVzLCBzY2FubmVyLkVycigpCn0K"
    fileExtensions["go"] = "go"
}