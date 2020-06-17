----------------------------
-- DATA

data_cop = { 
   ["position"] = {3,3}, ["renderable"] = {{0,0,255,255}, "c"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "cop"
}

data_thug = {
   ["position"] = {7,7}, ["renderable"] = {{255,0,0,255}, "h"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "thug"
}

data_medkit = {
   ["position"] = {4,4}, ["renderable"] = {{255,0,0,255}, "!"}, ["name"] = "medkit", ["comps"] = "medkit"
}

data_map = {
   ["wall"] = {"#", {255, 255,255,255}}, ["floor"] = {".", { 200,200,200, 255 } }
}

------------------------------
--FUNCTIONS

print("hello WASM from lua!!")

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

   --closest free position
   --the minus is the unary operator, dereferencing pointers
   pos_s = map:FreeGridInRange(20, -pos.Pos)
   --print(dump(pos_s))
   if #pos_s < 1 then
      return
   else
      pos.Pos = -pos_s[1]
      print(pos.Pos.X)
      print(pos.Pos.Y)
   end

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

function spawn_item(data)
   ent = Ent()
   ent:SetupComponentsMap()
   --add components from Lua
   pos = Position()
   -- lua indexes from 1 not 0 as Python does
   pos.Pos.X = data["position"][1]
   pos.Pos.Y = data["position"][2]

   --closest free position
   --the minus is the unary operator, dereferencing pointers
   pos_s = map:FreeGridInRange(20, -pos.Pos)
   --print(dump(pos_s))
   -- table length
   if #pos_s < 1 then
      return
   else
      pos.Pos = -pos_s[1]
   --   print(pos.Pos.X)
   --   print(pos.Pos.Y)
   end

   render = Renderable()
   render.Color.R = data["renderable"][1][1]
   render.Color.G = data["renderable"][1][2]
   render.Color.B = data["renderable"][1][3]
   render.Color.A = data["renderable"][1][4]
   render.Glyph = string.byte(data['renderable'][2], 1)

   name = Name()
   name.Name = data["name"]

   -- no need to define this one in data
   item = Item()
   
   --the minus is the unary operator, dereferencing pointers
   ent:AddComponent("position", -pos)
   ent:AddComponent("renderable", -render)
   ent:AddComponent("name", -name)
   ent:AddComponent("item", -item)

   --additional components
   if data["comps"] == "medkit" then
      comp = Medkit()
      comp.Heal = 4
      ent:AddComponent("medkit", -comp)
   end

   entities:add(ent)
end

function make_map(data)
   --print("Making map...")
   print(dump(data))

   wall_color = Color()
   wall_color.R = data["wall"][2][1]
   wall_color.G = data["wall"][2][2]
   wall_color.B = data["wall"][2][3]
   wall_color.A = data["wall"][2][4]

   floor_color = Color()
   floor_color.R = data["floor"][2][1]
   floor_color.G = data["floor"][2][2]
   floor_color.B = data["floor"][2][3]
   floor_color.A = data["floor"][2][4]

   --convert string to Go's rune equivalent (int)
   map:GenerateArenaMapData(string.byte(data["wall"][1], 1), string.byte(data["floor"][1],1), -wall_color, -floor_color)
end

-- do stuff!
spawn_npc(data_cop)
spawn_npc(data_thug)
spawn_item(data_medkit)

--print("Map...")
--make_map(data_map)