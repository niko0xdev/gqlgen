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
	- cp VERSION graphql/VERSION
	- git add --all
	- git commit -m "Bump version to $(major).$(minor).$(patch)"
	- git tag "$(major).$(minor).$(patch)" -m "Patch version update: $(major).$(minor).$(patch)"
	- git push origin $(major).$(minor).$(patch) $(git branch --show-current)

minor:
	- @echo "BUMPING MINOR"
	- @echo "Current Version: $(software_version)"
	- $(eval minor=$(shell echo $$(($(minor)+1))))
	- @echo "New Version: $(major).$(minor).0"
	- cp VERSION graphql/VERSION
	- git add --all
	- git commit -m "Bump version to $(major).$(minor).0"
	- git tag "$(major).$(minor).$(patch)" -m "Minor version update: $(major).$(minor).0"
	- git push origin $(major).$(minor).0 $(git branch --show-current)

major:
	- @echo "BUMPING MAJOR"
	- @echo "Current Version: $(software_version)"
	- $(eval major=$(shell echo $$(($(major)+1))))
	- @echo "New Version: $(major).0.0"
	- @printf $(major).0.0 > VERSION
	- cp VERSION graphql/VERSION
	- git add --all
	- git commit -m "Bump version to $(major).0.0"
	- git tag "$(major).0.0" -m "Major version update: $(major).0.0"
	- git push origin $(major).0.0 $(git branch --show-current)