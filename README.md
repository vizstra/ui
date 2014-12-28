Vistra UI is an experimental OpenGL based GUI library for the Go programming language. It a personal project and should be viewed as such, but feel free to use it and submit pull requests as bugs are found.

I chose OpenGL for rendering to have a less event driven UI experience and a simple architecture.  Each frame renders itself as often as possible in separate goroutine which is kicked off by the Window.Start() function.  

To get the library for use:
```
go get github.com/vizstra/ui
```

######PROTIPs
Use rerun to automatically test and rebuild changes.
```
go install https://github.com/skelterjohn/rerun
```

#### Current UI Goals
- [x] Button
- [ ] Text (This is a very large area currently being worked)
- [ ] Label
- [ ] Scrollbar
- [ ] List
- [ ] Radio
- [ ] Check
- [x] Progress Bar
- [ ] Image Button
- [ ] Tree
- [ ] Table
- [x] Line Chart
- [ ] Pie Chart

###### Layouts
- [x] Fill
- [x] Table
- [ ] Grid


