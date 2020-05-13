print("hello WASM from lua")



data_all = { 
    ["position"] = {2,2}, ["renderable"] = {{0,0,255,255}, "c"}, ["stats"] = { hp = 10, power =2 }, ["name"] = "cop"
}

--test Go interop (calling functions)
ent:AddComponentLua("TestComp")

--test Go interop (passing data)
data = {
   name = "Test"
}



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

 print(dump(data))

