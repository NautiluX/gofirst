APP_NAME=gofirst
PACKAGE_NAME=com.yourname.gofirst
AAR_PATH=./android/app/libs/$(APP_NAME).aar

.PHONY: help run bind-android bundle-android clean

run:
	go run main.go

bind-android:
	@echo "Binding Go to Android..."
	mkdir -p android/app/libs
	go tool ebitenmobile bind -target android -javapkg $(PACKAGE_NAME) -o $(AAR_PATH) ./mobile

bundle-android: bind-android
	@echo "Building Android App Bundle..."
	cd android && ./gradlew bundleRelease
	@echo "Done! Check: android/app/build/outputs/bundle/release/"

clean:
	rm -f $(AAR_PATH)
	cd android && ./gradlew clean
