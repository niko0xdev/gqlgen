software_version	  		=	$(shell cat VERSION)
version_array		    	=	$(subst ., ,$(software_version))
major				        = 	$(word 1,${version_array})
minor				        = 	$(word 2,${version_array})
patch				        = 	$(word 3,${version_array})

patch:
	- @echo "BUMPING PATCH"
	- @echo "Current Version: $(software_version)"
	- $(eval patch=$(shell echo $$(($(patch)+1))))
	- @echo "New Version: $(major).$(minor).$(patch)"
	- @printf $(major).$(minor).$(patch) > VERSION
	- git tag "$(major).$(minor).$(patch)" -m "Patch version update: $(major).$(minor).$(patch)"
	- git push origin $(major).$(minor).$(patch)

minor:
	- @echo "BUMPING MINOR"
	- @echo "Current Version: $(software_version)"
	- $(eval minor=$(shell echo $$(($(minor)+1))))
	- @echo "New Version: $(major).$(minor).0"
	- @printf $(major).$(minor).0 > VERSION
	- git tag "$(major).$(minor).$(patch)" -m "Minor version update: $(major).$(minor).$(patch)"
	- git push origin $(major).$(minor).$(patch)

major:
	- @echo "BUMPING MAJOR"
	- @echo "Current Version: $(software_version)"
	- $(eval major=$(shell echo $$(($(major)+1))))
	- @echo "New Version: $(major).0.0"
	- @printf $(major).0.0 > VERSION
	- git tag "$(major).$(minor).$(patch)" -m "Major version update: $(major).$(minor).$(patch)"
	- git push origin $(major).$(minor).$(patch)