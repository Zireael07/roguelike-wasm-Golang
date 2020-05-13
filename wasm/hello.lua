print("hello WASM from lua")

-- very basic function to dump table to readable format
function dump(o)
   if type(o) == 'table' then
      local s = '{ '
      for k,v in pairs(o) do
         if type(k) ~= 'number' then k = '"'..k..'"' end
         s = s .. '['..k..'] = ' .. dump(v) .. ','
      end
      return s .. '} '
   else
      return tostring(o)
   end
end


data_all = { 
    ["position"] = {2,2}, ["renderable"] = {{0,0,255,255}, "c"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "cop"
}

--test Go interop 
ent = Ent()
--print(ent)

ent:SetupComponentsMap()
--add components from Lua
pos = Position()
pos.Pos.X = 3
pos.Pos.Y = 3
-- print(pos.Pos.X)
-- print(pos.Pos.Y)

render = Renderable()
render.Color.R = 0
render.Color.G = 0
render.Color.B = 255
render.Color.A = 255
render.Glyph = 99 -- 'c'

-- print("Glyph: ", render.Glyph)
name = Name()
name.Name = "Cop"
--print("Name", name.Name)

block = Blocker()
npc = NPC()
stats = Stats()
stats.Hp = 10
stats.Max_hp = 10
stats.Power = 2

--the minus is the unary operator, dereferencing pointers
ent:AddComponent("position", -pos)
ent:AddComponent("renderable", -render)
ent:AddComponent("name", -name)
ent:AddComponent("blocker", -block)
ent:AddComponent("NPC", -npc)
ent:AddComponent("stats", -stats)

entities:add(ent)

--(calling functions)
--ent:AddComponentLua("TestComp")

--test Go interop (passing data)
-- data = {
--    name = "Test"
-- }



 --print(dump(data))
