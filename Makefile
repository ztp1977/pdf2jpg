# convert
in=data
out=out

build:
	go build -o pkg/pdf2jpg app/main.go
install: build
	cp pkg/pdf2jpg ${GOPATH}/bin/

# make convert in=${inDir} out=${outDir}
convert:
	pkg/pdf2jpg -in ${in} -out ${out}

# screenshot setting
# alias screenshot-switch="make -f /Users/Shared/Screenshots/Makefile screenshotSetting"
# screenshot-switch subDir=aaa prefix=img
scRoot=/Users/Shared/Screenshots/
subDir=
prefix=img
screenshotSetting:
	mkdir -p ${scRoot}${subDir}
	defaults write com.apple.screencapture location ${scRoot}${subDir}
	defaults write com.apple.screencapture name ${prefix}
	defaults write com.apple.screencapture type jpg
	defaults write com.apple.screencapture include-date true
	killall SystemUIServer
show:
	imgcat ${screenshotRoot}${subDir}/*.png

