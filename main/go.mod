module ikonchallenge/main

replace ikonchallenge/device => ../device

replace ikonchallenge/file => ../file

replace ikonchallenge/service => ../service

go 1.16

require (
	ikonchallenge/device v0.0.0-00010101000000-000000000000 // indirect
	ikonchallenge/file v0.0.0-00010101000000-000000000000
	ikonchallenge/service v0.0.0-00010101000000-000000000000
)
