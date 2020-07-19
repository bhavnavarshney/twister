#!/bin/sh

APP="NPTCalibration.app"
mkdir -p build/$APP/Contents/{MacOS,Resources}
go generate
go build -o build/$APP/Contents/MacOS/NPTCalibration
cat > build/$APP/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>NPTCalibration</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.zserge.lorca.example</string>
</dict>
</plist>
EOF
cp icons/icon.icns build/$APP/Contents/Resources/icon.icns
find build/$APP
