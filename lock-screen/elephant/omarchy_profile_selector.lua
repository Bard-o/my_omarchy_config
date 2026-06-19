Name = "omarchyProfileSelector"
NamePretty = "Omarchy Avatar Selector"
Cache = false
HideFromProviderlist = true
SearchName = true

local function ShellEscape(s)
  return "'" .. s:gsub("'", "'\\''") .. "'"
end

function FormatName(filename)
  local name = filename:gsub("^%d+", ""):gsub("^%-", "")
  name = name:gsub("%.[^%.]+$", "")
  name = name:gsub("-", " ")
  name = name:gsub("%S+", function(word)
    return word:sub(1, 1):upper() .. word:sub(2):lower()
  end)
  return name
end

function GetEntries()
  local entries = {}
  local home = os.getenv("HOME")
  local images_dir = home .. "/.config/omarchy/profile_images"

  local cmd = "/usr/bin/find -L " .. ShellEscape(images_dir)
    .. " -maxdepth 1 -type f \\( -name '*.jpg' -o -name '*.jpeg' "
    .. "-o -name '*.png' -o -name '*.gif' -o -name '*.bmp' "
    .. "-o -name '*.webp' \\) 2>/dev/null | /usr/bin/sort"

  local f = io.popen(cmd)
  if f then
    for filepath in f:lines() do
      local filename = filepath:match("([^/]+)$")
      if filename then
        table.insert(entries, {
          Text = FormatName(filename),
          Value = filepath,
          Actions = {
            activate = "cp " .. ShellEscape(filepath) .. " " .. ShellEscape(home .. "/.face"),
          },
          Preview = filepath,
          PreviewType = "file",
        })
      end
    end
    f:close()
  end

  return entries
end
