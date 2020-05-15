----------------------------
-- DATA

data_cop = { 
   ["position"] = {3,3}, ["renderable"] = {{0,0,255,255}, "c"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "cop"
}

data_thug = {
   ["position"] = {7,7}, ["renderable"] = {{255,0,0,255}, "h"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "thug"
}


------------------------------
--FUNCTIONS

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

function spawn_npc(data)
   --test Go interop 
   ent = Ent()
   --print(ent)

   ent:SetupComponentsMap()
   --add components from Lua
   pos = Position()
   -- lua indexes from 1 not 0 as Python does
   pos.Pos.X = data["position"][1]
   pos.Pos.Y = data["position"][2]

   render = Renderable()
   render.Color.R = data["renderable"][1][1]
   render.Color.G = data["renderable"][1][2]
   render.Color.B = data["renderable"][1][3]
   render.Color.A = data["renderable"][1][4]
   render.Glyph = string.byte(data['renderable'][2], 1) -- 99 -- 'c'

   -- print("Glyph: ", render.Glyph)
   name = Name()
   name.Name = data["name"]
   --print("Name", name.Name)

   stats = Stats()
   stats.Hp = data["stats"]["hp"]
   stats.Max_hp = data["stats"]["hp"]
   stats.Power = data["stats"]["power"]

   -- these two aren't defined in data because no need to
   block = Blocker()
   npc = NPC()

   --the minus is the unary operator, dereferencing pointers
   ent:AddComponent("position", -pos)
   ent:AddComponent("renderable", -render)
   ent:AddComponent("name", -name)
   ent:AddComponent("blocker", -block)
   ent:AddComponent("NPC", -npc)
   ent:AddComponent("stats", -stats)

   entities:add(ent)
end

-- do stuff!
spawn_npc(data_cop)
spawn_npc(data_thug)