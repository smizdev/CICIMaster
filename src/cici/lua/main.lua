for a=1, 1000 do
	local a = a
	thread(function()
		while true do			
			local rand = math.random(0, a*1000)
			print(string.format("thread %s %s %s", a, rand, os.clock()))			
			sleep(rand)						
		end
	end)
end