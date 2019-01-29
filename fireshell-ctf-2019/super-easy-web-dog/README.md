# Nome

Super-easy web dog

# Autor

diofeher

# Solves

2

# Descrição

Who doesn't love cute dogs?

# Flag

F#{G0L4NG_net/http_T3mpl4tE_1nJecTION}

# Write-up

Server-side template injection with Golang's html/template.

The valid payload is:

{{ $.Request }} or {{ .Request }}
