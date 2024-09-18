This is a language server for markdown files in Neovim.
With basic utils:

- Hover.
- Go to definition.

How to use:

1. Run

```sh
go build main.go
```

To build `main`.

2. Then under `lua/after/plugin` in Neovim config, create `something.lua`:

```lua
local client = vim.lsp.start_client {
  name = "edulsp",
  cmd = { "path/to/main/created/above" },
  on_attach = require("util.lsp").on_attach, -- If you have one.
}

if not client then
  vim.notify "Error: bad client"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

This LSP will attach to .md files allowing you to use keymaps in your LSP's `on_attach`.
To confirm the LSP is running, open a .md file in Neovim then `:LspInfo`.
